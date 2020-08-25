package executor

import (
	"context"
	"octavius/internal/control_plane/db/etcd"
	executorCPproto "octavius/internal/pkg/protofiles/executor_CP"

	"google.golang.org/protobuf/proto"
)

const (
	registerPrefix = "executor/register/"
)

//MetadataRepository interface for functions related to metadata repository
type ExecutorRepository interface {
	Save(ctx context.Context, key string, executorInfo *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error)
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
		return &executorCPproto.RegisterResponse{Registered: false}, err
	}
	err = e.etcdClient.PutValue(ctx, dbKey, string(val))
	if err != nil {
		return &executorCPproto.RegisterResponse{Registered: false}, err
	}
	return &executorCPproto.RegisterResponse{Registered: true}, nil
}
