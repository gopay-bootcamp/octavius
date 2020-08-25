package client

import (
	clientCPproto "octavius/internal/pkg/protofiles/client_CP"

	"github.com/stretchr/testify/mock"
)

type MockGrpcClient struct {
	mock.Mock
}

// CreateMetadata mock
func (m *MockGrpcClient) CreateMetadata(metadataPostRequest *clientCPproto.RequestToPostMetadata) (*clientCPproto.MetadataName, error) {
	args := m.Called(metadataPostRequest)
	return args.Get(0).(*clientCPproto.MetadataName), args.Error(1)
}

// ConnectClient mock
func (m *MockGrpcClient) ConnectClient(cpHost string) error {
	args := m.Called(cpHost)
	return args.Error(0)
}

func (m *MockGrpcClient) GetStreamLog(requestForStreamLog *clientCPproto.RequestForStreamLog) (*[]clientCPproto.Log, error) {
	args := m.Called(requestForStreamLog)
	return args.Get(0).(*[]clientCPproto.Log), args.Error(1)
}

func (m *MockGrpcClient) ExecuteJob(requestForExecute *clientCPproto.RequestForExecute) (*clientCPproto.Response, error) {
	args := m.Called(requestForExecute)
	return args.Get(0).(*clientCPproto.Response), args.Error(1)
}
