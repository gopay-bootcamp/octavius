package job

import (
	"octavius/internal/pkg/protofiles"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockGrpcClient struct {
	mock.Mock
}

//ConnectClient mock
func (m *MockGrpcClient) ConnectClient(cpHost string, connectionTimeOut time.Duration) error {
	args := m.Called(cpHost, connectionTimeOut)
	return args.Error(0)
}

//FetchJob mock
func (m *MockGrpcClient) FetchJob(start *protofiles.ExecutorID) (*protofiles.Job, error) {
	args := m.Called(start)
	return args.Get(0).(*protofiles.Job), args.Error(1)
}

//SendExecutionContext mock
func (m *MockGrpcClient) SendExecutionContext(executionData *protofiles.ExecutionContext) (*protofiles.Acknowledgement, error) {
	args := m.Called(executionData)
	return args.Get(0).(*protofiles.Acknowledgement), args.Error(1)
}
