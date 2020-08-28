package repository

import (
	"context"
	"octavius/internal/controller/server/scheduler"
	"octavius/internal/pkg/db/etcd"
	"octavius/internal/pkg/randomIdGenerator"
	"strconv"
)

type JobExecutionRepository interface {
	ExecuteJob(context.Context, string, map[string]string) (uint64, error)
}

type jobExecutionRepository struct {
	etcdClient etcd.Client
	scheduler  scheduler.Scheduler
	idGenerator randomIdGenerator.RandomIdGenerator
}

//NewJobExecutionRepository initializes jobExecutionRepository with the given etcdClient and scheduler
func NewJobExecutionRepository(client etcd.Client, scheduler scheduler.Scheduler, idGenerator randomIdGenerator.RandomIdGenerator) JobExecutionRepository {
	return &jobExecutionRepository{
		etcdClient: client,
		scheduler:  scheduler,
		idGenerator:idGenerator,
	}
}

func storeEnvVariablesInDatabase(ctx context.Context, etcdClient etcd.Client,jobId string, jobData map[string]string) error {

	for envName, value := range jobData {
		key := "jobs/" + jobId + "/env/" + envName
		err := etcdClient.PutValue(ctx, key, value)
		if err != nil {
			return err
		}
	}
	return nil
}


func (j jobExecutionRepository) ExecuteJob(ctx context.Context, jobName string, jobData map[string]string) (uint64,error) {

	jobId, err := j.idGenerator.Generate()
	if err != nil {
		return uint64(0),err
	}

	err = j.scheduler.AddToPendingList(jobId)
	if err!= nil {
		return uint64(0),err
	}

	jobIdString := strconv.FormatUint(jobId, 10)
	key := "jobs/" + jobIdString + "/metadataKeyName"
	value := "metadata/" + jobName
	err = j.etcdClient.PutValue(ctx, key, value)
	if err != nil {
		return uint64(0),err
	}

	err=storeEnvVariablesInDatabase(ctx,j.etcdClient,jobIdString, jobData)
	if err!= nil {
		return uint64(0),err
	}

	return jobId,nil
}


