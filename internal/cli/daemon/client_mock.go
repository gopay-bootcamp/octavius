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

func (m *MockClient) StartClient() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockClient) CreateMetadata(metadataFileHandler io.Reader, grpcClient client.Client) (*protobuf.Response, error) {
	args := m.Called(metadataFileHandler)
	return args.Get(0).(*protobuf.Response), args.Error(1)
}
