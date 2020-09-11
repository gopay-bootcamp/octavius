package kubernetes

import (
	"context"
	"io"
	"time"

	v1 "k8s.io/api/core/v1"

	"github.com/stretchr/testify/mock"
)

type MockKubernetesClient struct {
	mock.Mock
}

func (m *MockKubernetesClient) ExecuteJob(ctx context.Context, jobID string, imageName string, envMap map[string]string) (string, error) {
	args := m.Called(jobID, imageName, envMap)
	return args.String(0), args.Error(1)
}

func (m *MockKubernetesClient) ExecuteJobWithCommands(ctx context.Context, jobID string, imageName string, jobArgs map[string]string, commands []string) (string, error) {
	args := m.Called(jobID, imageName, jobArgs, commands)
	return args.String(0), args.Error(1)
}

func (m *MockKubernetesClient) JobExecutionStatus(ctx context.Context, executionName string) (string, error) {
	args := m.Called(executionName)
	return args.String(0), args.Error(1)
}

func (m *MockKubernetesClient) WaitForReadyJob(ctx context.Context, executionName string, waitTime time.Duration) error {
	args := m.Called(executionName, waitTime)
	return args.Error(0)
}

func (m *MockKubernetesClient) WaitForReadyPod(ctx context.Context, executionName string, waitTime time.Duration) (*v1.Pod, error) {
	args := m.Called(executionName, waitTime)
	return args.Get(0).(*v1.Pod), args.Error(1)
}

func (m *MockKubernetesClient) GetPodLogs(ctx context.Context, pod *v1.Pod) (io.ReadCloser, error) {
	args := m.Called(pod)
	return args.Get(0).(io.ReadCloser), args.Error(1)
}
