package executor

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	executor_cp "octavius/internal/pkg/protofiles/executor_cp"
)

type ExecutorMock struct {
	mock.Mock
}

func (e ExecutorMock) Save(ctx context.Context, key string, executorInfo *executor_cp.ExecutorInfo) (*executor_cp.RegisterResponse, error) {
	return nil, errors.New("not implemented yet")

}

func (e ExecutorMock) Get(ctx context.Context, key string) (*executor_cp.ExecutorInfo, error) {
	return nil, errors.New("not implemented yet")
}

func (e ExecutorMock) UpdateStatus(ctx context.Context, key string, health string) error {
	return errors.New("not implemented yet")
}
