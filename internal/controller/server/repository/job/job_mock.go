package job

import (
	"context"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"

	"github.com/stretchr/testify/mock"
)

type JobMock struct {
	mock.Mock
}

func (m *JobMock) Save(ctx context.Context, jobID uint64, executionData *clientCPproto.RequestForExecute) error {
	args := m.Called(jobID, executionData)
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

func (m *JobMock) ValidateJob(ctx context.Context, executionData *clientCPproto.RequestForExecute) (bool, error) {
	args := m.Called(executionData)
	return args.Bool(0), args.Error(1)
}
func (m *JobMock) SaveJobExecutionData(ctx context.Context, jobID string, executionData *executorCPproto.ExecutionContext) error {
	args := m.Called(jobID, executionData)
	return args.Error(0)
}
func (m *JobMock) GetLogs(ctx context.Context, jobID string) (string, error){
	args := m.Called(jobID) 
	return args.String(0),args.Error(1)
}