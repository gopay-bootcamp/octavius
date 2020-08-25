package daemon

import (
	"io"
	"octavius/internal/cli/client"
	clientCPproto "octavius/internal/pkg/protofiles/client_CP"

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
func (m *MockClient) CreateMetadata(metadataFileHandler io.Reader, grpcClient client.Client) (*clientCPproto.MetadataName, error) {
	args := m.Called(metadataFileHandler)
	return args.Get(0).(*clientCPproto.MetadataName), args.Error(1)
}

func (m *MockClient) GetStreamLog(jobName string, grpcClient client.Client) (*[]clientCPproto.Log, error) {
	args := m.Called(jobName)
	return args.Get(0).(*[]clientCPproto.Log), args.Error(1)
}

func (m *MockClient) ExecuteJob(jobName string, jobData map[string]string, grpcClient client.Client) (*clientCPproto.Response, error) {
	args := m.Called(jobName, jobData)
	return args.Get(0).(*clientCPproto.Response), args.Error(1)
}
