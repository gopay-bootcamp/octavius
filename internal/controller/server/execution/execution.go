package execution

import (
	"context"
	"errors"
	"fmt"
	"octavius/internal/controller/config"
	executorRepo "octavius/internal/controller/server/repository/executor"
	metadataRepo "octavius/internal/controller/server/repository/metadata"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/log"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"octavius/internal/pkg/util"
	"sync"
	"time"
)

// Execution interface for methods related to execution
type Execution interface {
	SaveMetadata(ctx context.Context, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error)
	ReadAllMetadata(ctx context.Context) (*clientCPproto.MetadataArray, error)
	RegisterExecutor(ctx context.Context, request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error)
	UpdateExecutorStatus(ctx context.Context, request *executorCPproto.Ping) (*executorCPproto.HealthResponse, error)
	StartExecutorHealthCheck(activeExecutorMap *sync.Map, id string, healthChan chan string, requestID uint64)
}

type execution struct {
	metadataRepo      metadataRepo.Repository
	executorRepo      executorRepo.Repository
	activeExecutorMap *sync.Map
}

// NewExec creates a new instance of metadata respository
func NewExec(metadataRepo metadataRepo.Repository, executorRepo executorRepo.Repository) Execution {
	return &execution{
		metadataRepo:      metadataRepo,
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

//StartExecutionHealthCheck checks for executor ping at regular interval
func (e *execution) StartExecutorHealthCheck(activeExecutorMap *sync.Map, id string, healthChan chan string, requestID uint64) {
	ctx := context.Background()
	timer := time.NewTimer(config.Config().ExecutorPingDeadline)
	cleanUpChan := make(chan struct{})
	log.Info(fmt.Sprintf("request ID: %v, opening connection with executor: %s", requestID, id))
	err := e.executorRepo.UpdateStatus(ctx, id, "free")
	if err != nil {
		log.Error(err, "")
		cleanUpChan <- struct{}{}
	}
	for {
		select {
		case health := <-healthChan:
			err := e.executorRepo.UpdateStatus(ctx, id, health)
			if err != nil {
				log.Error(err, "")
				cleanUpChan <- struct{}{}
			}
			timer.Stop()
			timer.Reset(config.Config().ExecutorPingDeadline)
		case <-timer.C:
			err := e.executorRepo.UpdateStatus(ctx, id, "expired")
			if err != nil {
				log.Error(err, "")
				cleanUpChan <- struct{}{}
			}
			log.Info(fmt.Sprintf("request ID: %v, deadline exceeded for executor with %s id", requestID, id))
			cleanUpChan <- struct{}{}
		case <-cleanUpChan:
			log.Info(fmt.Sprintf("request ID: %v, closing connection with executor: %s", requestID, id))
			activeExecutorMap.Delete(id)
			close(healthChan)
			timer.Stop()
			return
		}
	}
}

func (e *execution) UpdateExecutorStatus(ctx context.Context, request *executorCPproto.Ping) (*executorCPproto.HealthResponse, error) {
	executorID := request.ID
	requestID := ctx.Value(util.ContextKeyUUID)
	// construct to load channel if executor present in memory map
	if channel, ok := e.activeExecutorMap.Load(executorID); ok {
		channel.(chan string) <- request.State
		return &executorCPproto.HealthResponse{Recieved: true}, nil
	}

	_, err := e.executorRepo.Get(ctx, request.ID)
	if err != nil {
		if err.Error() == constant.NoValueFound {
			return nil, errors.New("executor not registered")
		}
		return nil, err
	}

	// construct to make a new channel and add the executor to the in memory map
	healthChan := make(chan string)

	e.activeExecutorMap.Store(executorID, healthChan)
	go e.StartExecutorHealthCheck(e.activeExecutorMap, executorID, healthChan, uint64(requestID.(uint64)))
	return &executorCPproto.HealthResponse{Recieved: true}, nil
}
