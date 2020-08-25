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
	SaveMetadataToDb(ctx context.Context, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error)
	ReadAllMetadata(ctx context.Context) (*clientCPproto.MetadataArray, error)
	RegisterExecutor(ctx context.Context, request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error)
	executorRegistered(ctx context.Context, id string) (bool, error)
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

//SaveMetadataToDb calls the repository/metadata Save() function and returns MetadataName
func (e *execution) SaveMetadataToDb(ctx context.Context, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error) {
	return e.metadataRepo.Save(ctx, metadata.Name, metadata)
}

//ReadAllMetadata calls the repository/metadata GetAll() and returns MetadataArray
func (e *execution) ReadAllMetadata(ctx context.Context) (*clientCPproto.MetadataArray, error) {
	return e.metadataRepo.GetAll(ctx)
}

func (e *execution) RegisterExecutor(ctx context.Context, request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error) {
	key := request.ID
	return e.executorRepo.Save(ctx, key, request)
}

func (e *execution) StartHealthCheck(ctx context.Context, activeExecutorMap *sync.Map, id string, healthChan chan string) {
	ticker := time.NewTicker(time.Second)
	expiredStatus := make(chan bool)
	deadline := 200
	presentTime := 0

	for {
		select {
		case health := <-healthChan:
			err := e.executorRepo.UpdateExecutorStatus(ctx, id, health)
			if err != nil {
				logger.Error(errors.New(fmt.Sprintf("error in updating status for executor with %s id", id)), "")
				expiredStatus <- true
			}
			presentTime = 0
		case expired := <-expiredStatus:
			if expired {
				err := e.executorRepo.UpdateExecutorStatus(ctx, id, "expired")
				if err != nil {
					logger.Error(errors.New(fmt.Sprintf("error in updating status for executor with %s id", id)), "")
				}
				logger.Info(fmt.Sprintf("deadline for executor with %s id expired", id))
				ticker.Stop()
				activeExecutorMap.Delete(id)
				close(healthChan)
			}
		case <-ticker.C:
			presentTime++
			if presentTime > deadline {
				expiredStatus <- true
			}
			expiredStatus <- false
		}
	}
}

func (e *execution) executorRegistered(ctx context.Context, id string) (bool, error) {
	return e.executorRepo.CheckIfPresent(ctx, id)
}

func (e *execution) UpdateExecutorStatus(ctx context.Context, request *executorCPproto.Ping) (*executorCPproto.HealthResponse, error) {
	executorID := request.ID
	if channel, ok := e.activeExecutorMap.Load(executorID); ok {
		channel.(chan string) <- request.State
		return &executorCPproto.HealthResponse{Recieved: true}, nil
	}

	registered, err := e.executorRegistered(ctx, request.ID)
	if err != nil {
		logger.Error(err, "error in registering executor")
		return nil, err
	}
	if !registered {
		return &executorCPproto.HealthResponse{Recieved: true}, errors.New("executor not registered. Please register this executor first!")
	}

	healthChan := make(chan string)
	e.activeExecutorMap.Store(executorID, healthChan)
	go e.StartHealthCheck(ctx, e.activeExecutorMap, executorID, healthChan)
	return &executorCPproto.HealthResponse{Recieved: true}, nil
}
