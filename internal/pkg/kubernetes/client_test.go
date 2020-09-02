package kubernetes

import (
	"context"
	"github.com/stretchr/testify/assert"
	batchV1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	fakeClientSet "k8s.io/client-go/kubernetes/fake"
	batch "k8s.io/client-go/kubernetes/typed/batch/v1"
	testingKubernetes "k8s.io/client-go/testing"
	"net/http"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/kubernetes/config"
	"os"
	"testing"
	"time"
)

var testKubeClient TestKubeClient

type TestKubeClient struct {
	testClient				KubeClient
	testKubernetesJobs  	batch.JobInterface
	fakeClientSet       	*fakeClientSet.Clientset
	jobName             	string
	podName             	string
	fakeClientSetStreaming  *fakeClientSet.Clientset
	fakeHTTPClient			*http.Client
	testClientStreaming		KubeClient
}

func setUpTest() {
	testKubeClient.fakeClientSet = fakeClientSet.NewSimpleClientset()
	testKubeClient.testClient = &kubeClient{
		clientSet: testKubeClient.fakeClientSet,
	}
	testKubeClient.jobName = "job1"
	testKubeClient.podName = "pod1"
	namespace = "default"
	testKubeClient.fakeClientSetStreaming = fakeClientSet.NewSimpleClientset(&v1.Pod{
		TypeMeta: meta.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: meta.ObjectMeta{
			Name:      testKubeClient.podName,
			Namespace: namespace,
			Labels: map[string]string{
				"tag": "",
				"job": testKubeClient.jobName,
			},
		},
		Status: v1.PodStatus{
			Phase: v1.PodSucceeded,
		},
	})

	testKubeClient.fakeHTTPClient =&http.Client{}
	testKubeClient.testClientStreaming = &kubeClient{
		clientSet: testKubeClient.fakeClientSetStreaming,
	}
}

func TestJobExecution(t *testing.T) {
	setUpTest()
	_ = os.Setenv("job_pod_annotations", "{\"key.one\":\"true\"}")
	_ = os.Setenv("service_account_name", "default")
	config.Reset()
	envVarsForContainer := map[string]string{"SAMPLE_ARG": "sample-value"}
	sampleImageName := "img1"

	executedJobName, err := testKubeClient.testClient.ExecuteJob(sampleImageName, envVarsForContainer)
	assert.NoError(t, err)

	typeMeta := meta.TypeMeta{
		Kind:       "Job",
		APIVersion: "batch/v1",
	}

	listOptions := meta.ListOptions{
		TypeMeta:      typeMeta,
		LabelSelector: jobLabelSelector(executedJobName),
	}
	namespace = "default"
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	listOfJobs, err := testKubeClient.fakeClientSet.BatchV1().Jobs(namespace).List(ctx, listOptions)
	assert.NoError(t, err)
	executedJob := listOfJobs.Items[0]

	assert.Equal(t, typeMeta, executedJob.TypeMeta)

	assert.Equal(t, executedJobName, executedJob.ObjectMeta.Name)
	assert.Equal(t, executedJobName, executedJob.Spec.Template.ObjectMeta.Name)

	expectedLabel := jobLabel(executedJobName)
	assert.Equal(t, expectedLabel, executedJob.ObjectMeta.Labels)
	assert.Equal(t, expectedLabel, executedJob.Spec.Template.ObjectMeta.Labels)
	assert.Equal(t, map[string]string{"key.one": "true"}, executedJob.Spec.Template.Annotations)
	assert.Equal(t, "default", executedJob.Spec.Template.Spec.ServiceAccountName)

	assert.Equal(t, 60, executedJob.Spec.ActiveDeadlineSeconds)
	assert.Equal(t, 0, executedJob.Spec.BackoffLimit)

	assert.Equal(t, v1.RestartPolicyNever, executedJob.Spec.Template.Spec.RestartPolicy)

	container := executedJob.Spec.Template.Spec.Containers[0]
	assert.Equal(t, executedJobName, container.Name)

	assert.Equal(t, sampleImageName, container.Image)

	expectedEnvVars := getEnvVars(envVarsForContainer)
	assert.Equal(t, expectedEnvVars, container.Env)

}

func TestWaitForReadyJob(t *testing.T) {
	setUpTest()

	var testJob batchV1.Job
	uniqueJobName := "octavius-job-1"
	label := jobLabel(uniqueJobName)
	objectMeta := meta.ObjectMeta{
		Name:   uniqueJobName,
		Labels: label,
	}
	testJob.ObjectMeta = objectMeta
	waitTime := 60 * time.Second

	watcher := watch.NewFake()
	testKubeClient.fakeClientSet.PrependWatchReactor("jobs", testingKubernetes.DefaultWatchReactor(watcher, nil))

	go func() {
		testJob.Status.Succeeded = 1
		watcher.Modify(&testJob)

		watcher.Stop()
	}()

	err := testKubeClient.testClient.WaitForReadyJob(uniqueJobName, waitTime)
	assert.NoError(t, err)
}

func TestShouldReturnSuccessJobExecutionStatus(t *testing.T) {
	setUpTest()

	watcher := watch.NewFake()
	testKubeClient.fakeClientSet.PrependWatchReactor("jobs", testingKubernetes.DefaultWatchReactor(watcher, nil))

	var activeJob batchV1.Job
	var succeededJob batchV1.Job
	uniqueJobName := "octavius-job-2"
	label := jobLabel(uniqueJobName)
	objectMeta := meta.ObjectMeta{
		Name:   uniqueJobName,
		Labels: label,
	}
	activeJob.ObjectMeta = objectMeta
	succeededJob.ObjectMeta = objectMeta

	go func() {
		activeJob.Status.Active = 1
		watcher.Modify(&activeJob)

		succeededJob.Status.Active = 0
		succeededJob.Status.Succeeded = 1
		watcher.Modify(&succeededJob)

		time.Sleep(time.Second * 1)
		watcher.Stop()
	}()

	jobExecutionStatus, err := testKubeClient.testClient.JobExecutionStatus(uniqueJobName)
	assert.NoError(t, err)

	assert.Equal(t, constant.JobSucceeded, jobExecutionStatus, "Should return succeeded")
}

func TestShouldReturnFailedJobExecutionStatus(t *testing.T) {
	setUpTest()

	watcher := watch.NewFake()
	testKubeClient.fakeClientSet.PrependWatchReactor("jobs", testingKubernetes.DefaultWatchReactor(watcher, nil))

	var activeJob batchV1.Job
	var failedJob batchV1.Job
	uniqueJobName := "octavius-job-3"
	label := jobLabel(uniqueJobName)
	objectMeta := meta.ObjectMeta{
		Name:   uniqueJobName,
		Labels: label,
	}
	activeJob.ObjectMeta = objectMeta
	failedJob.ObjectMeta = objectMeta

	go func() {
		activeJob.Status.Active = 1
		watcher.Modify(&activeJob)
		failedJob.Status.Active = 0
		failedJob.Status.Failed = 1
		watcher.Modify(&failedJob)

		time.Sleep(time.Second * 1)
		watcher.Stop()
	}()

	jobExecutionStatus, err := testKubeClient.testClient.JobExecutionStatus(uniqueJobName)
	assert.NoError(t, err)

	assert.Equal(t, constant.JobFailed, jobExecutionStatus, "Should return FAILED")
}

func TestJobExecutionStatusForNonDefinitiveStatus(t *testing.T) {
	setUpTest()

	watcher := watch.NewFake()
	testKubeClient.fakeClientSet.PrependWatchReactor("jobs", testingKubernetes.DefaultWatchReactor(watcher, nil))

	var testJob batchV1.Job
	uniqueJobName := "proctor-job-4"
	label := jobLabel(uniqueJobName)
	objectMeta := meta.ObjectMeta{
		Name:   uniqueJobName,
		Labels: label,
	}
	testJob.ObjectMeta = objectMeta

	go func() {
		testJob.Status.Active = 1
		watcher.Modify(&testJob)

		time.Sleep(time.Second * 1)
		watcher.Stop()
	}()

	jobExecutionStatus, err := testKubeClient.testClient.JobExecutionStatus(uniqueJobName)
	assert.NoError(t, err)

	assert.Equal(t, constant.NoDefinitiveJobExecutionStatusFound, jobExecutionStatus, "Should return NO_DEFINITIVE_JOB_EXECUTION_STATUS_FOUND")
}

func TestShouldReturnJobExecutionStatusFetchError(t *testing.T) {
	setUpTest()

	watcher := watch.NewFake()
	testKubeClient.fakeClientSet.PrependWatchReactor("jobs", testingKubernetes.DefaultWatchReactor(watcher, nil))

	var testJob batchV1.Job
	uniqueJobName := "proctor-job-5"
	label := jobLabel(uniqueJobName)
	objectMeta := meta.ObjectMeta{
		Name:   uniqueJobName,
		Labels: label,
	}
	testJob.ObjectMeta = objectMeta

	go func() {
		watcher.Error(&testJob)

		time.Sleep(time.Second * 1)
		watcher.Stop()
	}()

	jobExecutionStatus, err := testKubeClient.testClient.JobExecutionStatus(uniqueJobName)
	assert.NoError(t, err)

	assert.Equal(t, constant.JobExecutionStatusFetchError, jobExecutionStatus, "Should return JOB_EXECUTION_STATUS_FETCH_ERROR")
}