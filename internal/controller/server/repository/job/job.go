package job

import (
	"context"
	"errors"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/db/etcd"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
)

type Repository interface {
	CheckJobIsAvailable(ctx context.Context, jobName string) (bool, error)
	Save(ctx context.Context, jobID uint64, executionContext *clientCPproto.RequestForExecute) error
	Delete(ctx context.Context, key string) error
	FetchNextJob(ctx context.Context) (string, *clientCPproto.RequestForExecute, error)
}
type jobRepository struct {
	etcdClient etcd.Client
}

const (
	pendingPrefix = "jobs/pending/"
)

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
			return false, errors.New("job with given name not found")
		} else {
			return false, err
		}
	}

	return true, nil
}

// Save takes jobID and executionContext and save it in database as pendingList
func (j jobRepository) Save(ctx context.Context, jobID uint64, executionContext *clientCPproto.RequestForExecute) error {
	jobIDasString := strconv.FormatUint(jobID, 10)
	key := pendingPrefix + jobIDasString
	value, err := proto.Marshal(executionContext)
	if err != nil {
		return err
	}
	valueAsString := string(value)
	return j.etcdClient.PutValue(ctx, key, valueAsString)
}

// Delete function delete the job of given key from pendingList in database
func (j jobRepository) Delete(ctx context.Context, key string) error {
	_, err := j.etcdClient.DeleteKey(ctx, pendingPrefix+key)
	return err
}

// FetchNextJob returns jobID and executionContext from pendingList
func (j jobRepository) FetchNextJob(ctx context.Context) (string, *clientCPproto.RequestForExecute, error) {
	keys, values, err := j.etcdClient.GetAllKeyAndValues(ctx, pendingPrefix)
	if err != nil {
		return "", nil, err
	}
	if len(values) == 0 {
		return "", nil, errors.New("no pending job in pending job list")
	}
	nextJobID := strings.Split(keys[0], "/")[2]

	var nextExecutionContext *clientCPproto.RequestForExecute
	nextExecutionContext = &clientCPproto.RequestForExecute{}
	err = proto.Unmarshal([]byte(values[0]), nextExecutionContext)
	if err != nil {
		return "", nil, errors.New("error in unmarshalling job context")
	}
	return nextJobID, nextExecutionContext, nil
}
