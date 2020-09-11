package daemon

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"octavius/internal/executor/client"
	"octavius/internal/executor/config"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/kubernetes"
	"octavius/internal/pkg/log"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"time"

	v1 "k8s.io/api/core/v1"
)

// Client executor client interface
type Client interface {
	RegisterClient() (bool, error)
	StartClient(executorConfig config.OctaviusExecutorConfig) error
	StartPing()
	FetchJob() (*executorCPproto.Job, error)
	StartKubernetesService()
	startWatch(executionContext *executorCPproto.ExecutionContext)
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
	state                 string
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
	e.state = constant.IdleState
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
	return err
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
	ticker := time.NewTicker(e.pingInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				res, err := e.grpcClient.Ping(&executorCPproto.Ping{ID: e.id})
				if err != nil {
					log.Fatal(err.Error())
					return
				}
				if !res.Recieved {
					log.Error(errors.New("ping not acknowledeged by control plane"), "")
					return
				}
			}
		}
	}()
}

func (e *executorClient) StartKubernetesService() {
	for {
		job, err := e.FetchJob()
		if err != nil {
			log.Fatal(fmt.Sprintf("error in getting job from server, error details: %s", err.Error()))
		}
		if !job.HasJob {
			time.Sleep(5 * time.Second)
			continue
		}

		e.state = constant.RunningState
		res, err := e.grpcClient.Ping(&executorCPproto.Ping{ID: e.id, State: e.state})
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		if !res.Recieved {
			log.Error(errors.New("ping not acknowledeged by control plane"), "")
			return
		}
		log.Info(fmt.Sprintf("recieved job from controller, job details: %+v", job))
		if err != nil {
			_, err = e.sendResponse(&executorCPproto.ExecutionContext{Status: constant.CreationFailed})
			if err != nil {
				log.Error(err, "error in sending execution context")
			}
			return
		}
		imageName := job.ImageName
		executionArgs := job.JobData
		jobID := job.JobID
		jobContext := executorCPproto.ExecutionContext{
			JobID:      jobID,
			ImageName:  imageName,
			EnvArgs:    executionArgs,
			Status:     constant.Created,
			ExecutorID: e.id,
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

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
		e.state = constant.IdleState
		res, err = e.grpcClient.Ping(&executorCPproto.Ping{ID: e.id, State: e.state})
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		if !res.Recieved {
			log.Error(errors.New("ping not acknowledeged by control plane"), "")
			return
		}
	}
}

func (e *executorClient) sendResponse(jobContext *executorCPproto.ExecutionContext) (*executorCPproto.Acknowledgement, error) {
	return e.grpcClient.SendExecutionContext(jobContext)
}

func (e *executorClient) startWatch(executionContext *executorCPproto.ExecutionContext) {
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
	}

	output := buffer.String()

	executionContext.Output = output
	if executionContext.Status == constant.PodReady {
		executionContext.Status = constant.Finished
	}
	_, err = e.sendResponse(executionContext)
	if err != nil {
		log.Error(err, "error in sending execution context")
		return
	}
}

func (e *executorClient) FetchJob() (*executorCPproto.Job, error) {
	start := &executorCPproto.ExecutorID{Id: e.id}
	return e.grpcClient.FetchJob(start)
}
