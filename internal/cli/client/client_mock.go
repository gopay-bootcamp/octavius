package client

import (
	"octavius/pkg/protobuf"

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
