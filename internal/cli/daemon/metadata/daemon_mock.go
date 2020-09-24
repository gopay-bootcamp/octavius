package metadata

import (
	"io"
	"octavius/internal/cli/client/metadata"
	"octavius/internal/pkg/protofiles"

	"github.com/stretchr/testify/mock"
)

//MockClient is a mock for the Daemon
type MockClient struct {
	mock.Mock
}

// StartClient mock
func (m *MockClient) StartClient() error {
	args := m.Called()
	return args.Error(0)
}

//Post mock
func (m *MockClient) Post(metadataFileHandler io.Reader, grpcClient metadata.Client) (*protofiles.MetadataName, error) {
	args := m.Called(metadataFileHandler)
	return args.Get(0).(*protofiles.MetadataName), args.Error(1)
}

//Describe mock
func (m *MockClient) Describe(jobName string, grpcClient metadata.Client) (*protofiles.Metadata, error) {
	args := m.Called(jobName)
	return args.Get(0).(*protofiles.Metadata), args.Error(1)
}

//List mock
func (m *MockClient) List(c metadata.Client) (*protofiles.JobList, error) {
	args := m.Called()
	return args.Get(0).(*protofiles.JobList), args.Error(1)
}
