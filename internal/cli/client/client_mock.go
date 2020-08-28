package client

import (
	protobuf "octavius/internal/pkg/protofiles/client_cp"

	"github.com/stretchr/testify/mock"
)

type MockGrpcClient struct {
	mock.Mock
}

// CreateMetadata mock
func (m *MockGrpcClient) CreateMetadata(metadataPostRequest *protobuf.RequestToPostMetadata) (*protobuf.MetadataName, error) {
	args := m.Called(metadataPostRequest)
	return args.Get(0).(*protobuf.MetadataName), args.Error(1)
}

// ConnectClient mock
func (m *MockGrpcClient) ConnectClient(cpHost string) error {
	args := m.Called(cpHost)
	return args.Error(0)
}

func (m *MockGrpcClient) GetStreamLog(requestForStreamLog *protobuf.RequestForStreamLog) (*[]protobuf.Log, error) {
	args := m.Called(requestForStreamLog)
	return args.Get(0).(*[]protobuf.Log), args.Error(1)
}

func (m *MockGrpcClient) ExecuteJob(requestForExecute *protobuf.RequestForExecute) (*protobuf.Response, error) {
	args := m.Called(requestForExecute)
	return args.Get(0).(*protobuf.Response), args.Error(1)
}
