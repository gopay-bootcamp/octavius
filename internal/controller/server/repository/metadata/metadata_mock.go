// Package metadata implements metadata repository related functions
package metadata

import (
	"context"
	"octavius/internal/pkg/protofiles"

	"github.com/stretchr/testify/mock"
)

// MetadataMock mocks metadata repository
type MetadataMock struct {
	mock.Mock
}

// GetMetadata mocks GetMetadata functionality of repository
func (m *MetadataMock) GetMetadata(ctx context.Context, jobName string) (*protofiles.Metadata, error) {
	args := m.Called(jobName)
	return args.Get(0).(*protofiles.Metadata), args.Error(1)
}

// SaveMetadata mocks SaveMetadata functionality of repository
func (m *MetadataMock) SaveMetadata(ctx context.Context, key string, metadata *protofiles.Metadata) (*protofiles.MetadataName, error) {
	args := m.Called(key, metadata)
	return args.Get(0).(*protofiles.MetadataName), args.Error(1)
}

// GetAvailableJobs mocks GetAvailableJobs functionality of repository
func (m *MetadataMock) GetAvailableJobs(ctx context.Context) (*protofiles.JobList, error) {
	args := m.Called()
	return args.Get(0).(*protofiles.JobList), args.Error(1)
}
