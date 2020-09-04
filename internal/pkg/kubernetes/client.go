package kubernetes

import (
	"context"
	"errors"
	"fmt"
	"io"
	"octavius/internal/executor/config"
	"octavius/internal/pkg/constant"
	"os"
	"path/filepath"
	"time"

	batch "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	typeMeta     meta.TypeMeta
	ctx          context.Context
	cancel       context.CancelFunc
	timeoutError = errors.New("timeout when waiting job to be available")
)

func init() {
	typeMeta = meta.TypeMeta{
		Kind:       "Job",
		APIVersion: "batch/v1",
	}
}

type KubeClient interface {
	ExecuteJob(jobId string, imageName string, envMap map[string]string) (string, error)
	ExecuteJobWithCommand(jobId string, imageName string, args map[string]string, commands []string) (string, error)
	JobExecutionStatus(executionName string) (string, error)
	GetPodLogs(pod *v1.Pod) (io.ReadCloser, error)
	WaitForReadyJob(executionName string, waitTime time.Duration) error
	WaitForReadyPod(executionName string, waitTime time.Duration) (*v1.Pod, error)
}

type kubeClient struct {
	clientSet                    kubernetes.Interface
	namespace                    string
	kubeServiceAccountName       string
	jobPodAnnotations            map[string]string
	kubeJobActiveDeadlineSeconds int
	kubeJobRetries               int
	kubeWaitForResourcePollCount int
}

func NewClientSet(kubernetesConfig config.OctaviusExecutorConfig) (*kubernetes.Clientset, error) {
	var kubeConfig *rest.Config
	if kubernetesConfig.KubeConfig == "out-of-cluster" {

		home := os.Getenv("HOME")
		kubeConfigPath := filepath.Join(home, ".kube", "kubeconfig")

		configOverrides := &clientcmd.ConfigOverrides{}
		if kubernetesConfig.KubeContext != "default" {
			configOverrides.CurrentContext = kubernetesConfig.KubeContext
		}

		var err error
		kubeConfig, err = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath},
			configOverrides).ClientConfig()
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		kubeConfig, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	}

	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, err
	}

	return clientSet, nil
}

func NewKubernetesClient(kubernetesConfig config.OctaviusExecutorConfig) (KubeClient, error) {
	newClient := &kubeClient{
		namespace:                    kubernetesConfig.DefaultNamespace,
		kubeServiceAccountName:       kubernetesConfig.KubeServiceAccountName,
		jobPodAnnotations:            kubernetesConfig.JobPodAnnotations,
		kubeJobActiveDeadlineSeconds: kubernetesConfig.KubeJobActiveDeadlineSeconds,
		kubeJobRetries:               kubernetesConfig.KubeJobRetries,
		kubeWaitForResourcePollCount: kubernetesConfig.KubeWaitForResourcePollCount,
	}

	var err error
	newClient.clientSet, err = NewClientSet(kubernetesConfig)
	if err != nil {
		return nil, err
	}

	return newClient, nil
}

func jobLabel(executionName string) map[string]string {
	return map[string]string{
		"job": executionName,
	}
}

func getEnvVars(envMap map[string]string) []v1.EnvVar {
	var envVars []v1.EnvVar
	for k, v := range envMap {
		envVar := v1.EnvVar{
			Name:  k,
			Value: v,
		}
		envVars = append(envVars, envVar)
	}
	return envVars
}

func (client *kubeClient) ExecuteJob(jobId string, imageName string, envMap map[string]string) (string, error) {
	return client.ExecuteJobWithCommand(jobId, imageName, envMap, []string{})
}

func (client *kubeClient) ExecuteJobWithCommand(jobId string, imageName string, envMap map[string]string, command []string) (string, error) {
	executionName := "octavius-" + jobId

	label := jobLabel(executionName)

	batchV1 := client.clientSet.BatchV1()
	kubernetesJobs := batchV1.Jobs(client.namespace)

	container := v1.Container{
		Name:  executionName,
		Image: imageName,
		Env:   getEnvVars(envMap),
	}

	if len(command) != 0 {
		container.Command = command
	}

	podSpec := v1.PodSpec{
		Containers:         []v1.Container{container},
		RestartPolicy:      v1.RestartPolicyNever,
		ServiceAccountName: client.kubeServiceAccountName,
	}

	objectMeta := meta.ObjectMeta{
		Name:        executionName,
		Labels:      label,
		Annotations: client.jobPodAnnotations,
	}

	template := v1.PodTemplateSpec{
		ObjectMeta: objectMeta,
		Spec:       podSpec,
	}
	jobDeadline := int64(client.kubeJobActiveDeadlineSeconds)
	jobRetries := int32(client.kubeJobRetries)

	jobSpec := batch.JobSpec{
		Template:              template,
		ActiveDeadlineSeconds: &jobDeadline,
		BackoffLimit:          &jobRetries,
	}

	jobToRun := batch.Job{
		TypeMeta:   typeMeta,
		ObjectMeta: objectMeta,
		Spec:       jobSpec,
	}
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	_, err := kubernetesJobs.Create(ctx, &jobToRun, meta.CreateOptions{})
	if err != nil {
		return "", err
	}

	return executionName, nil
}

func jobLabelSelector(executionName string) string {
	return fmt.Sprintf("job=%s", executionName)
}

func (client *kubeClient) JobExecutionStatus(executionName string) (string, error) {
	batchV1 := client.clientSet.BatchV1()
	kubernetesJobs := batchV1.Jobs(client.namespace)
	listOptions := meta.ListOptions{
		TypeMeta:      typeMeta,
		LabelSelector: jobLabelSelector(executionName),
	}

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	watchJob, err := kubernetesJobs.Watch(ctx, listOptions)
	if err != nil {
		return constant.JobFailed, err
	}

	resultChan := watchJob.ResultChan()
	defer watchJob.Stop()
	var event watch.Event
	var jobEvent *batch.Job

	for event = range resultChan {
		if event.Type == watch.Error {
			return constant.JobExecutionStatusFetchError, nil
		}

		jobEvent = event.Object.(*batch.Job)
		if jobEvent.Status.Succeeded >= int32(1) {
			return constant.JobSucceeded, nil
		} else if jobEvent.Status.Failed >= int32(1) {
			return constant.JobFailed, nil
		}
	}

	return constant.NoDefinitiveJobExecutionStatusFound, nil
}

func (client *kubeClient) GetPodLogs(pod *v1.Pod) (io.ReadCloser, error) {
	podLogOpts := v1.PodLogOptions{
		Follow: true,
	}

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	request := client.clientSet.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &podLogOpts)
	response, err := request.Stream(ctx)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client *kubeClient) WaitForReadyJob(executionName string, waitTime time.Duration) error {
	jobs := client.clientSet.BatchV1().Jobs(client.namespace)
	listOptions := meta.ListOptions{
		TypeMeta:      typeMeta,
		LabelSelector: jobLabelSelector(executionName),
	}

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	var err error
	for i := 0; i < client.kubeWaitForResourcePollCount; i += 1 {
		watchJob, watchErr := jobs.Watch(ctx, listOptions)
		if watchErr != nil {
			err = watchErr
			continue
		}

		timeoutChan := time.After(waitTime)
		resultChan := watchJob.ResultChan()

		var job *batch.Job
		for {
			select {
			case event := <-resultChan:
				if event.Type == watch.Error {
					err = watcherError("jobs", listOptions)
					break
				}

				// Ignore empty events
				if event.Object == nil {
					continue
				}

				job = event.Object.(*batch.Job)
				if job.Status.Active >= 1 || job.Status.Succeeded >= 1 || job.Status.Failed >= 1 {
					watchJob.Stop()
					return nil
				}
			case <-timeoutChan:
				err = timeoutError
				break
			}
			if err != nil {
				watchJob.Stop()
				break
			}
		}
	}

	return err
}

func watcherError(resource string, listOptions meta.ListOptions) error {
	return fmt.Errorf("watch error when waiting for %s with list option %v", resource, listOptions)
}

func (client *kubeClient) WaitForReadyPod(executionName string, waitTime time.Duration) (*v1.Pod, error) {
	coreV1 := client.clientSet.CoreV1()
	kubernetesPods := coreV1.Pods(client.namespace)
	listOptions := meta.ListOptions{
		LabelSelector: jobLabelSelector(executionName),
	}

	var err error
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < client.kubeWaitForResourcePollCount; i += 1 {
		watchJob, watchErr := kubernetesPods.Watch(ctx, listOptions)
		if watchErr != nil {
			err = watchErr
			continue
		}

		timeoutChan := time.After(waitTime)
		resultChan := watchJob.ResultChan()

		var pod *v1.Pod
		for {
			select {
			case event := <-resultChan:
				if event.Type == watch.Error {
					err = watcherError("pods", listOptions)
					watchJob.Stop()
					break
				}

				// Ignore empty events
				if event.Object == nil {
					continue
				}

				pod = event.Object.(*v1.Pod)
				if pod.Status.Phase == v1.PodRunning || pod.Status.Phase == v1.PodSucceeded || pod.Status.Phase == v1.PodFailed {
					watchJob.Stop()
					return pod, nil
				}
			case <-timeoutChan:
				err = timeoutError
				watchJob.Stop()
				break
			}
			if err != nil {
				watchJob.Stop()
				break
			}
		}
	}

	return nil, err
}
