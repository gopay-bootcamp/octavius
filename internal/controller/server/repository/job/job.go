package job

import (
	"context"
	"octavius/internal/pkg/db/etcd"
	"strings"
)

type JobExecutionRepository interface {
	ExecuteJob(context.Context, string, string, map[string]string) error
	CheckJobMetadataIsAvailable(context.Context,string) (bool, error)
}

type jobExecutionRepository struct {
	etcdClient etcd.Client
}

//NewJobExecutionRepository initializes jobExecutionRepository with the given etcdClient and scheduler
func NewJobExecutionRepository(client etcd.Client) JobExecutionRepository {
	return &jobExecutionRepository{
		etcdClient: client,
	}
}

func storeEnvVariablesInDatabase(ctx context.Context, etcdClient etcd.Client, jobId string, jobData map[string]string) error {

	for envName, value := range jobData {
		key := "jobs/" + jobId + "/env/" + envName
		err := etcdClient.PutValue(ctx, key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (j jobExecutionRepository) ExecuteJob(ctx context.Context, jobIdString string, jobName string, jobData map[string]string) error {
	key := "jobs/" + jobIdString + "/metadataKeyName"
	value := "metadata/" + jobName
	err := j.etcdClient.PutValue(ctx, key, value)
	if err != nil {
		return err
	}

	err = storeEnvVariablesInDatabase(ctx, j.etcdClient, jobIdString, jobData)
	if err != nil {
		return err
	}

	return nil
}

func (j jobExecutionRepository) CheckJobMetadataIsAvailable(ctx context.Context, jobName string) (bool, error) {
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
