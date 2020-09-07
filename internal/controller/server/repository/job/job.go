package job

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/db/etcd"
	"octavius/internal/pkg/log"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
	"octavius/internal/pkg/util"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
)

type Repository interface {
	CheckJobIsAvailable(ctx context.Context, jobName string) (bool, error)
	Save(ctx context.Context, jobID uint64, executionData *clientCPproto.RequestForExecute) error
	Delete(ctx context.Context, key string) error
	FetchNextJob(ctx context.Context) (string, *clientCPproto.RequestForExecute, error)
	ValidateJob(context.Context, *clientCPproto.RequestForExecute) (bool, error)
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
func (j jobRepository) CheckJobIsAvailable(ctx context.Context, jobName string) (bool, error) {
	_, err := j.etcdClient.GetValue(ctx, "metadata/"+jobName)
	if err != nil {
		if err.Error() == constant.NoValueFound {
			return false, status.Error(codes.NotFound, constant.Etcd+fmt.Sprintf("job with %v name not found", jobName))
		}
		return false, status.Error(codes.Internal, err.Error())

	}
	return true, nil
}

// Save takes jobID and executionData and save it in database as pendingList
func (j jobRepository) Save(ctx context.Context, jobID uint64, executionData *clientCPproto.RequestForExecute) error {
	key := constant.JobPendingPrefix + strconv.FormatUint(jobID, 10)
	value, err := proto.Marshal(executionData)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	log.Info(fmt.Sprintf("Request ID: %v, saving executionData to etcd with value %+v", ctx.Value(util.ContextKeyUUID), executionData))
	return j.etcdClient.PutValue(ctx, key, string(value))
}

// Delete function delete the job of given key from pendingList in database
func (j jobRepository) Delete(ctx context.Context, key string) error {
	_, err := j.etcdClient.DeleteKey(ctx, constant.JobPendingPrefix+key)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

// FetchNextJob returns jobID and executionData from pendingList
func (j jobRepository) FetchNextJob(ctx context.Context) (string, *clientCPproto.RequestForExecute, error) {
	keys, values, err := j.etcdClient.GetAllKeyAndValues(ctx, constant.JobPendingPrefix)
	if err != nil {
		return "", nil, status.Error(codes.Internal, err.Error())
	}
	if len(values) == 0 {
		return "", nil, status.Error(codes.NotFound, constant.Controller+"no pending job in pending job list")
	}
	nextJobID := strings.Split(keys[0], "/")[2]
	nextExecutionData := &clientCPproto.RequestForExecute{}
	err = proto.Unmarshal([]byte(values[0]), nextExecutionData)
	if err != nil {
		return "", nil, status.Error(codes.Internal, err.Error())
	}
	return nextJobID, nextExecutionData, nil
}

//ValidateJob is used to validate the arguments of job when execution request is received
func (j jobRepository) ValidateJob(ctx context.Context, executionData *clientCPproto.RequestForExecute) (bool, error) {
	jobName := executionData.JobName
	jobData := executionData.JobData
	key := constant.MetadataPrefix + jobName
	res, err := j.etcdClient.GetValue(ctx, key)
	if err != nil {
		return false, status.Error(codes.Internal, err.Error())
	}

	metadata := &clientCPproto.Metadata{}
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

func isPresentInArgs(jobKey string, args []*clientCPproto.Arg) bool {
	for _, arg := range args {
		if arg.Name == jobKey {
			return true
		}
	}
	return false
}
