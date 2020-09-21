package job

import (
	"octavius/internal/cli/client/job"
	"octavius/internal/pkg/protofiles"

	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

// StartClient mock
func (m *MockClient) StartClient() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockClient) Logs(jobID string, grpcClient job.Client) (*protofiles.Log, error) {
	args := m.Called(jobID)
	return args.Get(0).(*protofiles.Log), args.Error(1)
}

func (m *MockClient) Execute(jobName string, jobData map[string]string, grpcClient job.Client) (*protofiles.Response, error) {
	args := m.Called(jobName, jobData)
	return args.Get(0).(*protofiles.Response), args.Error(1)
}
