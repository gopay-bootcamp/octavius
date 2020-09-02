package job
import (
	"context"
	"errors"
	"github.com/gogo/protobuf/proto"
	"octavius/internal/pkg/db/etcd"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
	"strconv"
	"strings"
)
type Repository interface {
	CheckJobMetadataIsAvailable(ctx context.Context,jobName string) (bool, error)
	Save(ctx context.Context, jobID uint64, jobContext *clientCPproto.RequestForExecute) error
	Delete(ctx context.Context, key string) error
	FetchNextJob(ctx context.Context) (string, *clientCPproto.RequestForExecute, error)
}
type jobRepository struct {
	etcdClient etcd.Client
}

const (
	pendingPrefix = "jobs/pending/"

)


//NewJobExecutionRepository initializes jobExecutionRepository with the given etcdClient and scheduler
func NewJobRepository(client etcd.Client) Repository {
	return &jobRepository{
		etcdClient: client,
	}
}
func (j jobRepository) CheckJobMetadataIsAvailable(ctx context.Context, jobName string) (bool, error) {
	jobNameListWithPrefix, _, err := j.etcdClient.GetAllKeyAndValues(ctx, "metadata/")
	if err != nil {
		return false, err
	}
	for _, jobNameWithPrefix := range jobNameListWithPrefix {
		availableJobName := strings.Split(jobNameWithPrefix, "/")[1]
		if availableJobName == jobName {
			return true, nil
		}
	}
	return false, nil
}

func (j jobRepository) Save(ctx context.Context, jobID uint64, jobContext *clientCPproto.RequestForExecute) error {
	jobIDasString := strconv.FormatUint(jobID, 10)
	key := "jobs/pending/" + jobIDasString
	value, err := proto.Marshal(jobContext)
	if err != nil {
		return  err
	}
	valueAsString := string(value)
	return j.etcdClient.PutValue(ctx, key, valueAsString)
}
func (j jobRepository) Delete(ctx context.Context, key string) error {
	_, err := j.etcdClient.DeleteKey(ctx, pendingPrefix+key)
	return err
}
func (j jobRepository) FetchNextJob(ctx context.Context) (string, *clientCPproto.RequestForExecute, error) {
	keys, values, err := j.etcdClient.GetAllKeyAndValues(ctx, pendingPrefix)
	if err != nil {
		return "", nil,  err
	}
	if len(values) == 0 {
		return "", nil , errors.New("no pending job in pending job list")
	}
	nextJobID := strings.Split(keys[0],"/")[2]

	var nextJobContext *clientCPproto.RequestForExecute
	err = proto.Unmarshal([]byte(values[0]), nextJobContext)
	if err != nil {
		return "",nil, errors.New("error in unmarshalling job context")
	}
	return nextJobID, nextJobContext, nil
}