package daemon

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"octavius/internal/executor/client"
	"octavius/internal/executor/config"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/kubernetes"
	"octavius/internal/pkg/log"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"time"

	v1 "k8s.io/api/core/v1"
)

const (
	Received          = "RECEIVED"
	RequirementNotMet = "REQUIREMENT_NOT_MET"
	Created           = "CREATED"
	CreationFailed    = "CREATION_FAILED"
	JobCreationFailed = "JOB_CREATION_FAILED"
	JobReady          = "JOB_READY"
	PodCreationFailed = "POD_CREATION_FAILED"
	PodReady          = "POD_READY"
	PodFailed         = "POD_FAILED"
	FetchPodLogFailed = "FETCH_POD_LOG_FAILED"
	Finished          = "FINISHED"
)

// Client executor client interface
type Client interface {
	RegisterClient() (bool, error)
	StartClient(executorConfig config.OctaviusExecutorConfig) error
	StartPing()
	FetchJob() (*executorCPproto.Job, error)
	StartKubernetesService()
}

type executorClient struct {
	id                    string
	grpcClient            client.Client
	cpHost                string
	accessToken           string
	connectionTimeoutSecs time.Duration
	pingInterval          time.Duration
	kubernetesClient      kubernetes.KubeClient
	kubeLogWaitTime       time.Duration
}

//NewExecutorClient returns new empty executor client
func NewExecutorClient(grpcClient client.Client) Client {
	return &executorClient{
		grpcClient: grpcClient,
	}
}

func (e *executorClient) StartClient(executorConfig config.OctaviusExecutorConfig) error {
	e.id = executorConfig.ID
	e.cpHost = executorConfig.CPHost
	e.accessToken = executorConfig.AccessToken
	e.connectionTimeoutSecs = executorConfig.ConnTimeOutSec
	e.pingInterval = executorConfig.PingInterval
	e.kubeLogWaitTime = 5 * time.Minute
	err := e.grpcClient.ConnectClient(e.cpHost, e.connectionTimeoutSecs)
	if err != nil {
		return err
	}

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
	e.kubernetesClient, err = kubernetes.NewKubernetesClient(kubeConfig)
	if err != nil {
		return err
	}
	return nil
}

func (e *executorClient) RegisterClient() (bool, error) {
	executorInfo := &executorCPproto.ExecutorInfo{
		Info: e.accessToken,
	}
	registerRequest := &executorCPproto.RegisterRequest{
		ID:           e.id,
		ExecutorInfo: executorInfo,
	}
	res, err := e.grpcClient.Register(registerRequest)
	if err != nil {
		return false, err
	}
	return res.Registered, nil
}

func (e *executorClient) StartPing() {
	for {
		res, err := e.grpcClient.Ping(&executorCPproto.Ping{ID: e.id, State: "stale"})
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		if !res.Recieved {
			log.Error(errors.New("ping not acknowledeged by control plane"), "")
			return
		}
		time.Sleep(e.pingInterval)
	}
}

func (e *executorClient) StartKubernetesService() {
	for {
		job, err := e.FetchJob()
		if err != nil {
			log.Fatal(fmt.Sprintf("error in getting job from server, error details: %s", err.Error()))
			time.Sleep(5 * time.Second)
			continue
		}
		if job.HasJob == "no" {
			time.Sleep(5 * time.Second)
			continue
		}

		log.Info(fmt.Sprintf("recieved job from controller, job details: %+v", job))
		contextID, err := idgen.NewRandomIdGenerator().Generate()
		if err != nil {
			e.sendResponse(&executorCPproto.ExecutionContext{Status: CreationFailed})
			return
		}
		imageName := job.ImageName
		executionArgs := job.JobData
		jobID := job.JobID

		//assign job to kubernetes
		jobContext := executorCPproto.ExecutionContext{
			ExecutionID: contextID,
			JobID:       jobID,
			JobName:     imageName,
			EnvArgs:     executionArgs,
			Status:      Created,
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		executionName, err := e.kubernetesClient.ExecuteJob(ctx, jobID, imageName, executionArgs)
		log.Info(fmt.Sprintln("Executed Job on Kubernetes got ", executionName, " execution jobName and ", err, "errors"))
		if err != nil {
			jobContext.Status = CreationFailed
			e.sendResponse(&jobContext)
			log.Error(err, "error while executing job")
			time.Sleep(10 * time.Second)
			continue
		}

		jobContext.Name = executionName

		e.startWatch(ctx, &jobContext)
		e.sendResponse(&jobContext)
		//get pod logs
		//send pod logs through StreamJobLog
	}
}

func (e *executorClient) sendResponse(jobContext *executorCPproto.ExecutionContext) (*executorCPproto.Acknowledgement, error) {
	return e.grpcClient.SendExecutionContext(jobContext)
}

func (e *executorClient) startWatch(ctx context.Context, executionContext *executorCPproto.ExecutionContext) {

	log.Info(fmt.Sprintf("Start Watch Process for Job With Context ID: %d , name: %s, and Status: %s", executionContext.ExecutionID, executionContext.Name, executionContext.Status))

	err := e.kubernetesClient.WaitForReadyJob(ctx, executionContext.Name, e.kubeLogWaitTime)

	if err != nil {
		executionContext.Status = JobCreationFailed
		return
	}

	executionContext.Status = JobReady
	log.Info(fmt.Sprintf("Job Ready for %d", executionContext.ExecutionID))

	pod, err := e.kubernetesClient.WaitForReadyPod(ctx, executionContext.Name, e.kubeLogWaitTime)
	if err != nil {
		log.Error(err, fmt.Sprintf("wait for ready pod ", pod.Name))
		executionContext.Status = PodCreationFailed
		return
	}
	if pod.Status.Phase == v1.PodFailed {
		executionContext.Status = PodFailed
		log.Info(fmt.Sprintf("Pod Failed for ", executionContext.ExecutionID, " reason: ", pod.Status.Reason, " message: ", pod.Status.Message))
	} else {
		executionContext.Status = PodReady
		log.Info(fmt.Sprintf("Pod Ready for ", executionContext.ExecutionID))
	}

	podLog, err := e.kubernetesClient.GetPodLogs(ctx, pod)
	if err != nil {
		executionContext.Status = FetchPodLogFailed
		return
	}

	scanner := bufio.NewScanner(podLog)
	scanner.Split(bufio.ScanLines)

	var buffer bytes.Buffer
	for scanner.Scan() {
		buffer.WriteString(scanner.Text() + "\n")
	}

	output := string(buffer.Bytes())

	executionContext.Output = output
	log.Info("Execution Output Produced " + string(executionContext.ExecutionID) + " with length " + string(len(output)))

	if executionContext.Status == PodReady {
		executionContext.Status = Finished
	}
	return
}

func (e *executorClient) FetchJob() (*executorCPproto.Job, error) {
	start := &executorCPproto.ExecutorID{Id: e.id}
	return e.grpcClient.FetchJob(start)
}
