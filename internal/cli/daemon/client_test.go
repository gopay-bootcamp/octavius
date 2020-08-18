package daemon

import (
	"octavius/internal/cli/client"
	"octavius/internal/cli/config"
	"octavius/pkg/protobuf"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
	testClient       Client
	mockConfigLoader config.MockLoader
	mockGrpcClient   client.MockGrpcClient
}

func (s *ClientTestSuite) SetupTest() {
	s.mockGrpcClient = client.MockGrpcClient{}
	s.mockConfigLoader = config.MockLoader{}

	s.testClient = NewClient(&s.mockConfigLoader)

}

func (s *ClientTestSuite) TestCreateMetadata() {
	t := s.T()
	testConfig := config.OctaviusConfig{
		Host:                  "localhost:5050",
		Email:                 "jaimin.rathod@go-jek.com",
		AccessToken:           "AllowMe",
		ConnectionTimeoutSecs: time.Second,
	}

	metadataTestFileHandler := strings.NewReader(
		`{
			"name": "test-name",
			"image_name": "test-image",
			"author": "test-author",
			"organization": "gopay-systems"
	}`)

	testMetadata := protobuf.Metadata{
		Name:         "test-name",
		ImageName:    "test-image",
		Author:       "test-author",
		Organization: "gopay-systems",
	}

	testRequestHeader := protobuf.ClientInfo{
		ClientEmail: "jaimin.rathod@go-jek.com",
		AccessToken: "AllowMe",
	}

	testPostRequest := protobuf.RequestForMetadataPost{
		Metadata:   &testMetadata,
		ClientInfo: &testRequestHeader,
	}

	testPostResponse := protobuf.Response{
		Status: "success",
	}

	s.mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	s.mockGrpcClient.On("ConnectClient", "localhost:5050").Return(nil).Once()
	s.mockGrpcClient.On("CreateJob", &testPostRequest).Return(&testPostResponse, nil).Once()
	res, err := s.testClient.CreateMetadata(metadataTestFileHandler, &s.mockGrpcClient)

	assert.Nil(t, err)
	assert.Equal(t, &testPostResponse, res)
	s.mockGrpcClient.AssertExpectations(t)
	s.mockConfigLoader.AssertExpectations(t)

}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
