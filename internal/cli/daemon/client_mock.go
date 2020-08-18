package daemon

import (
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) StartClient() error {
	panic("implement me")
}

func (m *MockClient) CreateMetadata(metadataFile string) error {
	args := m.Called()
	return args.Error(0)
}
