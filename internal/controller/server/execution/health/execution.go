// Package health implements functions related to executor health
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

// HealthExecution interface for methods related to healthExecution
type HealthExecution interface {
	UpdatePingStatus(ctx context.Context, request *protofiles.Ping, pingTimeOut time.Duration) (*protofiles.HealthResponse, error)
}

type healthExecution struct {
	executorRepo      executorRepo.Repository
	activeExecutorMap *activeExecutorMap
}

type activeExecutor struct {
	sessionID uint64
	pingChan  chan struct{}
	timer     *time.Timer
}
type activeExecutorMap struct {
	execMap *sync.Map
}

// Get takes executor key as a argument and checks if executor is active or not
func (m *activeExecutorMap) Get(key string) (*activeExecutor, bool) {
	exec, ok := m.execMap.Load(key)
	if ok {
		return exec.(*activeExecutor), ok
	}
	return nil, ok
}

// Put takes executor key as a argument and adds executor with that key to executor map
func (m *activeExecutorMap) Put(key string, executor *activeExecutor) {
	m.execMap.Store(key, executor)
}

// Delete takes executor key as a argument and deletes executor with that key from executor map
func (m *activeExecutorMap) Delete(key string) {
	m.execMap.Delete(key)
}

// NewHealthExec creates a new instance of HealthRepository
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

// StartExecutionHealthCheck checks for executor ping at regular interval
func startExecutorHealthCheck(e *healthExecution, activeExecutorMap *activeExecutorMap, id string) {
	executor, _ := activeExecutorMap.Get(id)
	ctx := context.Background()
	log.Info(fmt.Sprintf("session ID: %v, opening connection with executor: %s", executor.sessionID, id))
	for {
		select {
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

// UpdatePingStatus updates the executor alive status 
func (e *healthExecution) UpdatePingStatus(ctx context.Context, request *protofiles.Ping, pingTimeOut time.Duration) (*protofiles.HealthResponse, error) {
	executorID := request.ID
	// if executor is already active
	if executor, ok := e.activeExecutorMap.Get(executorID); ok {
		executor.pingChan <- struct{}{}
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
	pingChan := make(chan struct{})
	sessionID, err := idgen.NewRandomIdGenerator().Generate()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	timer := time.NewTimer(pingTimeOut)
	newActiveExecutor := activeExecutor{
		sessionID: sessionID,
		timer:     timer,
		pingChan:  pingChan,
	}
	e.activeExecutorMap.Put(executorID, &newActiveExecutor)
	go startExecutorHealthCheck(e, e.activeExecutorMap, executorID)
	return &protofiles.HealthResponse{Recieved: true}, nil
}

func getActiveExecutorMap(e *healthExecution) *activeExecutorMap {
	return e.activeExecutorMap
}
