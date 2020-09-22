package job

import (
	"octavius/internal/pkg/protofiles"

	"github.com/stretchr/testify/mock"
)

type MockGrpcClient struct {
	mock.Mock
}

// ConnectClient mock
func (m *MockGrpcClient) ConnectClient(cpHost string) error {
	args := m.Called(cpHost)
	return args.Error(0)
}

func (m *MockGrpcClient) Logs(requestForLog *protofiles.RequestToGetLogs) (*protofiles.Log, error) {
	args := m.Called(requestForLog)
	return args.Get(0).(*protofiles.Log), args.Error(1)
}

func (m *MockGrpcClient) Execute(requestForExecute *protofiles.RequestToExecute) (*protofiles.Response, error) {
	args := m.Called(requestForExecute)
	return args.Get(0).(*protofiles.Response), args.Error(1)
}

func (m *MockGrpcClient) List(requestForGetJobList *protofiles.RequestToGetJobList) (*protofiles.JobList, error) {
	args := m.Called(requestForGetJobList)
	return args.Get(0).(*protofiles.JobList), args.Error(1)
}
