package job

import (
	"context"
	"octavius/internal/pkg/protofiles"

	"github.com/stretchr/testify/mock"
)

type JobMock struct {
	mock.Mock
}

func (m *JobMock) Save(ctx context.Context, jobID uint64, executionData *protofiles.RequestToExecute) error {
	args := m.Called(jobID, executionData)
	return args.Error(0)
}

func (m *JobMock) Delete(ctx context.Context, key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func (m *JobMock) FetchNextJob(ctx context.Context) (string, *protofiles.RequestToExecute, error) {
	args := m.Called()
	return args.String(0), args.Get(1).(*protofiles.RequestToExecute), args.Error(2)
}

func (m *JobMock) CheckJobIsAvailable(ctx context.Context, jobName string) (bool, error) {
	args := m.Called(jobName)
	return args.Bool(0), args.Error(1)
}

func (m *JobMock) ValidateJob(ctx context.Context, executionData *protofiles.RequestToExecute) (bool, error) {
	args := m.Called(executionData)
	return args.Bool(0), args.Error(1)
}
func (m *JobMock) SaveJobExecutionData(ctx context.Context, jobID string, executionData *protofiles.ExecutionContext) error {
	args := m.Called(jobID, executionData)
	return args.Error(0)
}
func (m *JobMock) GetLogs(ctx context.Context, jobID string) (string, error) {
	args := m.Called(jobID)
	return args.String(0), args.Error(1)
}
func (m *JobMock) GetValue(ctx context.Context, jobName string) (*protofiles.Metadata, error) {
	args := m.Called(jobName)
	return args.Get(0).(*protofiles.Metadata), args.Error(1)
}
