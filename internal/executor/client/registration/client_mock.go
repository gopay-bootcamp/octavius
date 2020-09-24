package registration

import (
	"octavius/internal/pkg/protofiles"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockGrpcClient struct {
	mock.Mock
}

//Register mock
func (m *MockGrpcClient) Register(request *protofiles.RegisterRequest) (*protofiles.RegisterResponse, error) {
	args := m.Called(request)
	return args.Get(0).(*protofiles.RegisterResponse), args.Error(1)
}

//ConnectClient mock
func (m *MockGrpcClient) ConnectClient(cpHost string, connectionTimeOut time.Duration) error {
	args := m.Called(cpHost, connectionTimeOut)
	return args.Error(0)
}
