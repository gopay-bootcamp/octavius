package daemon

import (
	"github.com/stretchr/testify/mock"
	"io"
	"octavius/pkg/protobuf"
)

type MockClient struct {
	mock.Mock
}



func (m *MockClient) StartClient() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockClient) CreateMetadata(metadataFileHandler io.Reader) (*protobuf.Response, error) {
	 m.Called(metadataFileHandler)
	return nil,nil
}
