package job

import (
	"errors"
	"io/ioutil"
	"octavius/internal/executor/client"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/kubernetes"
	"octavius/internal/pkg/log"
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

func TestStartKubernetesServiceWithNoJob(t *testing.T) {
	mockGrpcClient := new(client.MockGrpcClient)
	testClient := NewJobServicesClient(mockGrpcClient)
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

	go testExecutorClient.StartKubernetesService()
	time.Sleep(1 * time.Second)

	testExecutorClient.statusLock.RLock()
	assert.Equal(t, constant.RunningState, testExecutorClient.state)
	testExecutorClient.statusLock.RUnlock()
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
