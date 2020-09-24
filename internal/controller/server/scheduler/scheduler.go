// Package scheduler implements scheduling related functions
package scheduler

import (
	"context"
	jobRepo "octavius/internal/controller/server/repository/job"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/protofiles"
	"sync"
)

// Scheduler interface contains all scheduler related functions
type Scheduler interface {
	AddToPendingList(context.Context, uint64, *protofiles.RequestToExecute) error
	FetchJob(ctx context.Context) (string, *protofiles.RequestToExecute, error)
	RemoveFromPendingList(context.Context, string) error
}
type scheduler struct {
	idGenerator idgen.RandomIdGenerator
	jobRepo     jobRepo.Repository
	mutex       sync.Mutex
}

// NewScheduler initializes new scheduler with randomIdGenerator and jobRepo
func NewScheduler(idGenerator idgen.RandomIdGenerator, schedulerRepo jobRepo.Repository) Scheduler {
	return &scheduler{
		idGenerator: idGenerator,
		jobRepo:     schedulerRepo,
	}
}

// AddToPendingList function add given job to pendingList
func (s *scheduler) AddToPendingList(ctx context.Context, jobID uint64, executionData *protofiles.RequestToExecute) error {
	return s.jobRepo.SaveJobArgs(ctx, jobID, executionData)
}

// RemoveFromPendigList function removes job with given key from pendingList
func (s *scheduler) RemoveFromPendingList(ctx context.Context, key string) error {
	return s.jobRepo.DeleteJob(ctx, key)
}

// FetchJob returns jobID and executionData from pendingList
func (s *scheduler) FetchJob(ctx context.Context) (string, *protofiles.RequestToExecute, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	nextJobID, nextExecutionData, err := s.jobRepo.GetNextJob(ctx)
	if err != nil {
		return "", nil, err
	}

	err = s.RemoveFromPendingList(ctx, nextJobID)
	if err != nil {
		return "", nil, err
	}

	return nextJobID, nextExecutionData, nil
}
