package executor

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return &executorCPproto.RegisterResponse{}, status.Error(codes.Internal, err.Error())
	}

	err = e.etcdClient.PutValue(ctx, dbKey, string(val))
	if err != nil {
		return &executorCPproto.RegisterResponse{}, status.Error(codes.Internal, err.Error())
	}

	log.Info(fmt.Sprintf("request ID: %v, saved executor %s with value %v", ctx.Value(util.ContextKeyUUID), key, executorInfo))
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
		if err.Error() == constant.NoValueFound {
			return nil, status.Error(codes.NotFound,constant.Etcd+constant.NoValueFound)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	executor := &executorCPproto.ExecutorInfo{}

	err = proto.Unmarshal([]byte(infoString), executor)
	if err != nil {
		return executor, status.Error(codes.Internal, err.Error())
	}
	return executor, nil
}
