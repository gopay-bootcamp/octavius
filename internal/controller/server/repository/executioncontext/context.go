package executioncontext

import (
	"context"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/db/etcd"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

const (
	contextPrefix = "context/"
)

type Repository interface {
	Save(ctx context.Context, key string, executionContext *executorCPproto.ExecutionContext) (*executorCPproto.Acknowledgement, error)
	Get(ctx context.Context, key string) (*executorCPproto.ExecutionContext, error)
}

type contextRepository struct {
	etcdClient etcd.Client
}

func NewContextRepository(client etcd.Client) Repository {
	return &contextRepository{
		etcdClient: client,
	}
}

func (c *contextRepository) Save(ctx context.Context, key string, executionContext *executorCPproto.ExecutionContext) (*executorCPproto.Acknowledgement, error) {
	dbKey := contextPrefix + key

	val, err := proto.Marshal(executionContext)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = c.etcdClient.PutValue(ctx, dbKey, string(val))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &executorCPproto.Acknowledgement{Recieved: true}, nil
}

func (c *contextRepository) Get(ctx context.Context, key string) (*executorCPproto.ExecutionContext, error) {
	dbKey := contextPrefix + key

	infoString, err := c.etcdClient.GetValue(ctx, dbKey)
	if err != nil {
		if err.Error() == constant.NoValueFound {
			return nil, status.Error(codes.NotFound, constant.Etcd+constant.NoValueFound)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	executor := &executorCPproto.ExecutionContext{}

	err = proto.Unmarshal([]byte(infoString), executor)
	if err != nil {
		return executor, status.Error(codes.Internal, err.Error())
	}
	return executor, nil
}
