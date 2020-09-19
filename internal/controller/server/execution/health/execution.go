package health

import (
	"context"
	"fmt"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/protofiles"

	executorRepo "octavius/internal/controller/server/repository/executor"
	"octavius/internal/pkg/log"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// healthExecution interface for methods related to healthExecution
type HealthExecution interface {
	UpdateExecutorStatus(ctx context.Context, request *protofiles.Ping, pingTimeOut time.Duration) (*protofiles.HealthResponse, error)
}
type healthExecution struct {
	executorRepo      executorRepo.Repository
	activeExecutorMap *activeExecutorMap
}

type activeExecutor struct {
	sessionID  uint64
	pingChan   chan string
	statusChan chan string
	timer      *time.Timer
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
func NewHealthExec(executorRepo executorRepo.Repository) HealthExecution {
	newActiveExecutorMap := &activeExecutorMap{
		execMap: new(sync.Map),
	}
	return &healthExecution{
		executorRepo:      executorRepo,
		activeExecutorMap: newActiveExecutorMap,
	}
}

func removeActiveExecutor(activeExecutorMap *activeExecutorMap, id string, executor *activeExecutor) {
	log.Info(fmt.Sprintf("session id: %d, executor id : %s, closing executor session", executor.sessionID, id))
	close(executor.pingChan)
	activeExecutorMap.Delete(id)
}

//StarthealthExecutionHealthCheck checks for executor ping at regular interval
func startExecutorHealthCheck(e *healthExecution, activeExecutorMap *activeExecutorMap, id string) {
	executor, _ := activeExecutorMap.Get(id)
	ctx := context.Background()
	log.Info(fmt.Sprintf("session ID: %v, opening connection with executor: %s", executor.sessionID, id))
	err := e.executorRepo.UpdateStatus(ctx, id, constant.IdleState)
	if err != nil {
		log.Error(err, fmt.Sprintf("session ID: %d, fail to write update status of executor with id: %s", executor.sessionID, id))
		removeActiveExecutor(activeExecutorMap, id, executor)
		return
	}
	for {
		select {
		case health := <-executor.statusChan:
			err := e.executorRepo.UpdateStatus(ctx, id, health)
			if err != nil {
				log.Error(err, fmt.Sprintf("session ID: %d, fail to write update status of executor with id: %s", executor.sessionID, id))
				removeActiveExecutor(activeExecutorMap, id, executor)
				return
			}

		case <-executor.timer.C:
			err := e.executorRepo.UpdateStatus(ctx, id, "expired")
			if err != nil {
				log.Error(err, fmt.Sprintf("session ID: %d, fail to write update status of executor with id: %s", executor.sessionID, id))
				removeActiveExecutor(activeExecutorMap, id, executor)
				return
			}
			log.Info(fmt.Sprintf("session ID: %v, deadline exceeded for executor with %s id", executor.sessionID, id))
			removeActiveExecutor(activeExecutorMap, id, executor)
			executor.timer.Stop()
			return
		case <-executor.pingChan:
		}
	}
}
func (e *healthExecution) UpdateExecutorStatus(ctx context.Context, request *protofiles.Ping, pingTimeOut time.Duration) (*protofiles.HealthResponse, error) {
	executorID := request.ID
	// if executor is already active
	if executor, ok := e.activeExecutorMap.Get(executorID); ok {
		if request.State != "ping" {
			executor.statusChan <- request.State
		} else {
			executor.pingChan <- request.State
		}
		executor.timer.Reset(pingTimeOut)
		return &protofiles.HealthResponse{Recieved: true}, nil
	}
	//if executor is not registered in database
	_, err := e.executorRepo.Get(ctx, request.ID)
	if err != nil {
		if err.Error() == status.Error(codes.NotFound, constant.Etcd+constant.NoValueFound).Error() {
			return nil, status.Error(codes.PermissionDenied, "executor not registered")
		}
		return nil, err
	}
	// if executor is registered and not yet active add it to activeExecutor map
	pingChan := make(chan string)
	statusChan := make(chan string)
	sessionID, err := idgen.NewRandomIdGenerator().Generate()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	timer := time.NewTimer(pingTimeOut)
	newActiveExecutor := activeExecutor{
		sessionID:  sessionID,
		timer:      timer,
		pingChan:   pingChan,
		statusChan: statusChan,
	}
	e.activeExecutorMap.Put(executorID, &newActiveExecutor)
	go startExecutorHealthCheck(e, e.activeExecutorMap, executorID)
	return &protofiles.HealthResponse{Recieved: true}, nil
}

func getActiveExecutorMap(e *healthExecution) *activeExecutorMap {
	return e.activeExecutorMap
}
