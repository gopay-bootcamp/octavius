package health

import (
	"errors"
	"io/ioutil"
	"octavius/internal/executor/client"
	"octavius/internal/executor/config"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/kubernetes"
	"octavius/internal/pkg/log"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	log.Init("info", "", false, 1)
}

func TestStartClient(t *testing.T) {
	mockGrpcClient := new(client.MockGrpcClient)

	testClient := NewExecutorClient(mockGrpcClient)
	testExecutorClient := testClient.(*executorClient)

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

	testKubeClient, _ := kubernetes.NewKubernetesClient(testConfig)

	finalExecutorClient := &executorClient{
		grpcClient:            mockGrpcClient,
		cpHost:                "test host",
		id:                    "test id",
		accessToken:           "test access",
		connectionTimeoutSecs: time.Second,
		pingInterval:          time.Second,
		kubeLogWaitTime:       5 * time.Minute,
		kubernetesClient:      testKubeClient,
		state:                 constant.IdleState,
	}

	mockGrpcClient.On("ConnectClient", "test host", time.Second).Return(nil)

	testExecutorClient.StartClient(testConfig)
	assert.Equal(t, finalExecutorClient, testExecutorClient)
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
	mockGrpcClient := new(client.MockGrpcClient)
	testClient := NewExecutorClient(mockGrpcClient)
	testExecutorClient := testClient.(*executorClient)
	testExecutorClient.id = "test id"
	mockGrpcClient.On("ConnectClient", "test host", time.Second).Return(nil)
	mockGrpcClient.On("Ping", &executorCPproto.Ping{ID: "test id", State: "idle"}).Return(&executorCPproto.HealthResponse{Recieved: true}, nil)
	testExecutorClient.StartClient(testConfig)
	testExecutorClient.StartPing()
	time.Sleep(6 * time.Second)
	mockGrpcClient.AssertExpectations(t)
}
