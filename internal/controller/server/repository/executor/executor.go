package executor

import (
	"context"
	"fmt"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/db/etcd"
	"octavius/internal/pkg/log"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"octavius/internal/pkg/util"

	"google.golang.org/protobuf/proto"
)

const (
	executorJobPrefix = "executor/job"
)

//Repository interface for functions related to metadata repository
type Repository interface {
	Save(ctx context.Context, key string, executorInfo *executorCPproto.ExecutorInfo) (*executorCPproto.RegisterResponse, error)
	Get(ctx context.Context, key string) (*executorCPproto.ExecutorInfo, error)
	UpdateStatus(ctx context.Context, key string, health string) error
	AddJob(ctx context.Context, key string, job *executorCPproto.Job, state string) (bool, error)
	RemoveJob(ctx context.Context, key string, state string) (bool, error)
	GetNextJob(ctx context.Context, key string, state string) (*executorCPproto.Job, error)
}

type executorRepository struct {
	etcdClient etcd.Client
}

//NewExecutorRepository initializes metadataRepository with the given etcdClient
func NewExecutorRepository(client etcd.Client) Repository {
	return &executorRepository{
		etcdClient: client,
	}
}

func (e *executorRepository) Save(ctx context.Context, key string, executorInfo *executorCPproto.ExecutorInfo) (*executorCPproto.RegisterResponse, error) {
	dbKey := constant.ExecutorRegistrationPrefix + key

	val, err := proto.Marshal(executorInfo)
	if err != nil {
		return &executorCPproto.RegisterResponse{}, err
	}

	err = e.etcdClient.PutValue(ctx, dbKey, string(val))
	if err != nil {
		return &executorCPproto.RegisterResponse{}, err
	}

	log.Info(fmt.Sprintf("request ID: %v, saved executor %s info to etcd with value %v", ctx.Value(util.ContextKeyUUID), key, executorInfo))
	return &executorCPproto.RegisterResponse{Registered: true}, nil
}

func (e *executorRepository) UpdateStatus(ctx context.Context, key string, health string) error {
	dbKey := constant.ExecutorStatusPrefix + key
	return e.etcdClient.PutValue(ctx, dbKey, health)
}

func (e *executorRepository) Get(ctx context.Context, key string) (*executorCPproto.ExecutorInfo, error) {
	dbKey := constant.ExecutorRegistrationPrefix + key

	infoString, err := e.etcdClient.GetValue(ctx, dbKey)
	if err != nil {
		return nil, err
	}
	executor := &executorCPproto.ExecutorInfo{}

	err = proto.Unmarshal([]byte(infoString), executor)
	return executor, err
}

func (e *executorRepository) AddJob(ctx context.Context, executorID string, job *executorCPproto.Job, state string) (bool, error) {
	dbKey := executorJobPrefix + state + executorID

	val, err := proto.Marshal(job)
	if err != nil {
		return false, err
	}
	err = e.etcdClient.PutValue(ctx, dbKey, string(val))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e *executorRepository) RemoveJob(ctx context.Context, executorID string, state string) (bool, error) {
	dbKey := executorJobPrefix + state + executorID

	return e.etcdClient.DeleteKey(ctx, dbKey)
}

func (e *executorRepository) GetNextJob(ctx context.Context, executorID string, state string) (*executorCPproto.Job, error) {
	//TODO: to be implemented
	return nil, nil
}
