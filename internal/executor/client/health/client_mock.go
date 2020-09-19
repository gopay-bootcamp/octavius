package health

import (
	"octavius/internal/pkg/protofiles"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockGrpcClient struct {
	mock.Mock
}

func (m *MockGrpcClient) Ping(ping *protofiles.Ping) (*protofiles.HealthResponse, error) {
	args := m.Called(ping)
	return args.Get(0).(*protofiles.HealthResponse), args.Error(1)
}

func (m *MockGrpcClient) ConnectClient(cpHost string, connectionTimeOut time.Duration) error {
	args := m.Called(cpHost, connectionTimeOut)
	return args.Error(0)
}
