package metadata

import (
	"context"
	"octavius/internal/pkg/protofiles"

	"github.com/stretchr/testify/mock"
)

type MetadataMock struct {
	mock.Mock
}

func (m *MetadataMock) GetValue(ctx context.Context, jobName string) (*protofiles.Metadata, error) {
	args := m.Called(jobName)
	return args.Get(0).(*protofiles.Metadata), args.Error(1)
}

// Save mock that takes key and metadata as args
func (m *MetadataMock) Save(ctx context.Context, key string, metadata *protofiles.Metadata) (*protofiles.MetadataName, error) {
	args := m.Called(key, metadata)
	return args.Get(0).(*protofiles.MetadataName), args.Error(1)
}

// GetAll mock that takes no args
func (m *MetadataMock) GetAll(ctx context.Context) (*protofiles.MetadataArray, error) {
	args := m.Called()
	return args.Get(0).(*protofiles.MetadataArray), args.Error(1)
}

func (m *MetadataMock) GetAvailableJobList(ctx context.Context) (*protofiles.JobList, error) {
	args := m.Called()
	return args.Get(0).(*protofiles.JobList), args.Error(1)
}
