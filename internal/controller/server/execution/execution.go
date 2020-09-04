package execution

import (
	"context"
	"errors"
	"fmt"
	"github.com/jonboulle/clockwork"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"octavius/internal/controller/config"
	executorRepo "octavius/internal/controller/server/repository/executor"
	jobRepo "octavius/internal/controller/server/repository/job"
	metadataRepo "octavius/internal/controller/server/repository/metadata"
	"octavius/internal/controller/server/scheduler"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"sync"
	"time"
)

// Execution interface for methods related to execution
type Execution interface {
	SaveMetadata(ctx context.Context, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error)
	ReadAllMetadata(ctx context.Context) (*clientCPproto.MetadataArray, error)
	RegisterExecutor(ctx context.Context, request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error)
	UpdateExecutorStatus(ctx context.Context, request *executorCPproto.Ping) (*executorCPproto.HealthResponse, error)
	ExecuteJob(ctx context.Context, request *clientCPproto.RequestForExecute) (uint64, error)
}
type execution struct {
	metadataRepo      metadataRepo.Repository
	executorRepo      executorRepo.Repository
	jobRepo           jobRepo.Repository
	idGenerator       idgen.RandomIdGenerator
	scheduler         scheduler.Scheduler
	activeExecutorMap *activeExecutorMap
}
type activeExecutor struct {
	sessionID  uint64
	healthChan chan string
	timer      <-chan time.Time
}
type activeExecutorMap struct {
	execMap *sync.Map
}

func (m *activeExecutorMap) Get(key string) (*activeExecutor, bool) {
	exec, ok := m.execMap.Load(key)
	if ok {
		return exec.(*activeExecutor), ok
	}
	return nil, ok
}
func (m *activeExecutorMap) Put(key string, executor *activeExecutor) {
	m.execMap.Store(key, executor)
}
func (m *activeExecutorMap) Delete(key string) {
	m.execMap.Delete(key)
}

// NewExec creates a new instance of metadata respository
func NewExec(metadataRepo metadataRepo.Repository, executorRepo executorRepo.Repository, jobRepo jobRepo.Repository, idGenerator idgen.RandomIdGenerator, scheduler scheduler.Scheduler) Execution {
	newActiveExecutorMap := &activeExecutorMap{
		execMap: new(sync.Map),
	}
	return &execution{
		metadataRepo:      metadataRepo,
		jobRepo:           jobRepo,
		executorRepo:      executorRepo,
		idGenerator:       idGenerator,
		scheduler:         scheduler,
		activeExecutorMap: newActiveExecutorMap,
	}
}

//SaveMetadata calls the repository/metadata Save() function and returns MetadataName
func (e *execution) SaveMetadata(ctx context.Context, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error) {
	return e.metadataRepo.Save(ctx, metadata.Name, metadata)
}

//ReadAllMetadata calls the repository/metadata GetAll() and returns MetadataArray
func (e *execution) ReadAllMetadata(ctx context.Context) (*clientCPproto.MetadataArray, error) {
	return e.metadataRepo.GetAll(ctx)
}

//RegisterExecutor saves executor information in DB
func (e *execution) RegisterExecutor(ctx context.Context, request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error) {
	key := request.ID
	value := request.ExecutorInfo
	return e.executorRepo.Save(ctx, key, value)
}
func removeActiveExecutor(activeExecutorMap *activeExecutorMap, id string, executor *activeExecutor) {
	log.Info(fmt.Sprintf("session id: %d, executor id : %s, closing executor session", executor.sessionID, id))
	close(executor.healthChan)
	activeExecutorMap.Delete(id)
}

//StartExecutionHealthCheck checks for executor ping at regular interval
func startExecutorHealthCheck(e *execution, activeExecutorMap *activeExecutorMap, id string) {
	executor, _ := activeExecutorMap.Get(id)
	ctx := context.Background()
	log.Info(fmt.Sprintf("session ID: %v, opening connection with executor: %s", executor.sessionID, id))
	err := e.executorRepo.UpdateStatus(ctx, id, "free")
	if err != nil {
		log.Error(err, fmt.Sprintf("session ID: %d, fail to write update status of executor with id: %s", executor.sessionID, id))
		removeActiveExecutor(activeExecutorMap, id, executor)
		return
	}
	for {
		select {
		case health := <-executor.healthChan:
			err := e.executorRepo.UpdateStatus(ctx, id, health)
			if err != nil {
				log.Error(err, fmt.Sprintf("session ID: %d, fail to write update status of executor with id: %s", executor.sessionID, id))
				removeActiveExecutor(activeExecutorMap, id, executor)
				return
			}
		case <-executor.timer:
			err := e.executorRepo.UpdateStatus(ctx, id, "expired")
			if err != nil {
				log.Error(err, fmt.Sprintf("session ID: %d, fail to write update status of executor with id: %s", executor.sessionID, id))
				removeActiveExecutor(activeExecutorMap, id, executor)
				return
			}
			log.Info(fmt.Sprintf("session ID: %v, deadline exceeded for executor with %s id", executor.sessionID, id))
			removeActiveExecutor(activeExecutorMap, id, executor)
			executor.timer = nil
			return
		}
	}
}
func (e *execution) UpdateExecutorStatus(ctx context.Context, request *executorCPproto.Ping) (*executorCPproto.HealthResponse, error) {
	executorID := request.ID
	clock := clockwork.NewRealClock()
	pingTimeOut := config.Config().ExecutorPingDeadline
	// if executor is already active
	if executor, ok := e.activeExecutorMap.Get(executorID); ok {
		executor.healthChan <- request.State
		executor.timer = clock.After(pingTimeOut)
		return &executorCPproto.HealthResponse{Recieved: true}, nil
	}
	//if executor is not registered in database
	_, err := e.executorRepo.Get(ctx, request.ID)
	if err != nil {
		if err.Error() == constant.NoValueFound {
			return nil, status.Error(codes.PermissionDenied, "executor not registered")
		}
		return nil, err
	}
	// if executor is registered and not yet active add it to activeExecutor map
	healthChan := make(chan string)
	sessionID, err := idgen.NewRandomIdGenerator().Generate()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	timer := clock.After(pingTimeOut)
	newActiveExecutor := activeExecutor{
		healthChan: healthChan,
		sessionID:  sessionID,
		timer:      timer,
	}
	e.activeExecutorMap.Put(executorID, &newActiveExecutor)
	go startExecutorHealthCheck(e, e.activeExecutorMap, executorID)
	return &executorCPproto.HealthResponse{Recieved: true}, nil
}
func getActiveExecutorMap(e *execution) *activeExecutorMap {
	return e.activeExecutorMap
}

// ExecuteJob function will call job repository and get jobId
func (e *execution) ExecuteJob(ctx context.Context, executionData *clientCPproto.RequestForExecute) (uint64, error) {
	jobAvailabilityStatus, err := e.jobRepo.CheckJobIsAvailable(ctx, executionData.JobName)
	if err != nil {
		return uint64(0), err
	}
	if jobAvailabilityStatus == false {
		return uint64(0), errors.New("job with given name not available")
	}
	jobId, err := e.idGenerator.Generate()
	if err != nil {
		return uint64(0), err
	}
	err = e.scheduler.AddToPendingList(ctx, jobId, executionData)
	if err != nil {
		return uint64(0), err
	}
	return jobId, err
}
