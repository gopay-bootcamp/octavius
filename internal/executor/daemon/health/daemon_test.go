package health

import (
	"octavius/internal/executor/client/health"
	"octavius/internal/executor/config"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/protofiles"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	log.Init("info", "", false, 1)
}

func TestStartClient(t *testing.T) {
	mockGrpcClient := new(health.MockGrpcClient)

	testClient := NewHealthServicesClient(mockGrpcClient)
	testhealthServicesClient := testClient.(*healthServicesClient)

	testConfig := config.OctaviusExecutorConfig{
		CPHost:                       "test host",
		ID:                           "test id",
		AccessToken:                  "test access",
		ConnTimeOutSec:               time.Second,
		PingInterval:                 time.Second,
		KubeConfig:                   "test kube config",
		KubeContext:                  "test context",
		DefaultNamespace:             "default",
		KubeServiceAccountName:       "test",
		JobPodAnnotations:            map[string]string{"test pod": "test annotation"},
		KubeJobActiveDeadlineSeconds: 1,
		KubeJobRetries:               1,
		KubeWaitForResourcePollCount: 1,
	}

	finalhealthServicesClient := &healthServicesClient{
		grpcClient:            mockGrpcClient,
		cpHost:                "test host",
		id:                    "test id",
		accessToken:           "test access",
		connectionTimeoutSecs: time.Second,
		pingInterval:          time.Second,
	}

	mockGrpcClient.On("ConnectClient", "test host", time.Second).Return(nil)

	testhealthServicesClient.connectClient(testConfig)
	assert.Equal(t, finalhealthServicesClient, testhealthServicesClient)
	mockGrpcClient.AssertExpectations(t)
}

func TestStartPing(t *testing.T) {
	testConfig := config.OctaviusExecutorConfig{
		CPHost:                       "test host",
		ID:                           "test id",
		AccessToken:                  "test access",
		ConnTimeOutSec:               time.Second,
		PingInterval:                 time.Second,
		KubeConfig:                   "test kube config",
		KubeContext:                  "test context",
		DefaultNamespace:             "default",
		KubeServiceAccountName:       "test",
		JobPodAnnotations:            map[string]string{"test pod": "test annotation"},
		KubeJobActiveDeadlineSeconds: 1,
		KubeJobRetries:               1,
		KubeWaitForResourcePollCount: 1,
	}
	mockGrpcClient := new(health.MockGrpcClient)
	testClient := NewHealthServicesClient(mockGrpcClient)
	testhealthServicesClient := testClient.(*healthServicesClient)
	testhealthServicesClient.id = "test id"
	mockGrpcClient.On("ConnectClient", "test host", time.Second).Return(nil)
	mockGrpcClient.On("Ping", &protofiles.Ping{ID: "test id"}).Return(&protofiles.HealthResponse{Received: true}, nil)
	testhealthServicesClient.StartPing(testConfig)
	time.Sleep(2 * time.Second)
	mockGrpcClient.AssertExpectations(t)
}
