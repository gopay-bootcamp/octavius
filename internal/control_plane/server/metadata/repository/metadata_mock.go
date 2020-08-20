package repository

import (
	"context"
	"octavius/pkg/protobuf"

	"github.com/stretchr/testify/mock"
)

type MetadataMock struct {
	mock.Mock
}

// Save mock that takes key and metadata as args
func (m *MetadataMock) Save(ctx context.Context, key string, metadata *protobuf.Metadata) (*protobuf.MetadataName,error) {
	args := m.Called(key, metadata)
	return args.Get(0).(*protobuf.MetadataName),args.Error(1)
}

// GetAll mock that takes no args
func (m *MetadataMock) GetAll(ctx context.Context) (*protobuf.MetadataArray,error) {
	args := m.Called()
	return args.Get(0).(*protobuf.MetadataArray),args.Error(1)
}
