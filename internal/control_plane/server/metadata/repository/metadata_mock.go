package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"octavius/pkg/protobuf"
)

type MetadataMock struct {
	mock.Mock
}

func (m *MetadataMock) Save(ctx context.Context, key string, metadata *protobuf.Metadata) *protobuf.MetadataName {
	args := m.Called(key, metadata)
	return args.Get(0).(*protobuf.MetadataName)
}

func (m *MetadataMock) GetAll(ctx context.Context) *protobuf.MetadataArray {
	args := m.Called()
	return args.Get(0).(*protobuf.MetadataArray)
}
