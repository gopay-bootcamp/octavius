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

type Repository interface {
	GetValue(ctx context.Context, jobName string) (*protofiles.Metadata, error)
	CheckJobIsAvailable(ctx context.Context, jobName string) (bool, error)
	Save(ctx context.Context, jobID uint64, executionData *protofiles.RequestToExecute) error
	Delete(ctx context.Context, key string) error
	UpdateStatus(ctx context.Context, key string, health string) error
	FetchNextJob(ctx context.Context) (string, *protofiles.RequestToExecute, error)
	ValidateJob(context.Context, *protofiles.RequestToExecute) (bool, error)
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

// CheckJobIsAvailable returns true if given job is available otherwise returns false
func (j *jobRepository) CheckJobIsAvailable(ctx context.Context, jobName string) (bool, error) {
	_, err := j.etcdClient.GetValue(ctx, "metadata/"+jobName)
	if err != nil {
		if err.Error() == constant.NoValueFound {
			return false, status.Error(codes.NotFound, constant.Etcd+fmt.Sprintf("job with %v name not found", jobName))
		}
		return false, status.Error(codes.Internal, err.Error())

	}
	return true, nil
}

func (e *jobRepository) UpdateStatus(ctx context.Context, key string, health string) error {
	dbKey := constant.ExecutorStatusPrefix + key
	return e.etcdClient.PutValue(ctx, dbKey, health)
}

// Save takes jobID and executionData and save it in database as pendingList
func (j *jobRepository) Save(ctx context.Context, jobID uint64, executionData *protofiles.RequestToExecute) error {
	key := constant.JobPendingPrefix + strconv.FormatUint(jobID, 10)
	value, err := proto.Marshal(executionData)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	log.Info(fmt.Sprintf("Request ID: %v, saving executionData to etcd with value %+v", ctx.Value(util.ContextKeyUUID), executionData))
	return j.etcdClient.PutValue(ctx, key, string(value))
}

// Delete function delete the job of given key from pendingList in database
func (j *jobRepository) Delete(ctx context.Context, key string) error {
	_, err := j.etcdClient.DeleteKey(ctx, constant.JobPendingPrefix+key)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

// FetchNextJob returns jobID and executionData from pendingList
func (j *jobRepository) FetchNextJob(ctx context.Context) (string, *protofiles.RequestToExecute, error) {
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

//ValidateJob is used to validate the arguments of job when execution request is received
func (j *jobRepository) ValidateJob(ctx context.Context, executionData *protofiles.RequestToExecute) (bool, error) {
	jobName := executionData.JobName
	jobData := executionData.JobData
	key := constant.MetadataPrefix + jobName
	res, err := j.etcdClient.GetValue(ctx, key)
	if err != nil {
		return false, status.Error(codes.Internal, err.Error())
	}

	metadata := &protofiles.Metadata{}
	err = proto.Unmarshal([]byte(res), metadata)
	if err != nil {
		return false, status.Error(codes.Internal, err.Error())
	}

	args := metadata.EnvVars.Args

	for _, arg := range args {
		if arg.Required {
			if _, ok := jobData[arg.Name]; !ok {
				return false, nil
			}
		}
	}
	for jobKey := range jobData {
		if !isPresentInArgs(jobKey, args) {
			return false, nil
		}
	}
	return true, nil
}

func isPresentInArgs(jobKey string, args []*protofiles.Arg) bool {
	for _, arg := range args {
		if arg.Name == jobKey {
			return true
		}
	}
	return false
}

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

func (j *jobRepository) SaveJobExecutionData(ctx context.Context, jobname string, executionData *protofiles.ExecutionContext) error {
	key := constant.ExecutionDataPrefix + jobname
	value, err := proto.Marshal(executionData)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	log.Info(fmt.Sprintf("Request ID: %v, saving executionData to etcd with value %+v", ctx.Value(util.ContextKeyUUID), executionData))
	return j.etcdClient.PutValue(ctx, key, string(value))
}

func (c *jobRepository) GetValue(ctx context.Context, jobName string) (*protofiles.Metadata, error) {
	dbKey := constant.MetadataPrefix + jobName
	gr, err := c.etcdClient.GetValue(ctx, dbKey)

	if err == errors.New(constant.NoValueFound) {
		return &protofiles.Metadata{}, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		return &protofiles.Metadata{}, status.Error(codes.Internal, err.Error())
	}

	metadata := &protofiles.Metadata{}
	err = proto.Unmarshal([]byte(gr), metadata)
	if err != nil {
		return &protofiles.Metadata{}, status.Error(codes.Internal, err.Error())
	}
	return metadata, nil
}
