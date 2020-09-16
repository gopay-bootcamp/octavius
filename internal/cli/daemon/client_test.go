package daemon

import (
	"errors"
	"octavius/internal/cli/client"
	"octavius/internal/cli/config"
	protobuf "octavius/internal/pkg/protofiles/client_cp"
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

func TestExecuteJob_ExecuteJobError(t *testing.T) {
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
	executedResponse := &protobuf.Response{
		Status: "success",
	}
	mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	mockGrpcClient.On("ConnectClient", "localhost:5050").Return(errors.New("some error")).Once()
	mockGrpcClient.On("ExecuteJob", &testExecuteRequest).Return(executedResponse, nil).Once()
	_, err := testClient.ExecuteJob("DemoJob", jobData, &mockGrpcClient)

	assert.Error(t, err)
}

func TestGetLogs(t *testing.T) {
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
	testGetLogRequest := protobuf.RequestForLogs{
		ClientInfo: &testRequestHeader,
		JobName:    "DemoJob",
	}
	logResponse := protobuf.Log{
		Log: "sample log 1",
	}
	mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	mockGrpcClient.On("ConnectClient", "localhost:5050").Return(nil).Once()
	mockGrpcClient.On("GetLogs", &testGetLogRequest).Return(&logResponse, nil).Once()
	res, err := testClient.GetLogs("DemoJob", &mockGrpcClient)

	assert.Nil(t, err)
	assert.Equal(t, &logResponse, res)
	mockGrpcClient.AssertExpectations(t)
	mockConfigLoader.AssertExpectations(t)
}

func TestDescribeJob(t *testing.T) {
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
	testDescribeRequest := protobuf.RequestForDescribe{
		ClientInfo: &testRequestHeader,
		JobName:    "DemoJob",
	}
	describeResponse := protobuf.Metadata{Name: "DemoJob"}
	mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	mockGrpcClient.On("ConnectClient", "localhost:5050").Return(nil).Once()
	mockGrpcClient.On("DescribeJob", &testDescribeRequest).Return(&describeResponse, nil).Once()
	res, err := testClient.DescribeJob("DemoJob", &mockGrpcClient)
	assert.Nil(t, err)
	assert.Equal(t, &describeResponse, res)
	mockGrpcClient.AssertExpectations(t)
	mockConfigLoader.AssertExpectations(t)

}

func TestGetJobList(t *testing.T) {

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
	testGetJobListRequest := protobuf.RequestForGetJobList{
		ClientInfo: &testRequestHeader,
	}

	var jobList []string
	jobList = append(jobList, "demo-image-name")
	jobList = append(jobList, "demo-image-name-1")

	response := &protobuf.JobList{
		Jobs: jobList,
	}

	mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	mockGrpcClient.On("ConnectClient", "localhost:5050").Return(nil).Once()
	mockGrpcClient.On("GetJobList", &testGetJobListRequest).Return(response, nil).Once()
	res, err := testClient.GetJobList(&mockGrpcClient)

	assert.Nil(t, err)
	assert.Equal(t, response, res)
	mockGrpcClient.AssertExpectations(t)
	mockConfigLoader.AssertExpectations(t)
}

func TestGetJobListForGetJobListFailure(t *testing.T) {

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
	testGetJobListRequest := protobuf.RequestForGetJobList{
		ClientInfo: &testRequestHeader,
	}

	mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	mockGrpcClient.On("ConnectClient", "localhost:5050").Return(nil).Once()
	mockGrpcClient.On("GetJobList", &testGetJobListRequest).Return(&protobuf.JobList{}, errors.New("error in getJobList function")).Once()
	_, err := testClient.GetJobList(&mockGrpcClient)

	assert.Equal(t, "error in getJobList function", err.Error())

	mockGrpcClient.AssertExpectations(t)
	mockConfigLoader.AssertExpectations(t)
}
