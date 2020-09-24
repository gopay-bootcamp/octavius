package job

import (
	"octavius/internal/pkg/protofiles"
	"time"

	"github.com/stretchr/testify/mock"
)

//MockGrpcClient is the mock of the job services Grpc client
type MockGrpcClient struct {
	mock.Mock
}

//ConnectClient mocks the connect client method
func (m *MockGrpcClient) ConnectClient(cpHost string, connectionTimeOut time.Duration) error {
	args := m.Called(cpHost, connectionTimeOut)
	return args.Error(0)
}

//FetchJob mocks the fetch job method
func (m *MockGrpcClient) FetchJob(start *protofiles.ExecutorID) (*protofiles.Job, error) {
	args := m.Called(start)
	return args.Get(0).(*protofiles.Job), args.Error(1)
}

//SendExecutionContext mocks the send executionContext method
func (m *MockGrpcClient) SendExecutionContext(executionData *protofiles.ExecutionContext) (*protofiles.Acknowledgement, error) {
	args := m.Called(executionData)
	return args.Get(0).(*protofiles.Acknowledgement), args.Error(1)
}

//PostExecutorStatus mocks the post executorStatus method
func (m *MockGrpcClient) PostExecutorStatus(stat *protofiles.Status) (*protofiles.Acknowledgement, error) {
	args := m.Called(stat)
	return args.Get(0).(*protofiles.Acknowledgement), args.Error(1)
}
