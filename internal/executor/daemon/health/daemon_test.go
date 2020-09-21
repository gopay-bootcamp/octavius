package health

import (
	"octavius/internal/executor/client/health"
	"octavius/internal/executor/config"
	"octavius/internal/pkg/protofiles"
	"testing"
	"time"
)

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
	testHealthServicesClient := NewHealthServicesClient(mockGrpcClient)
	mockGrpcClient.On("ConnectClient", "test host", time.Second).Return(nil)
	mockGrpcClient.On("Ping", &protofiles.Ping{ID: "test id", State: "ping"}).Return(&protofiles.HealthResponse{Recieved: true}, nil)
	testHealthServicesClient.StartPing(testConfig)
	time.Sleep(6 * time.Second)
	mockGrpcClient.AssertExpectations(t)
}
