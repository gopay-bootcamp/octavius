package daemon

import (
	"octavius/internal/cli/client"
	"octavius/internal/cli/config"
	"octavius/pkg/protobuf"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateMetadata(t *testing.T) {

	mockGrpcClient := client.MockGrpcClient{}
	mockConfigLoader := config.MockLoader{}
	testClient := NewClient(&mockConfigLoader)

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

	mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	mockGrpcClient.On("ConnectClient", "localhost:5050").Return(nil).Once()
	mockGrpcClient.On("CreateJob", &testPostRequest).Return(&testPostResponse, nil).Once()
	res, err := testClient.CreateMetadata(metadataTestFileHandler, &mockGrpcClient)

	assert.Nil(t, err)
	assert.Equal(t, &testPostResponse, res)
	mockGrpcClient.AssertExpectations(t)
	mockConfigLoader.AssertExpectations(t)
}
