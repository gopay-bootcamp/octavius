package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type JobExecutorMock struct {
	mock.Mock
}

func (m *JobExecutorMock) ExecuteJob(ctx context.Context, jobName string, jobData map[string]string) (uint64,error) {
	args := m.Called(jobName, jobData)
	return args.Get(0).(uint64), args.Error(1)
}