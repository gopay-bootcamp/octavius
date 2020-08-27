package executor

import (
	"context"
	"fmt"
	"octavius/internal/pkg/db/etcd"
	"octavius/internal/pkg/log"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"octavius/internal/pkg/util"

	"google.golang.org/protobuf/proto"
)

const (
	registerPrefix = "executor/register/"
	statusPrefix   = "executor/status/"
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
	dbKey := registerPrefix + key

	val, err := proto.Marshal(executorInfo)
	if err != nil {
		return &executorCPproto.RegisterResponse{}, err
	}

	err = e.etcdClient.PutValue(ctx, dbKey, string(val))
	if err != nil {
		return &executorCPproto.RegisterResponse{}, err
	}

	log.Info(fmt.Sprintf("request id:%v, saved executor %s info to etcd", util.ContextKeyUUID, key))
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

	log.Info(fmt.Sprintf("request id:%v, recieved executor id:%s info and info value is %s", util.ContextKeyUUID, key, infoString))
	executor := &executorCPproto.ExecutorInfo{}

	err = proto.Unmarshal([]byte(infoString), executor)
	return executor, err
}
