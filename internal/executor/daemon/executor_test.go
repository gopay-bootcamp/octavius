package daemon

import (
	"io/ioutil"
	"octavius/internal/executor/client"
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
	log.Init("info", "", false)
}

func TestStartClient(t *testing.T) {

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

	pod := &v1.Pod{
		TypeMeta: meta.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: meta.ObjectMeta{
			Name:      "test pod",
			Namespace: "default",
			Labels: map[string]string{
				"tag": "",
				"job": "test job",
			},
		},
		Status: v1.PodStatus{
			Phase: v1.PodSucceeded,
		},
	}

	stringReadCloser := ioutil.NopCloser(strings.NewReader(""))

	testExecutionContext := &executorCPproto.ExecutionContext{
		Name:       "test execution",
		JobID:      "123",
		JobName:    "test image",
		ExecutorID: "test id",
		Status:     "FINISHED",
		EnvArgs:    testArgs,
		Output:     "",
	}

	mockGrpcClient.On("FetchJob", &executorCPproto.ExecutorID{Id: "test id"}).Return(testJob, nil)
	mockKubeClient.On("ExecuteJob", "123", "test image", testArgs).Return("test execution", nil)
	mockKubeClient.On("WaitForReadyJob", "test execution", time.Second).Return(nil)
	mockKubeClient.On("WaitForReadyPod", "test execution", time.Second).Return(pod, nil)
	mockKubeClient.On("GetPodLogs", pod).Return(stringReadCloser, nil)
	mockGrpcClient.On("SendExecutionContext", testExecutionContext).Return(&executorCPproto.Acknowledgement{}, nil)

	go testExecutorClient.StartKubernetesService()
	time.Sleep(1 * time.Second)

	assert.Equal(t, RunningState, testExecutorClient.state)
	mockGrpcClient.AssertExpectations(t)
	mockKubeClient.AssertExpectations(t)
}
