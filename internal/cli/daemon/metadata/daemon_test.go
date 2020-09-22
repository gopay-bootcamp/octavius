package metadata

import (
	"errors"
	"octavius/internal/cli/client/metadata"
	"octavius/internal/cli/config"
	"octavius/internal/pkg/protofiles"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPost(t *testing.T) {
	mockGrpcClient := metadata.MockGrpcClient{}
	mockConfigLoader := config.MockLoader{}
	testClient := NewClient(&mockConfigLoader)

	testConfig := config.OctaviusConfig{
		Host:                  "localhost:5050",
		Email:                 "jaimin.rathod@go-jek.com",
		AccessToken:           "AllowMe",
		ConnectionTimeoutSecs: time.Second,
	}
	arg := protofiles.Arg{
		Name:        "test",
		Description: "test description",
		Required:    false,
	}
	var args []*protofiles.Arg
	args = append(args, &arg)
	var testEnvVars = &protofiles.EnvVars{
		Args: args,
	}
	metadataTestFileHandler := strings.NewReader(`{
		"name": "test-name",
		"image_name": "hello-world",
		"author": "Gopay System Automation Team",
		"organization": "gopay-systems",
		"env_vars": {
		"args": [{
			"name": "test",
			"description": "test description",
			"required" : false
			}]
		}
		}
`)

	testMetadata := protofiles.Metadata{
		Name:         "test-name",
		ImageName:    "hello-world",
		Author:       "Gopay System Automation Team",
		Organization: "gopay-systems",
		EnvVars:      testEnvVars,
	}

	testRequestHeader := protofiles.ClientInfo{
		ClientEmail: "jaimin.rathod@go-jek.com",
		AccessToken: "AllowMe",
	}

	testPostRequest := protofiles.RequestToPostMetadata{
		Metadata:   &testMetadata,
		ClientInfo: &testRequestHeader,
	}

	testPostMetadataName := protofiles.MetadataName{
		Name: "success",
	}

	mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	mockGrpcClient.On("ConnectClient", "localhost:5050").Return(nil).Once()
	mockGrpcClient.On("Post", &testPostRequest).Return(&testPostMetadataName, nil).Once()
	res, err := testClient.Post(metadataTestFileHandler, &mockGrpcClient)

	assert.Nil(t, err)
	assert.Equal(t, &testPostMetadataName, res)
	mockGrpcClient.AssertExpectations(t)
	mockConfigLoader.AssertExpectations(t)
}

func TestDescribe(t *testing.T) {
	mockGrpcClient := metadata.MockGrpcClient{}
	mockConfigLoader := config.MockLoader{}
	testClient := NewClient(&mockConfigLoader)

	testConfig := config.OctaviusConfig{
		Host:                  "localhost:5050",
		Email:                 "akshay.busa@go-jek.com",
		AccessToken:           "AllowMe",
		ConnectionTimeoutSecs: time.Second,
	}
	testRequestHeader := protofiles.ClientInfo{
		ClientEmail: "akshay.busa@go-jek.com",
		AccessToken: "AllowMe",
	}
	testDescribeRequest := protofiles.RequestToDescribe{
		ClientInfo: &testRequestHeader,
		JobName:    "DemoJob",
	}
	describeResponse := protofiles.Metadata{Name: "DemoJob"}
	mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	mockGrpcClient.On("ConnectClient", "localhost:5050").Return(nil).Once()
	mockGrpcClient.On("Describe", &testDescribeRequest).Return(&describeResponse, nil).Once()
	res, err := testClient.Describe("DemoJob", &mockGrpcClient)
	assert.Nil(t, err)
	assert.Equal(t, &describeResponse, res)
	mockGrpcClient.AssertExpectations(t)
	mockConfigLoader.AssertExpectations(t)

}

func TestList(t *testing.T) {
	mockGrpcClient := metadata.MockGrpcClient{}
	mockConfigLoader := config.MockLoader{}
	testClient := NewClient(&mockConfigLoader)

	testConfig := config.OctaviusConfig{
		Host:                  "localhost:5050",
		Email:                 "akshay.busa@go-jek.com",
		AccessToken:           "AllowMe",
		ConnectionTimeoutSecs: time.Second,
	}
	testRequestHeader := protofiles.ClientInfo{
		ClientEmail: "akshay.busa@go-jek.com",
		AccessToken: "AllowMe",
	}
	testGetJobListRequest := protofiles.RequestToGetJobList{
		ClientInfo: &testRequestHeader,
	}

	var jobList []string
	jobList = append(jobList, "demo-image-name")
	jobList = append(jobList, "demo-image-name-1")

	response := &protofiles.JobList{
		Jobs: jobList,
	}

	mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	mockGrpcClient.On("ConnectClient", "localhost:5050").Return(nil).Once()
	mockGrpcClient.On("List", &testGetJobListRequest).Return(response, nil).Once()
	res, err := testClient.List(&mockGrpcClient)

	assert.Nil(t, err)
	assert.Equal(t, response, res)
	mockGrpcClient.AssertExpectations(t)
	mockConfigLoader.AssertExpectations(t)
}

func TestListForFailure(t *testing.T) {
	mockGrpcClient := metadata.MockGrpcClient{}
	mockConfigLoader := config.MockLoader{}
	testClient := NewClient(&mockConfigLoader)

	testConfig := config.OctaviusConfig{
		Host:                  "localhost:5050",
		Email:                 "akshay.busa@go-jek.com",
		AccessToken:           "AllowMe",
		ConnectionTimeoutSecs: time.Second,
	}
	testRequestHeader := protofiles.ClientInfo{
		ClientEmail: "akshay.busa@go-jek.com",
		AccessToken: "AllowMe",
	}
	testGetJobListRequest := protofiles.RequestToGetJobList{
		ClientInfo: &testRequestHeader,
	}

	mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	mockGrpcClient.On("ConnectClient", "localhost:5050").Return(nil).Once()
	mockGrpcClient.On("List", &testGetJobListRequest).Return(&protofiles.JobList{}, errors.New("error in getJobList function")).Once()
	_, err := testClient.List(&mockGrpcClient)

	assert.Equal(t, "error in getJobList function", err.Error())

	mockGrpcClient.AssertExpectations(t)
	mockConfigLoader.AssertExpectations(t)
}
