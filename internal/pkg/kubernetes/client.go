package kubernetes

import (
	"context"
	"errors"
	"fmt"
	"io"
	"octavius/internal/executor/config"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/log"
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
	typeMeta meta.TypeMeta
)

func init() {
	typeMeta = meta.TypeMeta{
		Kind:       "Job",
		APIVersion: "batch/v1",
	}
}

// KubeClient implements methods to interact with kubernetes cluster
type KubeClient interface {
	ExecuteJob(ctx context.Context, jobID string, imageName string, envMap map[string]string) (string, error)
	ExecuteJobWithCommands(ctx context.Context, jobID string, imageName string, args map[string]string, commands []string) (string, error)
	JobExecutionStatus(ctx context.Context, executionName string) (string, error)
	GetPodLogs(ctx context.Context, pod *v1.Pod) (io.ReadCloser, error)
	WaitForReadyJob(ctx context.Context, executionName string, waitTime time.Duration) error
	WaitForReadyPod(ctx context.Context, executionName string, waitTime time.Duration) (*v1.Pod, error)
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

func newClientSet(kubernetesConfig config.OctaviusExecutorConfig) (*kubernetes.Clientset, error) {
	var kubeConfig *rest.Config
	if kubernetesConfig.KubeConfig == constant.OutOfClustor {

		home := os.Getenv("HOME")
		kubeConfigPath := filepath.Join(home, ".kube", "config")

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

//NewKubernetesClient creates new kubernetes client with configurations
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
	newClient.clientSet, err = newClientSet(kubernetesConfig)
	if err != nil {
		return nil, err
	}

	return newClient, nil
}

func jobLabel(executionName string) map[string]string {
	return map[string]string{
		"app.kubernetes.io/job": executionName,
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

func (client *kubeClient) ExecuteJob(ctx context.Context, jobID string, imageName string, envMap map[string]string) (string, error) {
	return client.ExecuteJobWithCommands(ctx, jobID, imageName, envMap, []string{})
}

func (client *kubeClient) ExecuteJobWithCommands(ctx context.Context, jobID string, imageName string, envMap map[string]string, commands []string) (string, error) {
	executionName := fmt.Sprintf("%s-%s", constant.ExecutionKey, jobID)

	label := jobLabel(executionName)

	batchV1 := client.clientSet.BatchV1()
	kubernetesJobs := batchV1.Jobs(client.namespace)

	container := v1.Container{
		Name:  executionName,
		Image: imageName,
		Env:   getEnvVars(envMap),
	}

	if len(commands) != 0 {
		container.Command = commands
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

	_, err := kubernetesJobs.Create(ctx, &jobToRun, meta.CreateOptions{})
	if err != nil {
		return "", err
	}

	return executionName, nil
}

func jobLabelSelector(executionName string) string {
	return fmt.Sprintf("app.kubernetes.io/job=%s", executionName)
}

func (client *kubeClient) JobExecutionStatus(ctx context.Context, executionName string) (string, error) {
	batchV1 := client.clientSet.BatchV1()
	kubernetesJobs := batchV1.Jobs(client.namespace)
	listOptions := meta.ListOptions{
		TypeMeta:      typeMeta,
		LabelSelector: jobLabelSelector(executionName),
	}

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

func (client *kubeClient) GetPodLogs(ctx context.Context, pod *v1.Pod) (io.ReadCloser, error) {
	podLogOpts := v1.PodLogOptions{
		Follow: true,
	}

	request := client.clientSet.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &podLogOpts)
	response, err := request.Stream(ctx)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client *kubeClient) WaitForReadyJob(ctx context.Context, executionName string, waitTime time.Duration) error {
	jobs := client.clientSet.BatchV1().Jobs(client.namespace)
	listOptions := meta.ListOptions{
		TypeMeta:      typeMeta,
		LabelSelector: jobLabelSelector(executionName),
	}

	var (
		err      error
		watchJob watch.Interface
	)
	for i := 0; i < client.kubeWaitForResourcePollCount; i++ {
		watchJob, err = jobs.Watch(ctx, listOptions)
		if err != nil {
			log.Error(err, "error on watching job")
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
				err = errors.New(constant.TimeOutError)
				watchJob.Stop()
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

func (client *kubeClient) WaitForReadyPod(ctx context.Context, executionName string, waitTime time.Duration) (*v1.Pod, error) {
	coreV1 := client.clientSet.CoreV1()
	kubernetesPods := coreV1.Pods(client.namespace)
	listOptions := meta.ListOptions{
		LabelSelector: jobLabelSelector(executionName),
	}
	var (
		err      error
		watchJob watch.Interface
	)

	for i := 0; i < client.kubeWaitForResourcePollCount; i++ {
		watchJob, err = kubernetesPods.Watch(ctx, listOptions)
		if err != nil {
			log.Error(err, "error on watching pod")
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
				err = errors.New(constant.TimeOutError)
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
