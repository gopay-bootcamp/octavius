// Package job implements job repository related functions
package job

import (
	"context"
	"octavius/internal/pkg/protofiles"

	"github.com/stretchr/testify/mock"
)

// JobMock mocks job repository
type JobMock struct {
	mock.Mock
}

// Save mocks Save functionality of repository
func (m *JobMock) Save(ctx context.Context, jobID uint64, executionData *protofiles.RequestToExecute) error {
	args := m.Called(jobID, executionData)
	return args.Error(0)
}

// DeleteJob mocks DeleteJob functionality of repository
func (m *JobMock) DeleteJob(ctx context.Context, key string) error {
	args := m.Called(key)
	return args.Error(0)
}

// GetNextJob mocks GetNextJob functionality of repository
func (m *JobMock) GetNextJob(ctx context.Context) (string, *protofiles.RequestToExecute, error) {
	args := m.Called()
	return args.String(0), args.Get(1).(*protofiles.RequestToExecute), args.Error(2)
}

// SaveJobExecutionData mocks SaveJobExecutionData functionality of repository
func (m *JobMock) SaveJobExecutionData(ctx context.Context, jobID string, executionData *protofiles.ExecutionContext) error {
	args := m.Called(jobID, executionData)
	return args.Error(0)
}

// GetLogs mocks GetLogs functionality of repository
func (m *JobMock) GetLogs(ctx context.Context, jobID string) (string, error) {
	args := m.Called(jobID)
	return args.String(0), args.Error(1)
}

// GetMetadata mocks GetMetadata functionality of repository
func (m *JobMock) GetMetadata(ctx context.Context, jobName string) (*protofiles.Metadata, error) {
	args := m.Called(jobName)
	return args.Get(0).(*protofiles.Metadata), args.Error(1)
}

//UpdateStatus mocks UpdateStatus functionality of repository
func (m *JobMock) UpdateStatus(ctx context.Context, key string, health string) error {
	args := m.Called(key, health)
	return args.Error(1)
}
