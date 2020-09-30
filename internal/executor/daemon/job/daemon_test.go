package job

import (
	"errors"
	"io/ioutil"
	"octavius/internal/executor/client/job"
	client "octavius/internal/executor/client/job"
	"octavius/internal/executor/config"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/kubernetes"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/protofiles"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	v1 "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func TestStartKubernetesServiceWithNoJob(t *testing.T) {
	mockGrpcClient := new(job.MockGrpcClient)
	testClient := NewJobServicesClient(mockGrpcClient)
	testJobServicesClient := testClient.(*jobServicesClient)
	testJobServicesClient.id = "test id"
	testJobServicesClient.state = "idle"

	mockGrpcClient.On("FetchJob", &protofiles.ExecutorID{ID: "test id"}).Return(&protofiles.Job{HasJob: false}, nil)
	mockGrpcClient.On("ConnectClient", "test host", time.Second).Return(nil)
	mockGrpcClient.On("PostExecutorStatus", &protofiles.Status{ID: "test id", Status: "idle"}).Return(&protofiles.Acknowledgement{Received: true}, nil)

	go testJobServicesClient.StartKubernetesService(testConfig)
	time.Sleep(2 * time.Second)
	mockGrpcClient.AssertExpectations(t)
}

func TestStartKubernetesServiceCreationFailed(t *testing.T) {
	mockGrpcClient := new(client.MockGrpcClient)
	mockKubeClient := new(kubernetes.MockKubernetesClient)
	testClient := NewJobServicesClient(mockGrpcClient)
	testJobServicesClient := testClient.(*jobServicesClient)
	testJobServicesClient.id = "test id"
	testJobServicesClient.state = "idle"
	testJobServicesClient.kubernetesClient = mockKubeClient
	testJobServicesClient.kubeLogWaitTime = time.Second
	testArgs := map[string]string{"data": "test data"}

	testJob := &protofiles.Job{
		HasJob:    true,
		JobID:     "123",
		ImageName: "test image",
		JobData:   testArgs,
	}

	testExecutionContext := &protofiles.ExecutionContext{
		JobK8SName: "",
		JobID:      "123",
		ImageName:  "test image",
		ExecutorID: "test id",
		Status:     "CREATION_FAILED",
		EnvArgs:    testArgs,
	}

	mockGrpcClient.On("FetchJob", &protofiles.ExecutorID{ID: "test id"}).Return(testJob, nil)
	mockKubeClient.On("ExecuteJob", "123", "test image", testArgs).Return("", errors.New("test error"))
	mockGrpcClient.On("SendExecutionContext", testExecutionContext).Return(&protofiles.Acknowledgement{}, nil)
	mockGrpcClient.On("ConnectClient", "test host", time.Second).Return(nil)
	mockGrpcClient.On("PostExecutorStatus", mock.Anything).Return(&protofiles.Acknowledgement{Received: true}, nil)

	go testJobServicesClient.StartKubernetesService(testConfig)
	time.Sleep(2 * time.Second)

	testJobServicesClient.statusLock.RLock()
	assert.Equal(t, constant.IdleState, testJobServicesClient.state)
	testJobServicesClient.statusLock.RUnlock()
	mockGrpcClient.AssertExpectations(t)
	mockKubeClient.AssertExpectations(t)
}

func TestStartKubernetesService(t *testing.T) {
	mockGrpcClient := new(client.MockGrpcClient)
	mockKubeClient := new(kubernetes.MockKubernetesClient)
	testClient := NewJobServicesClient(mockGrpcClient)
	testJobServicesClient := testClient.(*jobServicesClient)
	testJobServicesClient.id = "test id"
	testJobServicesClient.state = "idle"
	testJobServicesClient.kubernetesClient = mockKubeClient
	testJobServicesClient.kubeLogWaitTime = time.Second
	testArgs := map[string]string{"data": "test data"}

	testJob := &protofiles.Job{
		HasJob:    true,
		JobID:     "123",
		ImageName: "test image",
		JobData:   testArgs,
	}

	testExecutionContext := &protofiles.ExecutionContext{
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

	mockGrpcClient.On("ConnectClient", "test host", time.Second).Return(nil)
	mockGrpcClient.On("PostExecutorStatus", mock.Anything).Return(&protofiles.Acknowledgement{Received: true}, nil)
	mockGrpcClient.On("FetchJob", &protofiles.ExecutorID{ID: "test id"}).Return(testJob, nil)
	mockKubeClient.On("ExecuteJob", "123", "test image", testArgs).Return("", nil)
	mockKubeClient.On("WaitForReadyJob", "", time.Second).Return(nil)
	mockKubeClient.On("WaitForReadyPod", "", time.Second).Return(pod, nil)
	mockKubeClient.On("GetPodLogs", pod).Return(stringReadCloser, nil)
	mockGrpcClient.On("SendExecutionContext", testExecutionContext).Return(&protofiles.Acknowledgement{}, nil)

	go testJobServicesClient.StartKubernetesService(testConfig)
	time.Sleep(1 * time.Second)

	mockGrpcClient.AssertExpectations(t)
	mockKubeClient.AssertExpectations(t)
}

func TestStartWatch(t *testing.T) {
	mockGrpcClient := new(client.MockGrpcClient)
	mockKubeClient := new(kubernetes.MockKubernetesClient)
	testClient := NewJobServicesClient(mockGrpcClient)
	testJobServicesClient := testClient.(*jobServicesClient)
	testJobServicesClient.id = "test id"
	testJobServicesClient.state = "idle"
	testJobServicesClient.kubernetesClient = mockKubeClient
	testJobServicesClient.kubeLogWaitTime = time.Second
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

	testExecutionContext := &protofiles.ExecutionContext{
		JobK8SName: "test execution",
		JobID:      "123",
		ImageName:  "test image",
		ExecutorID: "test id",
		Status:     "CREATED",
		EnvArgs:    testArgs,
	}

	intermediateExecutionContext := &protofiles.ExecutionContext{
		JobK8SName: "test execution",
		JobID:      "123",
		ImageName:  "test image",
		ExecutorID: "test id",
		Status:     "running",
		EnvArgs:    testArgs,
		Output:     "test logs\n",
	}

	finalExecutionContext := &protofiles.ExecutionContext{
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
	mockGrpcClient.On("SendExecutionContext", intermediateExecutionContext).Return(&protofiles.Acknowledgement{}, nil)
	mockGrpcClient.On("SendExecutionContext", finalExecutionContext).Return(&protofiles.Acknowledgement{}, nil)

	testClient.startWatch(testExecutionContext)

	mockKubeClient.AssertExpectations(t)
	//mockGrpcClient.AssertExpectations(t)
}
