package daemon

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
	log.Init("info", "", false,1)
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
	mockGrpcClient.On("Ping", &executorCPproto.Ping{ID: "test id"}).Return(&executorCPproto.HealthResponse{Recieved: true}, nil)

	testExecutorClient.StartPing()
	time.Sleep(6 * time.Second)
	mockGrpcClient.AssertExpectations(t)
}

func TestStartKubernetesServiceWithNoJob(t *testing.T) {
	mockGrpcClient := new(client.MockGrpcClient)
	testClient := NewExecutorClient(mockGrpcClient)
	testExecutorClient := testClient.(*executorClient)
	testExecutorClient.id = "test id"
	testExecutorClient.state = "idle"

	mockGrpcClient.On("FetchJob", &executorCPproto.ExecutorID{Id: "test id"}).Return(&executorCPproto.Job{HasJob: false}, nil)

	go testExecutorClient.StartKubernetesService()
	time.Sleep(1 * time.Second)
	assert.Equal(t, "idle", testExecutorClient.state)
	mockGrpcClient.AssertExpectations(t)
}

func TestStartKubernetesServiceCreationFailed(t *testing.T) {
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
		HasJob:    true,
		JobID:     "123",
		ImageName: "test image",
		JobData:   testArgs,
	}

	testExecutionContext := &executorCPproto.ExecutionContext{
		JobK8SName: "",
		JobID:      "123",
		ImageName:  "test image",
		ExecutorID: "test id",
		Status:     "CREATION_FAILED",
		EnvArgs:    testArgs,
	}

	mockGrpcClient.On("FetchJob", &executorCPproto.ExecutorID{Id: "test id"}).Return(testJob, nil)
	mockKubeClient.On("ExecuteJob", "123", "test image", testArgs).Return("", errors.New("test error"))
	mockGrpcClient.On("SendExecutionContext", testExecutionContext).Return(&executorCPproto.Acknowledgement{}, nil)
	mockGrpcClient.On("Ping", &executorCPproto.Ping{ID: "test id", State: "running"}).Return(&executorCPproto.HealthResponse{Recieved: true}, nil)

	go testExecutorClient.StartKubernetesService()
	time.Sleep(1 * time.Second)

	assert.Equal(t, constant.RunningState, testExecutorClient.state)
	mockGrpcClient.AssertExpectations(t)
	mockKubeClient.AssertExpectations(t)
}

func TestStartKubernetesService(t *testing.T) {
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
		HasJob:    true,
		JobID:     "123",
		ImageName: "test image",
		JobData:   testArgs,
	}

	testExecutionContext := &executorCPproto.ExecutionContext{
		JobK8SName: "",
		JobID:      "123",
		ImageName:  "test image",
		ExecutorID: "test id",
		Status:     "FINISHED",
		EnvArgs:    testArgs,
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

	mockGrpcClient.On("FetchJob", &executorCPproto.ExecutorID{Id: "test id"}).Return(testJob, nil)
	mockKubeClient.On("ExecuteJob", "123", "test image", testArgs).Return("", nil)
	mockKubeClient.On("WaitForReadyJob", "", time.Second).Return(nil)
	mockKubeClient.On("WaitForReadyPod", "", time.Second).Return(pod, nil)
	mockKubeClient.On("GetPodLogs", pod).Return(stringReadCloser, nil)
	mockGrpcClient.On("SendExecutionContext", testExecutionContext).Return(&executorCPproto.Acknowledgement{}, nil)
	mockGrpcClient.On("Ping", &executorCPproto.Ping{ID: "test id", State: "idle"}).Return(&executorCPproto.HealthResponse{Recieved: true}, nil)
	mockGrpcClient.On("Ping", &executorCPproto.Ping{ID: "test id", State: "running"}).Return(&executorCPproto.HealthResponse{Recieved: true}, nil)

	go testExecutorClient.StartKubernetesService()
	time.Sleep(1 * time.Second)

	mockGrpcClient.AssertExpectations(t)
	mockKubeClient.AssertExpectations(t)
}

func TestStartWatch(t *testing.T) {
	mockGrpcClient := new(client.MockGrpcClient)
	mockKubeClient := new(kubernetes.MockKubernetesClient)
	testClient := NewExecutorClient(mockGrpcClient)
	testExecutorClient := testClient.(*executorClient)
	testExecutorClient.id = "test id"
	testExecutorClient.state = "idle"
	testExecutorClient.kubernetesClient = mockKubeClient
	testExecutorClient.kubeLogWaitTime = time.Second
	testArgs := map[string]string{"data": "test data"}

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

	testExecutionContext := &executorCPproto.ExecutionContext{
		JobK8SName: "test execution",
		JobID:      "123",
		ImageName:  "test image",
		ExecutorID: "test id",
		Status:     "CREATED",
		EnvArgs:    testArgs,
	}

	finalExecutionContext := &executorCPproto.ExecutionContext{
		JobK8SName: "test execution",
		JobID:      "123",
		ImageName:  "test image",
		ExecutorID: "test id",
		Status:     "FINISHED",
		EnvArgs:    testArgs,
		Output:     "test logs\n",
	}

	stringReadCloser := ioutil.NopCloser(strings.NewReader("test logs"))

	mockKubeClient.On("WaitForReadyJob", "test execution", time.Second).Return(nil)
	mockKubeClient.On("WaitForReadyPod", "test execution", time.Second).Return(pod, nil)
	mockKubeClient.On("GetPodLogs", pod).Return(stringReadCloser, nil)
	mockGrpcClient.On("SendExecutionContext", finalExecutionContext).Return(&executorCPproto.Acknowledgement{}, nil)

	testClient.startWatch(testExecutionContext)

	mockKubeClient.AssertExpectations(t)
	mockGrpcClient.AssertExpectations(t)
}
