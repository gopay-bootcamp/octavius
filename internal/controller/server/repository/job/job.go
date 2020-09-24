// Package job implements job repository related functions
package job

import (
	"context"
	"errors"
	"fmt"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/db/etcd"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/protofiles"
	"octavius/internal/pkg/util"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/proto"
)

// Repository interface for job repository functions
type Repository interface {
	GetMetadata(ctx context.Context, jobName string) (*protofiles.Metadata, error)
	SaveJobArgs(ctx context.Context, jobID uint64, executionData *protofiles.RequestToExecute) error
	DeleteJob(ctx context.Context, key string) error
	UpdateStatus(ctx context.Context, key string, health string) error
	GetNextJob(ctx context.Context) (string, *protofiles.RequestToExecute, error)
	GetLogs(context.Context, string) (string, error)
	SaveJobExecutionData(ctx context.Context, jobID string, executionData *protofiles.ExecutionContext) error
}

type jobRepository struct {
	etcdClient etcd.Client
}

// NewJobRepository initializes jobRepository with the given etcdClient
func NewJobRepository(client etcd.Client) Repository {
	return &jobRepository{
		etcdClient: client,
	}
}

func (j *jobRepository) UpdateStatus(ctx context.Context, key string, health string) error {
	dbKey := constant.ExecutorStatusPrefix + key
	return j.etcdClient.PutValue(ctx, dbKey, health)
}

// SaveJobArgs takes jobID and executionData and save it in database as pendingList
func (j *jobRepository) SaveJobArgs(ctx context.Context, jobID uint64, executionData *protofiles.RequestToExecute) error {
	key := constant.JobPendingPrefix + strconv.FormatUint(jobID, 10)
	value, err := proto.Marshal(executionData)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	log.Info(fmt.Sprintf("Request ID: %v, saving executionData to etcd with value %+v", ctx.Value(util.ContextKeyUUID), executionData))
	return j.etcdClient.PutValue(ctx, key, string(value))
}

// DeleteJob function delete the job of given key from pendingList in database
func (j *jobRepository) DeleteJob(ctx context.Context, key string) error {
	_, err := j.etcdClient.DeleteKey(ctx, constant.JobPendingPrefix+key)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

// GetNextJob returns jobID and executionData from pendingList
func (j *jobRepository) GetNextJob(ctx context.Context) (string, *protofiles.RequestToExecute, error) {
	keys, values, err := j.etcdClient.GetAllKeyAndValues(ctx, constant.JobPendingPrefix)
	if err != nil {
		return "", nil, status.Error(codes.Internal, err.Error())
	}
	if len(values) == 0 {
		return "", nil, status.Error(codes.NotFound, constant.Controller+"no pending job")
	}
	nextJobID := strings.Split(keys[0], "/")[2]
	nextExecutionData := &protofiles.RequestToExecute{}
	err = proto.Unmarshal([]byte(values[0]), nextExecutionData)
	if err != nil {
		return "", nil, status.Error(codes.Internal, err.Error())
	}
	return nextJobID, nextExecutionData, nil
}

// GetLogs is used to fetch logs of job executing/executed
func (j *jobRepository) GetLogs(ctx context.Context, jobName string) (string, error) {
	jobKey := constant.ExecutionDataPrefix + constant.KubeOctaviusPrefix + jobName
	res, err := j.etcdClient.GetValue(ctx, jobKey)
	if err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}

	execContext := &protofiles.ExecutionContext{}
	err = proto.Unmarshal([]byte(res), execContext)
	if err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}

	return execContext.Output, nil

}

// SaveJobExecutionData  saves Execution data of given jobName in executor/logs
func (j *jobRepository) SaveJobExecutionData(ctx context.Context, jobname string, executionData *protofiles.ExecutionContext) error {
	key := constant.ExecutionDataPrefix + jobname
	value, err := proto.Marshal(executionData)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	log.Info(fmt.Sprintf("Request ID: %v, saving executionData to etcd with value %+v", ctx.Value(util.ContextKeyUUID), executionData))
	return j.etcdClient.PutValue(ctx, key, string(value))
}

// GetMetadata returns metadata of given jobName
func (j *jobRepository) GetMetadata(ctx context.Context, jobName string) (*protofiles.Metadata, error) {
	dbKey := constant.MetadataPrefix + jobName
	gr, err := j.etcdClient.GetValue(ctx, dbKey)
	if err != nil {
		if err.Error() == errors.New(constant.NoValueFound).Error() {
			return &protofiles.Metadata{}, status.Error(codes.NotFound, err.Error())
		}
		return &protofiles.Metadata{}, status.Error(codes.Internal, err.Error())
	}

	metadata := &protofiles.Metadata{}
	err = proto.Unmarshal([]byte(gr), metadata)
	if err != nil {
		return &protofiles.Metadata{}, status.Error(codes.Internal, err.Error())
	}
	return metadata, nil
}
