package metadata

import (
	"context"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"

	"github.com/stretchr/testify/mock"
)

type MetadataMock struct {
	mock.Mock
}

func (m *MetadataMock) GetValue(ctx context.Context, jobName string) (*clientCPproto.Metadata, error) {
	args := m.Called(jobName)
	return args.Get(0).(*clientCPproto.Metadata), args.Error(1)
}

// Save mock that takes key and metadata as args
func (m *MetadataMock) Save(ctx context.Context, key string, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error) {
	args := m.Called(key, metadata)
	return args.Get(0).(*clientCPproto.MetadataName), args.Error(1)
}

// GetAll mock that takes no args
func (m *MetadataMock) GetAll(ctx context.Context) (*clientCPproto.MetadataArray, error) {
	args := m.Called()
	return args.Get(0).(*clientCPproto.MetadataArray), args.Error(1)
}

func (m *MetadataMock) GetAvailableJobList(ctx context.Context) (*clientCPproto.JobList, error) {
	args := m.Called()
	return args.Get(0).(*clientCPproto.JobList), args.Error(1)
}
