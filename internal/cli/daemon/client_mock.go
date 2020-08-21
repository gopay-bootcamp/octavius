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

func (m *MockClient) CreateMetadata(metadataFileHandler io.Reader, grpcClient client.Client) (*protobuf.MetadataName, error) {
	args := m.Called(metadataFileHandler)
	return args.Get(0).(*protobuf.MetadataName), args.Error(1)
}

func (m *MockClient)  GetStreamLog(jobName string, grpcClient client.Client)  (error) {
	args := m.Called(jobName)
	return args.Error(0)
}

func (m *MockClient) ExecuteJob(jobName string, jobData map[string]string, grpcClient client.Client) error {
	args := m.Called(jobName,jobData)
	return args.Error(0)
}

