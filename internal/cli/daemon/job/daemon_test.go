package job

import (
	"errors"
	"octavius/internal/cli/client/job"
	"octavius/internal/cli/config"
	"octavius/internal/pkg/protofiles"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	mockGrpcClient := job.MockGrpcClient{}
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
	var jobData = map[string]string{
		"Namespace": "default",
	}
	testExecuteRequest := protofiles.RequestToExecute{
		ClientInfo: &testRequestHeader,
		JobName:    "DemoJob",
		JobData:    jobData,
	}
	executedResponse := protofiles.Response{
		Status: "success",
	}
	mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	mockGrpcClient.On("ConnectClient", "localhost:5050").Return(nil).Once()
	mockGrpcClient.On("Execute", &testExecuteRequest).Return(&executedResponse, nil).Once()
	res, err := testClient.Execute("DemoJob", jobData, &mockGrpcClient)

	assert.Nil(t, err)
	assert.Equal(t, &executedResponse, res)
	mockGrpcClient.AssertExpectations(t)
	mockConfigLoader.AssertExpectations(t)
}

func TestExecuteForError(t *testing.T) {
	mockGrpcClient := job.MockGrpcClient{}
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
	var jobData = map[string]string{
		"Namespace": "default",
	}
	testExecuteRequest := protofiles.RequestToExecute{
		ClientInfo: &testRequestHeader,
		JobName:    "DemoJob",
		JobData:    jobData,
	}
	executedResponse := &protofiles.Response{
		Status: "success",
	}
	mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	mockGrpcClient.On("ConnectClient", "localhost:5050").Return(errors.New("some error")).Once()
	mockGrpcClient.On("ExecuteJob", &testExecuteRequest).Return(executedResponse, nil).Once()
	_, err := testClient.Execute("DemoJob", jobData, &mockGrpcClient)

	assert.Error(t, err)
}

func TestLogs(t *testing.T) {
	mockGrpcClient := job.MockGrpcClient{}
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
	testGetLogRequest := protofiles.RequestToGetLogs{
		ClientInfo: &testRequestHeader,
		JobName:    "DemoJob",
	}
	logResponse := protofiles.Log{
		Log: "sample log 1",
	}
	mockConfigLoader.On("Load").Return(testConfig, config.ConfigError{}).Once()
	mockGrpcClient.On("ConnectClient", "localhost:5050").Return(nil).Once()
	mockGrpcClient.On("Logs", &testGetLogRequest).Return(&logResponse, nil).Once()
	res, err := testClient.Logs("DemoJob", &mockGrpcClient)

	assert.Nil(t, err)
	assert.Equal(t, &logResponse, res)
	mockGrpcClient.AssertExpectations(t)
	mockConfigLoader.AssertExpectations(t)
}
