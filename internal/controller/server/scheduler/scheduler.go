package scheduler

import (
	"context"
	jobRepo "octavius/internal/controller/server/repository/job"
	"octavius/internal/pkg/idgen"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
)

type Scheduler interface {
	AddToPendingList(context.Context, uint64, *clientCPproto.RequestForExecute) error
	FetchJob(ctx context.Context) (string, *clientCPproto.RequestForExecute, error)
	RemoveFromPendingList(context.Context, string) error
}
type scheduler struct {
	idGenerator idgen.RandomIdGenerator
	jobRepo     jobRepo.Repository
}

// NewScheduler initializes new scheduler with randomIdGenerator and jobRepo
func NewScheduler(idGenerator idgen.RandomIdGenerator, schedulerRepo jobRepo.Repository) Scheduler {
	return &scheduler{
		idGenerator: idGenerator,
		jobRepo:     schedulerRepo,
	}
}

// AddToPendingList function add given job to pendingList
func (s *scheduler) AddToPendingList(ctx context.Context, jobId uint64, jobContext *clientCPproto.RequestForExecute) error {

	return s.jobRepo.Save(ctx, jobId, jobContext)
}

// RemoveFromPendigList function removes job with given key from pendingList
func (s *scheduler) RemoveFromPendingList(ctx context.Context, key string) error {
	return s.jobRepo.Delete(ctx, key)
}

// FetchJob returns jobID and jobContext from pendingList
func (s *scheduler) FetchJob(ctx context.Context) (string, *clientCPproto.RequestForExecute, error) {

	nextJobID, nextJobContext, err := s.jobRepo.FetchNextJob(ctx)
	if err != nil {
		return "", nil, err
	}

	err = s.RemoveFromPendingList(ctx, nextJobID)
	if err != nil {
		return "", nil, err
	}

	return nextJobID, nextJobContext, nil
}
