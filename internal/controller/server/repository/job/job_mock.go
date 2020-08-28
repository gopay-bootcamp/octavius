package job

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type JobExecutorMock struct {
	mock.Mock
}

func (m *JobExecutorMock) ExecuteJob(ctx context.Context,jobIdString string, jobName string, jobData map[string]string) (error) {
	args := m.Called(jobIdString, jobName, jobData)
	return  args.Error(0)
}