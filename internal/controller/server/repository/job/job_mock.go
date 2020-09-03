package job

import (
	"context"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"

	"github.com/stretchr/testify/mock"
)

type JobMock struct {
	mock.Mock
}

func (m *JobMock) Save(ctx context.Context, jobID uint64, executionContext *clientCPproto.RequestForExecute) error {
	args := m.Called(jobID, executionContext)
	return args.Error(0)
}

func (m *JobMock) Delete(ctx context.Context, key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func (m *JobMock) FetchNextJob(ctx context.Context) (string, *clientCPproto.RequestForExecute, error) {
	args := m.Called()
	return args.String(0), args.Get(1).(*clientCPproto.RequestForExecute), args.Error(2)
}

func (m *JobMock) CheckJobIsAvailable(ctx context.Context, jobName string) (bool, error) {
	args := m.Called(jobName)
	return args.Bool(0), args.Error(1)
}
