package job

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	client "octavius/internal/executor/client/job"
	"octavius/internal/executor/config"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/kubernetes"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/protofiles"
	"sync"
	"time"

	v1 "k8s.io/api/core/v1"
)

// JobServicesClient interface
type JobServicesClient interface {
	connectClient(executorConfig config.OctaviusExecutorConfig) error
	FetchJob() (*protofiles.Job, error)
	StartKubernetesService(executorConfig config.OctaviusExecutorConfig)
	startWatch(executionContext *protofiles.ExecutionContext)
	ConfigureKubernetesClient(executorConfig config.OctaviusExecutorConfig) error
}

type jobServicesClient struct {
	id                    string
	grpcClient            client.Client
	cpHost                string
	accessToken           string
	connectionTimeoutSecs time.Duration
	pingInterval          time.Duration
	kubernetesClient      kubernetes.KubeClient
	kubeLogWaitTime       time.Duration
	state                 string
	statusLock            sync.RWMutex
}

//NewJobServicesClient returns new empty job services client
func NewJobServicesClient(grpcClient client.Client) JobServicesClient {
	return &jobServicesClient{
		grpcClient: grpcClient,
	}
}

// ConfigureKubernetesClient is used to configure the k8s client from values taken from Config file
func (e *jobServicesClient) ConfigureKubernetesClient(executorConfig config.OctaviusExecutorConfig) error {
	var kubeConfig = config.OctaviusExecutorConfig{
		KubeConfig:                   executorConfig.KubeConfig,
		KubeContext:                  executorConfig.KubeContext,
		DefaultNamespace:             executorConfig.DefaultNamespace,
		KubeServiceAccountName:       executorConfig.KubeServiceAccountName,
		JobPodAnnotations:            executorConfig.JobPodAnnotations,
		KubeJobActiveDeadlineSeconds: executorConfig.KubeJobActiveDeadlineSeconds,
		KubeJobRetries:               executorConfig.KubeJobRetries,
		KubeWaitForResourcePollCount: executorConfig.KubeWaitForResourcePollCount,
	}
	kubernetesClient, err := kubernetes.NewKubernetesClient(kubeConfig)
	if err != nil {
		return err
	}
	e.kubernetesClient = kubernetesClient
	return nil
}

func (e *jobServicesClient) connectClient(executorConfig config.OctaviusExecutorConfig) error {
	e.id = executorConfig.ID
	e.cpHost = executorConfig.CPHost
	e.accessToken = executorConfig.AccessToken
	e.connectionTimeoutSecs = executorConfig.ConnTimeOutSec
	e.pingInterval = executorConfig.PingInterval
	e.kubeLogWaitTime = time.Duration(executorConfig.KubeJobActiveDeadlineSeconds) * time.Second
	e.state = constant.IdleState
	err := e.grpcClient.ConnectClient(e.cpHost, e.connectionTimeoutSecs)

	return err
}

func (e *jobServicesClient) postExecutorStatus(stat string) (*protofiles.Acknowledgement, error) {
	return e.grpcClient.PostExecutorStatus(&protofiles.Status{ID: e.id, Status: stat})
}

// StartKubernetesService is used to start a blocking kuberentes service that fetches jobs from executor 
// at regular intervals
func (e *jobServicesClient) StartKubernetesService(executorConfig config.OctaviusExecutorConfig) {

	err := e.connectClient(executorConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = e.postExecutorStatus(constant.IdleState)
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		job, err := e.FetchJob()
		if err != nil {
			log.Fatal(fmt.Sprintf("error in getting job from server, error details: %s", err.Error()))
		}
		if !job.HasJob {
			time.Sleep(5 * time.Second)
			continue
		}
		e.statusLock.Lock()
		_, err = e.postExecutorStatus(constant.RunningState)
		if err != nil {
			log.Fatal(err.Error())
		}
		e.statusLock.Unlock()
		log.Info(fmt.Sprintf("received job from controller, job details: %+v", job))

		imageName := job.ImageName
		executionArgs := job.JobData
		jobID := job.JobID
		jobContext := protofiles.ExecutionContext{
			JobID:      jobID,
			ImageName:  imageName,
			EnvArgs:    executionArgs,
			Status:     constant.Created,
			ExecutorID: e.id,
		}
		ctx:= context.Background()

		log.Info("calling kubernets")
		executionName, err := e.kubernetesClient.ExecuteJob(ctx, jobID, imageName, executionArgs)
		log.Info(fmt.Sprintln("Executed Job on Kubernetes got ", executionName, " execution jobName and ", err, "errors"))
		if err != nil {
			jobContext.Status = constant.CreationFailed
			_, err := e.sendResponse(&jobContext)
			if err != nil {
				log.Fatal(err.Error())
			}
			log.Error(err, "error while executing job")
			time.Sleep(10 * time.Second)
			continue
		}

		jobContext.JobK8SName = executionName
		go e.startWatch(&jobContext)

		e.statusLock.Lock()
		_, err = e.postExecutorStatus(constant.IdleState)
		if err != nil {
			log.Fatal(err.Error())
		}
		e.statusLock.Unlock()
	}
}

func (e *jobServicesClient) sendResponse(jobContext *protofiles.ExecutionContext) (*protofiles.Acknowledgement, error) {
	return e.grpcClient.SendExecutionContext(jobContext)
}

func (e *jobServicesClient) startWatch(executionContext *protofiles.ExecutionContext) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Info(fmt.Sprintf("start watch process for job With k8s name: %s , job ID: %s, and status: %s", executionContext.JobK8SName, executionContext.JobID, executionContext.Status))

	err := e.kubernetesClient.WaitForReadyJob(ctx, executionContext.JobK8SName, e.kubeLogWaitTime)

	if err != nil {
		executionContext.Status = constant.JobCreationFailed
		return
	}

	executionContext.Status = constant.JobReady
	log.Info(fmt.Sprintf("Job Ready for %s", executionContext.JobK8SName))

	pod, err := e.kubernetesClient.WaitForReadyPod(ctx, executionContext.JobK8SName, e.kubeLogWaitTime)

	if err != nil {
		log.Error(err, fmt.Sprintf("wait for ready pod %s", pod.Name))
		executionContext.Status = constant.PodCreationFailed
		return
	}
	if pod.Status.Phase == v1.PodFailed {
		executionContext.Status = constant.PodFailed
		log.Info(fmt.Sprintf("Pod Failed for %s with reason: %s and message: %s", executionContext.JobK8SName, pod.Status.Reason, pod.Status.Message))
	} else {
		executionContext.Status = constant.PodReady
		log.Info(fmt.Sprintf("Pod Ready for %s", executionContext.JobK8SName))
	}

	podLog, err := e.kubernetesClient.GetPodLogs(ctx, pod)

	if err != nil {
		executionContext.Status = constant.FetchPodLogFailed
		return
	}
	defer podLog.Close()

	scanner := bufio.NewScanner(podLog)
	scanner.Split(bufio.ScanLines)

	var buffer bytes.Buffer
	for scanner.Scan() {
		buffer.WriteString(scanner.Text() + "\n")
		output := buffer.String()
		executionContext.Output = output
		if executionContext.Status == constant.PodReady {
			executionContext.Status = constant.RunningState
		}

		_, err = e.sendResponse(executionContext)
		if err != nil {
			log.Error(err, "error in sending execution context")
			return
		}
	}
	output := buffer.String()
	executionContext.Output = output

	if executionContext.Status == constant.PodReady || executionContext.Status == constant.RunningState {
		executionContext.Status = constant.Finished
	}
	_, err = e.sendResponse(executionContext)
	if err != nil {
		log.Error(err, "error in sending execution context")
		return
	}
}

// FetchJob is used to fetch jobs from the controller
func (e *jobServicesClient) FetchJob() (*protofiles.Job, error) {
	start := &protofiles.ExecutorID{ID: e.id}
	return e.grpcClient.FetchJob(start)
}
