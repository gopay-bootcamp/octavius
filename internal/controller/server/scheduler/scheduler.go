package scheduler

import (
	"context"
	"errors"
	"octavius/internal/pkg/db/etcd"
	"octavius/internal/pkg/idgen"
	"strconv"
)

type Scheduler interface {
	AddToPendingList(uint64) error
	FetchJob() (string, error)
	RemoveFromPendingList(string) error
}

type scheduler struct {
	etcdClient  etcd.Client
	idGenerator idgen.RandomIdGenerator
}

func NewScheduler(etcdClient etcd.Client, idGenerator idgen.RandomIdGenerator) Scheduler {
	return &scheduler{
		etcdClient:  etcdClient,
		idGenerator: idGenerator,
	}
}

func (s *scheduler) AddToPendingList(jobId uint64) error {
	jobIdString := strconv.FormatUint(jobId, 10)
	key := "jobs/pending/" + jobIdString

	return s.etcdClient.PutValue(context.Background(), key, jobIdString)

}

func (s *scheduler) RemoveFromPendingList(key string) error {
	_, err := s.etcdClient.DeleteKey(context.Background(), key)
	if err != nil {
		return err
	}
	return nil
}

func (s *scheduler) FetchJob() (string, error) {
	prefix := "jobs/pending/"

	keys, values, err := s.etcdClient.GetAllKeyAndValues(context.Background(), prefix)
	if err != nil {
		return "", err
	}
	if len(values) == 0 {
		return "", errors.New("no pending job in pending job list")
	}
	err = s.RemoveFromPendingList(prefix + keys[0])
	if err != nil {
		return "", err
	}

	return values[0], nil
}
