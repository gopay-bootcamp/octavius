package executor

import (
	"context"
	"octavius/internal/control_plane/db/etcd"
	executorCPproto "octavius/internal/pkg/protofiles/executor_CP"

	"google.golang.org/protobuf/proto"
)

const (
	registerPrefix = "executor/register/"
	statusPrefix   = "executor/status/"
)

//ExecutorRepository interface for functions related to metadata repository
type ExecutorRepository interface {
	Save(ctx context.Context, key string, executorInfo *executorCPproto.ExecutorInfo) (*executorCPproto.RegisterResponse, error)
	Get(ctx context.Context, key string) (*executorCPproto.ExecutorInfo, error)
	UpdateStatus(ctx context.Context, key string, health string) error
}

type executorRepository struct {
	etcdClient etcd.EtcdClient
}

//NewExecutorRepository initializes metadataRepository with the given etcdClient
func NewExecutorRepository(client etcd.EtcdClient) ExecutorRepository {
	return &executorRepository{
		etcdClient: client,
	}
}

func (e *executorRepository) Save(ctx context.Context, key string, executorInfo *executorCPproto.ExecutorInfo) (*executorCPproto.RegisterResponse, error) {
	dbKey := registerPrefix + key
	val, err := proto.Marshal(executorInfo)
	if err != nil {
		return &executorCPproto.RegisterResponse{}, err
	}
	err = e.etcdClient.PutValue(ctx, dbKey, string(val))
	if err != nil {
		return &executorCPproto.RegisterResponse{}, err
	}
	return &executorCPproto.RegisterResponse{Registered: true}, nil
}

func (e *executorRepository) UpdateStatus(ctx context.Context, key string, health string) error {
	dbKey := statusPrefix + key
	return e.etcdClient.PutValue(ctx, dbKey, health)
}

func (e *executorRepository) Get(ctx context.Context, key string) (*executorCPproto.ExecutorInfo, error) {
	dbKey := registerPrefix + key
	infoString, err := e.etcdClient.GetValue(ctx, dbKey)
	if err != nil {
		return nil, err
	}
	executor := &executorCPproto.ExecutorInfo{}
	err = proto.Unmarshal([]byte(infoString), executor)
	if err != nil {
		return nil, err
	}
	return executor, nil
}
