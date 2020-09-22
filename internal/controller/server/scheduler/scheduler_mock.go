package scheduler

import (
	"context"
	"octavius/internal/pkg/protofiles"

	"github.com/stretchr/testify/mock"
)

type SchedulerMock struct {
	mock.Mock
}

func (m *SchedulerMock) AddToPendingList(ctx context.Context, jobID uint64, executionData *protofiles.RequestToExecute) error {
	args := m.Called(jobID, executionData)
	return args.Error(0)
}

func (m *SchedulerMock) RemoveFromPendingList(ctx context.Context, key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func (m *SchedulerMock) FetchJob(ctx context.Context) (string, *protofiles.RequestToExecute, error) {
	args := m.Called()
	return args.String(0), args.Get(1).(*protofiles.RequestToExecute), args.Error(2)
}
