package execution

import (
	"context"
	"errors"
	"fmt"
	"octavius/internal/control_plane/logger"
	executorRepo "octavius/internal/control_plane/server/repository/executor"
	metadataRepo "octavius/internal/control_plane/server/repository/metadata"
	clientCPproto "octavius/internal/pkg/protofiles/client_CP"
	executorCPproto "octavius/internal/pkg/protofiles/executor_CP"
	"sync"
	"time"
)

// Execution interface for methods related to execution
type Execution interface {
	SaveMetadata(ctx context.Context, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error)
	ReadAllMetadata(ctx context.Context) (*clientCPproto.MetadataArray, error)
	RegisterExecutor(ctx context.Context, request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error)
	UpdateExecutorStatus(ctx context.Context, request *executorCPproto.Ping) (*executorCPproto.HealthResponse, error)
}

type execution struct {
	metadataRepo      metadataRepo.MetadataRepository
	executorRepo      executorRepo.ExecutorRepository
	activeExecutorMap *sync.Map
}

// NewExec creates a new instance of metadata respository
func NewExec(metadataRepo metadataRepo.MetadataRepository, executorRepo executorRepo.ExecutorRepository) Execution {
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

func (e *execution) StartHealthCheck(ctx context.Context, activeExecutorMap *sync.Map, id string, healthChan chan string) {
	timer := time.NewTimer(time.Minute)
	deadline := timer.C
	cleanUpChan := make(chan struct{})
	err := e.executorRepo.UpdateStatus(ctx, id, "free")
	if err != nil {
		logger.Error(fmt.Errorf("error in updating status for executor with %s id", id), "")
		cleanUpChan <- struct{}{}
	}
	for {
		select {
		case health := <-healthChan:
			err := e.executorRepo.UpdateStatus(ctx, id, health)
			if err != nil {
				logger.Error(fmt.Errorf("error in updating status for executor with %s id", id), "")
				cleanUpChan <- struct{}{}
			}
			timer.Stop()
			timer.Reset(time.Minute)
		case <-deadline:
			err := e.executorRepo.UpdateStatus(ctx, id, "expired")
			if err != nil {
				logger.Error(fmt.Errorf("error in updating status for executor with %s id", id), "")
				cleanUpChan <- struct{}{}
			}
			logger.Info(fmt.Sprintf("deadline for executor with %s id expired", id))
			cleanUpChan <- struct{}{}
		case <-cleanUpChan:
			activeExecutorMap.Delete(id)
			close(healthChan)
			timer.Stop()
			return
		}
	}
}

func (e *execution) UpdateExecutorStatus(ctx context.Context, request *executorCPproto.Ping) (*executorCPproto.HealthResponse, error) {
	executorID := request.ID
	if channel, ok := e.activeExecutorMap.Load(executorID); ok {
		channel.(chan string) <- request.State
		return &executorCPproto.HealthResponse{Recieved: true}, nil
	}

	_, err := e.executorRepo.Get(ctx, request.ID)
	if err != nil {
		if err.Error() == "no value found" {
			return &executorCPproto.HealthResponse{Recieved: true}, errors.New("executor not registered. Please register this executor first!")
		}
		logger.Error(err, "error in getting executor from repo")
		return nil, err
	}

	healthChan := make(chan string)
	e.activeExecutorMap.Store(executorID, healthChan)
	go e.StartHealthCheck(ctx, e.activeExecutorMap, executorID, healthChan)
	return &executorCPproto.HealthResponse{Recieved: true}, nil
}
