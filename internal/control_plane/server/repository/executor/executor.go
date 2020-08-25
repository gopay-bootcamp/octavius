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

//MetadataRepository interface for functions related to metadata repository
type ExecutorRepository interface {
	Save(ctx context.Context, key string, executorInfo *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error)
	CheckIfPresent(ctx context.Context, key string) (bool, error)
	UpdateExecutorStatus(ctx context.Context, key string, health string) error
}

type executorRepository struct {
	etcdClient etcd.EtcdClient
}

//NewMetadataRepository initializes metadataRepository with the given etcdClient
func NewExecutorRepository(client etcd.EtcdClient) ExecutorRepository {
	return &executorRepository{
		etcdClient: client,
	}
}

func (e *executorRepository) Save(ctx context.Context, key string, register *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error) {
	dbKey := registerPrefix + key
	val, err := proto.Marshal(register)
	if err != nil {
		return &executorCPproto.RegisterResponse{}, err
	}
	err = e.etcdClient.PutValue(ctx, dbKey, string(val))
	if err != nil {
		return &executorCPproto.RegisterResponse{}, err
	}
	return &executorCPproto.RegisterResponse{Registered: true}, nil
}

func (e *executorRepository) UpdateExecutorStatus(ctx context.Context, key string, health string) error {
	dbKey := statusPrefix + key
	return e.etcdClient.PutValue(ctx, dbKey, health)
}

func (e *executorRepository) CheckIfPresent(ctx context.Context, key string) (bool, error) {
	dbKey := registerPrefix + key
	gr, err := e.etcdClient.GetValue(ctx, dbKey)
	if err != nil {
		if err.Error() != "no value found" {
			return true, err
		}
	}
	if gr != "" {
		return true, nil
	}
	return false, nil
}
