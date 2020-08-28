package scheduler

import (
	"github.com/stretchr/testify/mock"
)

type SchedulerMock struct {

	mock.Mock
}

func (m *SchedulerMock) AddToPendingList(jobId uint64) error {
	args:= m.Called(jobId)
	return args.Error(0)
}

func (m *SchedulerMock) RemoveFromPendingList(key string) error {
	args:= m.Called(key)
	return args.Error(0)
}

func (m *SchedulerMock) FetchJob() (string, error) {
	args:= m.Called()
	return args.Get(0).(string), args.Error(1)
}
