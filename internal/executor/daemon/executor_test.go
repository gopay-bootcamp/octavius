package daemon

import (
	"octavius/internal/executor/client"
	"octavius/internal/pkg/kubernetes"
	"octavius/internal/pkg/log"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	log.Init("", "", false)
}

func TestRegisterClient(t *testing.T) {
	mockGrpcClient := new(client.MockGrpcClient)
	testClient := NewExecutorClient(mockGrpcClient)
	testExecutorClient := testClient.(*executorClient)
	testExecutorClient.id = "test id"
	testExecutorClient.accessToken = "test token"
	testInfo := &executorCPproto.ExecutorInfo{
		Info: testExecutorClient.accessToken,
	}

	request := &executorCPproto.RegisterRequest{
		ID:           "test id",
		ExecutorInfo: testInfo,
	}

	mockGrpcClient.On("Register", request).Return(&executorCPproto.RegisterResponse{Registered: true}, nil)

	res, err := testExecutorClient.RegisterClient()
	mockGrpcClient.AssertExpectations(t)
	assert.Equal(t, true, res)
	assert.Nil(t, err)
}

func TestStartPing(t *testing.T) {
	mockGrpcClient := new(client.MockGrpcClient)
	testClient := NewExecutorClient(mockGrpcClient)
	testExecutorClient := testClient.(*executorClient)
	testExecutorClient.id = "test id"
	testExecutorClient.state = "test state"
	mockGrpcClient.On("Ping", &executorCPproto.Ping{ID: "test id", State: "test state"}).Return(&executorCPproto.HealthResponse{Recieved: true}, nil)

	go testExecutorClient.StartPing()
	time.Sleep(1 * time.Second)
	mockGrpcClient.AssertExpectations(t)
}

func TestStartKubernetesServiceWithNoJob(t *testing.T) {
	mockGrpcClient := new(client.MockGrpcClient)
	testClient := NewExecutorClient(mockGrpcClient)
	testExecutorClient := testClient.(*executorClient)
	testExecutorClient.id = "test id"
	testExecutorClient.state = "idle"

	mockGrpcClient.On("FetchJob", &executorCPproto.ExecutorID{Id: "test id"}).Return(&executorCPproto.Job{HasJob: "no"}, nil)

	go testExecutorClient.StartKubernetesService()
	time.Sleep(1 * time.Second)
	assert.Equal(t, "idle", testExecutorClient.state)
	mockGrpcClient.AssertExpectations(t)
}

func TestStartKubernetesServiceHasJob(t *testing.T) {
	mockGrpcClient := new(client.MockGrpcClient)
	mockKubeClient := new(kubernetes.MockKubernetesClient)
	testClient := NewExecutorClient(mockGrpcClient)
	testExecutorClient := testClient.(*executorClient)
	testExecutorClient.id = "test id"
	testExecutorClient.state = "idle"
	testExecutorClient.kubernetesClient = mockKubeClient
	testExecutorClient.kubeLogWaitTime = time.Second
	testArgs := map[string]string{"data": "test data"}

	testJob := &executorCPproto.Job{
		HasJob:    "yes",
		JobID:     "123",
		ImageName: "test image",
		JobData:   testArgs,
	}

	mockGrpcClient.On("FetchJob", &executorCPproto.ExecutorID{Id: "test id"}).Return(testJob, nil)
	mockKubeClient.On("ExecuteJob", "123", "test image", testArgs).Return("test execution", nil)
	mockKubeClient.On("WaitForReadyJob", "test execution", time.Second).Return(nil)

	go testExecutorClient.StartKubernetesService()
	time.Sleep(time.Second)

	assert.Equal(t, RunningState, testExecutorClient.state)
	mockGrpcClient.AssertExpectations(t)
}
