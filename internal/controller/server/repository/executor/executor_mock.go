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

// Save mock that takes key and executor as args
func (m *ExecutorMock) Save(ctx context.Context, key string, executorInfo *protofiles.ExecutorInfo) (*protofiles.RegisterResponse, error) {
	args := m.Called(key, executorInfo)
	return args.Get(0).(*protofiles.RegisterResponse), args.Error(1)
}

// Get mocks the get functionality of repository
func (m *ExecutorMock) Get(ctx context.Context, key string) (*protofiles.ExecutorInfo, error) {
	args := m.Called(key)
	return args.Get(0).(*protofiles.ExecutorInfo), args.Error(1)
}

//Update mocks update functionality of repository
func (m *ExecutorMock) Update(ctx context.Context, key string, value string) error {
	args := m.Called(key, value)
	return args.Error(0)
}
