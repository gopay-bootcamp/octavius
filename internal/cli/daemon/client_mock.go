package daemon

import (
	"io"
	"octavius/internal/cli/client"
	protobuf "octavius/internal/pkg/protofiles/client_cp"

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

func (m *MockClient) GetStreamLog(jobName string, grpcClient client.Client) (*[]protobuf.Log, error) {
	args := m.Called(jobName)
	return args.Get(0).(*[]protobuf.Log), args.Error(1)
}

func (m *MockClient) ExecuteJob(jobName string, jobData map[string]string, grpcClient client.Client) (*protobuf.Response, error) {
	args := m.Called(jobName, jobData)
	return args.Get(0).(*protobuf.Response), args.Error(1)
}

func (m *MockClient) GetJobList(c client.Client) (*protobuf.JobList, error) {
	args := m.Called()
	return args.Get(0).(*protobuf.JobList), args.Error(1)
}
