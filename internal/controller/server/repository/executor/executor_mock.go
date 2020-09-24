// Package executor implements executor repository related functions
package executor

import (
	"context"
	"octavius/internal/pkg/protofiles"

	"github.com/stretchr/testify/mock"
)

// ExecutorMock mock executor repository
type ExecutorMock struct {
	mock.Mock
}

// SaveExecutorInfo mock that takes key and executor as args
func (m *ExecutorMock) SaveExecutorInfo(ctx context.Context, key string, executorInfo *protofiles.ExecutorInfo) (*protofiles.RegisterResponse, error) {
	args := m.Called(key, executorInfo)
	return args.Get(0).(*protofiles.RegisterResponse), args.Error(1)
}

// GetExecutorInfo mocks the GetExecutorInfo functionality of repository
func (m *ExecutorMock) GetExecutorInfo(ctx context.Context, key string) (*protofiles.ExecutorInfo, error) {
	args := m.Called(key)
	return args.Get(0).(*protofiles.ExecutorInfo), args.Error(1)
}

//UpdateExecutorHealth mocks UpdateExecutorHealth functionality of repository
func (m *ExecutorMock) UpdateExecutorHealth(ctx context.Context, key string, value string) error {
	args := m.Called(key, value)
	return args.Error(0)
}
