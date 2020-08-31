package executor

import (
	"context"
	"github.com/stretchr/testify/mock"
	executor_cp "octavius/internal/pkg/protofiles/executor_cp"
)

type ExecutorMock struct {
	mock.Mock
}

func (e ExecutorMock) Save(ctx context.Context, key string, executorInfo *executor_cp.ExecutorInfo) (*executor_cp.RegisterResponse, error) {
	panic("implement me")
}

func (e ExecutorMock) Get(ctx context.Context, key string) (*executor_cp.ExecutorInfo, error) {
	panic("implement me")
}

func (e ExecutorMock) UpdateStatus(ctx context.Context, key string, health string) error {
	panic("implement me")
}
