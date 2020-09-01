package executor

import (
	"context"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"

	"github.com/stretchr/testify/mock"
)

//ExecutorMock mock executor repository
type ExecutorMock struct {
	mock.Mock
}

// Save mock that takes key and executor as args
func (m *ExecutorMock) Save(ctx context.Context, key string, executorInfo *executorCPproto.ExecutorInfo) (*executorCPproto.RegisterResponse, error) {
	args := m.Called(key, executorInfo)
	return args.Get(0).(*executorCPproto.RegisterResponse), args.Error(1)
}

// Get mocks the get fucntionality of repository
func (m *ExecutorMock) Get(ctx context.Context, key string) (*executorCPproto.ExecutorInfo, error) {
	args := m.Called(key)
	return args.Get(0).(*executorCPproto.ExecutorInfo), args.Error(1)
}

//UpdateStatus mocks update status functionality of repository
func (m *ExecutorMock) UpdateStatus(ctx context.Context, key string, health string) error {
	args := m.Called(key, health)
	return args.Error(0)
}
