package scheduler

import (
	"context"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"

	"github.com/stretchr/testify/mock"
)

type SchedulerMock struct {
	mock.Mock
}

func (m *SchedulerMock) AddToPendingList(ctx context.Context, jobID uint64, executionData *clientCPproto.RequestForExecute) error {
	args := m.Called(jobID, executionData)
	return args.Error(0)
}

func (m *SchedulerMock) RemoveFromPendingList(ctx context.Context, key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func (m *SchedulerMock) FetchJob(ctx context.Context) (string, *clientCPproto.RequestForExecute, error) {
	args := m.Called()
	return args.String(0), args.Get(1).(*clientCPproto.RequestForExecute), args.Error(2)
}
