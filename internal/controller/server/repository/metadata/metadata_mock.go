package metadata

import (
	"context"
	clientCPproto "octavius/internal/pkg/protofiles/client_CP"

	"github.com/stretchr/testify/mock"
)

type MetadataMock struct {
	mock.Mock
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
