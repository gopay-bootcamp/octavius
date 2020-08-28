package execution

import (
	"context"
	"fmt"
	"octavius/internal/controller/config"
	executorRepo "octavius/internal/controller/server/repository/executor"
	jobExecutorRepo "octavius/internal/controller/server/repository/jobExecutor"
	metadataRepo "octavius/internal/controller/server/repository/metadata"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Execution interface for methods related to execution
type Execution interface {
	SaveMetadata(ctx context.Context, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error)
	ReadAllMetadata(ctx context.Context) (*clientCPproto.MetadataArray, error)
	RegisterExecutor(ctx context.Context, request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error)
	UpdateExecutorStatus(ctx context.Context, request *executorCPproto.Ping) (*executorCPproto.HealthResponse, error)
	StartExecutorHealthCheck(activeExecutorMap *sync.Map, id string, executor activeExecutor)
	ExecuteJob(ctx context.Context, name string, data map[string]string) (uint64, error)
}

type execution struct {
	metadataRepo      metadataRepo.Repository
	executorRepo      executorRepo.Repository
	jobExecutorRepo   jobExecutorRepo.JobExecutionRepository
	activeExecutorMap *sync.Map

}

type activeExecutor struct {
	sessionID  uint64
	healthChan chan string
}

// NewExec creates a new instance of metadata respository
func NewExec(metadataRepo metadataRepo.Repository, executorRepo executorRepo.Repository, jobExecutorRepo jobExecutorRepo.JobExecutionRepository) Execution {
	return &execution{
		metadataRepo:      metadataRepo,
		jobExecutorRepo:  jobExecutorRepo,
		executorRepo:      executorRepo,
		activeExecutorMap: new(sync.Map),
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

func removeActiveExecutor(activeExecutorMap *sync.Map, id string, executor activeExecutor) {
	log.Info(fmt.Sprintf("session id: %d, executor id : %s, closing executor session", executor.sessionID, id))
	close(executor.healthChan)
	activeExecutorMap.Delete(id)
}

//StartExecutionHealthCheck checks for executor ping at regular interval
func (e *execution) StartExecutorHealthCheck(activeExecutorMap *sync.Map, id string, executor activeExecutor) {
	ctx := context.Background()
	timer := time.NewTimer(config.Config().ExecutorPingDeadline)
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
			timer.Stop()
			timer.Reset(config.Config().ExecutorPingDeadline)
		case <-timer.C:
			err := e.executorRepo.UpdateStatus(ctx, id, "expired")
			if err != nil {
				log.Error(err, fmt.Sprintf("session ID: %d, fail to write update status of executor with id: %s", executor.sessionID, id))
				removeActiveExecutor(activeExecutorMap, id, executor)
				return
			}
			log.Info(fmt.Sprintf("session ID: %v, deadline exceeded for executor with %s id", executor.sessionID, id))
			removeActiveExecutor(activeExecutorMap, id, executor)
			timer.Stop()
			return
		}
	}
}

func (e *execution) UpdateExecutorStatus(ctx context.Context, request *executorCPproto.Ping) (*executorCPproto.HealthResponse, error) {
	executorID := request.ID

	// construct to load channel if executor present in memory map
	if executor, ok := e.activeExecutorMap.Load(executorID); ok {
		executor.(activeExecutor).healthChan <- request.State
		return &executorCPproto.HealthResponse{Recieved: true}, nil
	}

	_, err := e.executorRepo.Get(ctx, request.ID)
	if err != nil {
		if err.Error() == constant.NoValueFound {
			return nil, status.Error(codes.PermissionDenied, "executor not registered")
		}
		return nil, err
	}

	// construct to make a new channel and add the executor to the in memory map
	healthChan := make(chan string)
	sessionID, err := idgen.NextID()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	newActiveExecutor := activeExecutor{
		healthChan: healthChan,
		sessionID:  sessionID,
	}
	e.activeExecutorMap.Store(executorID, newActiveExecutor)
	go e.StartExecutorHealthCheck(e.activeExecutorMap, executorID, newActiveExecutor)
	return &executorCPproto.HealthResponse{Recieved: true}, nil
}

//ExecuteJob function will call jobExecutor repository and get jonId
func (e *execution) ExecuteJob(ctx context.Context, jobName string, jobData map[string]string) (uint64, error) {
	jobId, err := e.jobExecutorRepo.ExecuteJob(ctx, jobName, jobData)
	return jobId, err
}
