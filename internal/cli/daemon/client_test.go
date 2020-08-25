package daemon

import (
	"octavius/internal/cli/client"
	"octavius/internal/cli/config"
	protobuf "octavius/internal/pkg/protofiles/client_CP"
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

	testPostRequest := protobuf.RequestToPostMetadata{
		Metadata:   &testMetadata,
		ClientInfo: &testRequestHeader,
	}

	testPostMetadataName := protobuf.MetadataName{
		Name: "success",
	}

	mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	mockGrpcClient.On("ConnectClient", "localhost:5050").Return(nil).Once()
	mockGrpcClient.On("CreateMetadata", &testPostRequest).Return(&testPostMetadataName, nil).Once()
	res, err := testClient.CreateMetadata(metadataTestFileHandler, &mockGrpcClient)

	assert.Nil(t, err)
	assert.Equal(t, &testPostMetadataName, res)
	mockGrpcClient.AssertExpectations(t)
	mockConfigLoader.AssertExpectations(t)
}

func TestExecuteJob(t *testing.T) {
	mockGrpcClient := client.MockGrpcClient{}
	mockConfigLoader := config.MockLoader{}
	testClient := NewClient(&mockConfigLoader)

	testConfig := config.OctaviusConfig{
		Host:                  "localhost:5050",
		Email:                 "akshay.busa@go-jek.com",
		AccessToken:           "AllowMe",
		ConnectionTimeoutSecs: time.Second,
	}
	testRequestHeader := protobuf.ClientInfo{
		ClientEmail: "akshay.busa@go-jek.com",
		AccessToken: "AllowMe",
	}
	var jobData = map[string]string{
		"Namespace": "default",
	}
	testExecuteRequest := protobuf.RequestForExecute{
		ClientInfo: &testRequestHeader,
		JobName:    "DemoJob",
		JobData:    jobData,
	}
	executedResponse := protobuf.Response{
		Status: "success",
	}
	mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	mockGrpcClient.On("ConnectClient", "localhost:5050").Return(nil).Once()
	mockGrpcClient.On("ExecuteJob", &testExecuteRequest).Return(&executedResponse, nil).Once()
	res, err := testClient.ExecuteJob("DemoJob", jobData, &mockGrpcClient)

	assert.Nil(t, err)
	assert.Equal(t, &executedResponse, res)
	mockGrpcClient.AssertExpectations(t)
	mockConfigLoader.AssertExpectations(t)

}

func TestGetStream(t *testing.T) {
	mockGrpcClient := client.MockGrpcClient{}
	mockConfigLoader := config.MockLoader{}
	testClient := NewClient(&mockConfigLoader)

	testConfig := config.OctaviusConfig{
		Host:                  "localhost:5050",
		Email:                 "akshay.busa@go-jek.com",
		AccessToken:           "AllowMe",
		ConnectionTimeoutSecs: time.Second,
	}
	testRequestHeader := protobuf.ClientInfo{
		ClientEmail: "akshay.busa@go-jek.com",
		AccessToken: "AllowMe",
	}
	testGetStreamRequest := protobuf.RequestForStreamLog{
		ClientInfo: &testRequestHeader,
		JobName:    "DemoJob",
	}
	var logResponse []protobuf.Log
	log1 := &protobuf.Log{
		Log: "Test log 1",
	}
	log2 := &protobuf.Log{
		Log: "Test log 2",
	}
	logResponse = append(logResponse, *log1)
	logResponse = append(logResponse, *log2)

	mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	mockGrpcClient.On("ConnectClient", "localhost:5050").Return(nil).Once()
	mockGrpcClient.On("GetStreamLog", &testGetStreamRequest).Return(&logResponse, nil).Once()
	res, err := testClient.GetStreamLog("DemoJob", &mockGrpcClient)

	assert.Nil(t, err)
	assert.Equal(t, &logResponse, res)
	mockGrpcClient.AssertExpectations(t)
	mockConfigLoader.AssertExpectations(t)
}
