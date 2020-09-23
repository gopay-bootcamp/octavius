// Package executor implements executor repository related functions
package executor

import (
	"context"
	"fmt"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/db/etcd"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/protofiles"
	"octavius/internal/pkg/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Repository interface for functions related to executor repository
type Repository interface {
	Save(ctx context.Context, key string, executorInfo *protofiles.ExecutorInfo) (*protofiles.RegisterResponse, error)
	Get(ctx context.Context, key string) (*protofiles.ExecutorInfo, error)
	UpdateStatus(ctx context.Context, key string, health string) error
}

type executorRepository struct {
	etcdClient etcd.Client
}

// NewExecutorRepository initializes metadataRepository with the given etcdClient
func NewExecutorRepository(client etcd.Client) Repository {
	return &executorRepository{
		etcdClient: client,
	}
}

// Save takes exexcutorInfo and key as arguments and saves it to executor/register
func (e *executorRepository) Save(ctx context.Context, key string, executorInfo *protofiles.ExecutorInfo) (*protofiles.RegisterResponse, error) {
	dbKey := constant.ExecutorRegistrationPrefix + key

	val, err := proto.Marshal(executorInfo)
	if err != nil {
		return &protofiles.RegisterResponse{}, status.Error(codes.Internal, err.Error())
	}

	err = e.etcdClient.PutValue(ctx, dbKey, string(val))
	if err != nil {
		return &protofiles.RegisterResponse{}, status.Error(codes.Internal, err.Error())
	}

	log.Info(fmt.Sprintf("request ID: %v, saved executor %s with value %v", ctx.Value(util.ContextKeyUUID), key, executorInfo))
	return &protofiles.RegisterResponse{Registered: true}, nil
}

// UpdateStatus takes executor key and health as arguments and updates status of executor in executor/status
func (e *executorRepository) UpdateStatus(ctx context.Context, key string, health string) error {
	dbKey := constant.ExecutorStatusPrefix + key
	return e.etcdClient.PutValue(ctx, dbKey, health)
}

// Get takes executor key as an argument and returns information about that particular executor
func (e *executorRepository) Get(ctx context.Context, key string) (*protofiles.ExecutorInfo, error) {
	dbKey := constant.ExecutorRegistrationPrefix + key

	infoString, err := e.etcdClient.GetValue(ctx, dbKey)
	if err != nil {
		if err.Error() == constant.NoValueFound {
			return nil, status.Error(codes.NotFound, constant.Etcd+constant.NoValueFound)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	executor := &protofiles.ExecutorInfo{}

	err = proto.Unmarshal([]byte(infoString), executor)
	if err != nil {
		return executor, status.Error(codes.Internal, err.Error())
	}
	return executor, nil
}
