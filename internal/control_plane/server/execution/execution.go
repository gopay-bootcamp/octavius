package execution

import (
	"context"
	executorRepo "octavius/internal/control_plane/server/repository/executor"
	metadataRepo "octavius/internal/control_plane/server/repository/metadata"
	"octavius/internal/control_plane/server/timer"
	clientCPproto "octavius/internal/pkg/protofiles/client_CP"
	executorCPproto "octavius/internal/pkg/protofiles/executor_CP"
)

// Execution interface for methods related to execution
type Execution interface {
	SaveMetadataToDb(ctx context.Context, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error)
	ReadAllMetadata(ctx context.Context) (*clientCPproto.MetadataArray, error)
	RegisterExecutor(ctx context.Context, request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error)
}

type execution struct {
	metadataRepo      metadataRepo.MetadataRepository
	executorRepo      executorRepo.ExecutorRepository
	activeExecutorMap map[string](chan int64)
}

// NewExec creates a new instance of metadata respository
func NewExec(metadataRepo metadataRepo.MetadataRepository, executorRepo executorRepo.ExecutorRepository) Execution {
	return &execution{
		metadataRepo: metadataRepo,
		executorRepo: executorRepo,
	}
}

//SaveMetadataToDb calls the repository/metadata Save() function and returns MetadataName
func (e *execution) SaveMetadataToDb(ctx context.Context, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error) {
	result, err := e.metadataRepo.Save(ctx, metadata.Name, metadata)
	return result, err
}

//ReadAllMetadata calls the repository/metadata GetAll() and returns MetadataArray
func (e *execution) ReadAllMetadata(ctx context.Context) (*clientCPproto.MetadataArray, error) {
	result, err := e.metadataRepo.GetAll(ctx)
	return result, err
}

func (e *execution) RegisterExecutor(ctx context.Context, request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error) {
	key := request.ID
	res, err := e.executorRepo.Save(ctx, key, request)
	return res, err
}

func StartHealthCheck(activeExecutorMap map[string](chan int64), healthChan chan int64) {
	//start a timer

	select:
		timerChan: //deadline cross
		heathChan:	//when ping is recieved

	//if timer expires
		//close channel
		//delete from map
}

func (e *execution) UpdateExecutorStatus(ctx context.Context, request *executorCPproto.Ping) (*executorCPproto.HealthResponse, error) {
	executorID := request.ID
	if _, ok := e.activeExecutorMap[executorID]; ok {
		e.activeExecutorMap[executorID] <- request.Health
		return &executorCPproto.HealthResponse{Recieved: true}, nil
	}
	//if the executor id is present in the etcd create a health chan and pass it to the checking routine
	if notPresentInEtcd {
		return &executorCPproto.HealthResponse{Recieved: false}, nil
	}
	if inPresentInEtcd {
		healthChan := make(chan int64)
		e.activeExecutorMap[executorID] = healthChan
		go StartHealthCheck(e.activeExecutorMap, healthChan)
	}
	timerInstance := timer.GetTimer()
	// map of channels
	// initialize timer in channel
	// reset timer in channel
	//

	//check if the executor is registered
	//timer
	//channel
	//if timer expires update status as dead
	//update etcd with present ping status
}
