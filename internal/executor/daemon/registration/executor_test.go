package registration

import (
	"octavius/internal/executor/client/registration"
	"octavius/internal/executor/config"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/protofiles"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testConfig = config.OctaviusExecutorConfig{
	CPHost:                       "test host",
	ID:                           "test id",
	AccessToken:                  "test access",
	ConnTimeOutSec:               time.Second,
	PingInterval:                 time.Second,
	KubeConfig:                   "out-of-cluster",
	KubeContext:                  "default",
	DefaultNamespace:             "default",
	KubeServiceAccountName:       "test",
	JobPodAnnotations:            map[string]string{"test pod": "test annotation"},
	KubeJobActiveDeadlineSeconds: 1,
	KubeJobRetries:               1,
	KubeWaitForResourcePollCount: 1,
}

func init() {
	log.Init("info", "", false, 1)
}

func TestRegisterClient(t *testing.T) {
	mockGrpcClient := new(registration.MockGrpcClient)
	testClient := NewRegistrationServicesClient(mockGrpcClient)
	testRegistrationServicesClient := testClient.(*registrationServicesClient)
	testRegistrationServicesClient.id = "test id"
	testRegistrationServicesClient.accessToken = "test access"
	testInfo := &protofiles.ExecutorInfo{
		Info: testRegistrationServicesClient.accessToken,
	}

	request := &protofiles.RegisterRequest{
		ID:           "test id",
		ExecutorInfo: testInfo,
	}

	mockGrpcClient.On("Register", request).Return(&protofiles.RegisterResponse{Registered: true}, nil)
	mockGrpcClient.On("ConnectClient", "test host", time.Second).Return(nil)

	res, err := testRegistrationServicesClient.RegisterClient(testConfig)
	mockGrpcClient.AssertExpectations(t)
	assert.Equal(t, true, res)
	assert.Nil(t, err)
}
