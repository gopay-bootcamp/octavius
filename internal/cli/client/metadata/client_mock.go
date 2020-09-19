package metadata

import (
	protobuf "octavius/internal/pkg/protofiles"

	"github.com/stretchr/testify/mock"
)

type MockGrpcClient struct {
	mock.Mock
}

// Describe mock
func (m *MockGrpcClient) Describe(requestForDescribe *protobuf.RequestToDescribe) (*protobuf.Metadata, error) {
	args := m.Called(requestForDescribe)
	return args.Get(0).(*protobuf.Metadata), args.Error(1)
}

// CreateMetadata mock
func (m *MockGrpcClient) Post(metadataPostRequest *protobuf.RequestToPostMetadata) (*protobuf.MetadataName, error) {
	args := m.Called(metadataPostRequest)
	return args.Get(0).(*protobuf.MetadataName), args.Error(1)
}

// ConnectClient mock
func (m *MockGrpcClient) ConnectClient(cpHost string) error {
	args := m.Called(cpHost)
	return args.Error(0)
}
