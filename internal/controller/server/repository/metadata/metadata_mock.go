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

// GetValue mocks GetValue functionality of repository
func (m *MetadataMock) GetValue(ctx context.Context, jobName string) (*protofiles.Metadata, error) {
	args := m.Called(jobName)
	return args.Get(0).(*protofiles.Metadata), args.Error(1)
}

// Save mocks Save functionality of repository
func (m *MetadataMock) Save(ctx context.Context, key string, metadata *protofiles.Metadata) (*protofiles.MetadataName, error) {
	args := m.Called(key, metadata)
	return args.Get(0).(*protofiles.MetadataName), args.Error(1)
}

// GetAll mocks GetAll functionality of repository
func (m *MetadataMock) GetAll(ctx context.Context) (*protofiles.MetadataArray, error) {
	args := m.Called()
	return args.Get(0).(*protofiles.MetadataArray), args.Error(1)
}

// GetAllKeys mocks GetAllKeys functionality of repository
func (m *MetadataMock) GetAllKeys(ctx context.Context) (*protofiles.JobList, error) {
	args := m.Called()
	return args.Get(0).(*protofiles.JobList), args.Error(1)
}
