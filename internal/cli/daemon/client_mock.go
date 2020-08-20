package daemon

import (
	"io"
	"octavius/internal/cli/client"
	"octavius/pkg/protobuf"

	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

// StartClient mock
func (m *MockClient) StartClient() error {
	args := m.Called()
	return args.Error(0)
}

// CreateMetadata mock
func (m *MockClient) CreateMetadata(metadataFileHandler io.Reader, grpcClient client.Client) (*protobuf.MetadataName, error) {
	args := m.Called(metadataFileHandler)
	return args.Get(0).(*protobuf.MetadataName), args.Error(1)
}
