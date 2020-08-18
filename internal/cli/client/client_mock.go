package client

import (
	"octavius/pkg/protobuf"

	"github.com/stretchr/testify/mock"
)

type MockGrpcClient struct {
	mock.Mock
}

func (m *MockGrpcClient) CreateJob(metadataPostRequest *protobuf.RequestForMetadataPost) (*protobuf.Response, error) {
	args := m.Called(metadataPostRequest)
	return args.Get(0).(*protobuf.Response), args.Error(1)
}
