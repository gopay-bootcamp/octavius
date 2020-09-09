package client

import (
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockGrpcClient struct {
	mock.Mock
}

//Register mock executor
func (m *MockGrpcClient) Register(request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error) {
	args := m.Called(request)
	return args.Get(0).(*executorCPproto.RegisterResponse), args.Error(1)
}

func (m *MockGrpcClient) Ping(ping *executorCPproto.Ping) (*executorCPproto.HealthResponse, error) {
	args := m.Called(ping)
	return args.Get(0).(*executorCPproto.HealthResponse), args.Error(1)
}

func (m *MockGrpcClient) ConnectClient(cpHost string, connectionTimeOut time.Duration) error {
	args := m.Called(cpHost, connectionTimeOut)
	return args.Error(0)
}

func (m *MockGrpcClient) FetchJob(start *executorCPproto.ExecutorID) (*executorCPproto.Job, error) {
	args := m.Called(start)
	return args.Get(0).(*executorCPproto.Job), args.Error(1)
}

func (m *MockGrpcClient) SendExecutionContext(executionData *executorCPproto.ExecutionContext) (*executorCPproto.Acknowledgement, error) {
	args := m.Called(executionData)
	return args.Get(0).(*executorCPproto.Acknowledgement), args.Error(1)
}
