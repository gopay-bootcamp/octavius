package daemon

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"octavius/internal/executor/client"
	"octavius/internal/executor/config"
	"octavius/internal/pkg/kubernetes"
	"octavius/internal/pkg/log"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"time"

	"github.com/klauspost/compress/gzip"
	v1 "k8s.io/api/core/v1"
)

// Client executor client interface
type Client interface {
	RegisterClient() (bool, error)
	StartClient(executorConfig config.OctaviusExecutorConfig) error
	StartPing()
	GetJob() (*executorCPproto.Job, error)
	StartKubernetesService()
	StreamJobLog()
}

type executorClient struct {
	id                    string
	grpcClient            client.Client
	cpHost                string
	accessToken           string
	connectionTimeoutSecs time.Duration
	pingInterval          time.Duration
	jobChan               chan *executorCPproto.Job
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
			log.Error(err, "error in ping")
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
		job, err := e.GetJob()
		if err != nil {
			log.Fatal(fmt.Sprintf("error in getting job from server, error details: %s", err.Error()))
		}

		if !job.HasJob {
			time.Sleep(5 * time.Second)
			continue
		}
		log.Info(fmt.Sprintf("recieved job from controller, job details: %+v", job))
		time.Sleep(5 * time.Second)
		//assign job to kubernetes
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		executionName, err := e.kubernetesClient.ExecuteJob(ctx, job.JobID, job.ImageName, job.JobData)
		log.Info(fmt.Sprint("Executed Job on Kubernetes got ", executionName, " execution jobName and ", err, "error"))
		if err == nil {
			go startWatch(ctx, e, executionName)
		}

		//get pod logs
		//send pod logs through StreamJobLog
	}
}

func startWatch(ctx context.Context, e *executorClient, executionName string) {

	err := e.kubernetesClient.WaitForReadyJob(ctx, executionName, e.kubeLogWaitTime)

	if err != nil {
		return
	}
	pod, err := e.kubernetesClient.WaitForReadyPod(ctx, executionName, e.kubeLogWaitTime)
	log.Error(err, fmt.Sprintln("wait for ready pod", pod))
	if err != nil {
		return
	}
	if pod.Status.Phase == v1.PodFailed {
		log.Info(fmt.Sprintln("Pod Failed for ", executionName, " reason: ", pod.Status.Reason, " message: ", pod.Status.Message))
	} else {
		log.Info(fmt.Sprintln("Pod Ready for ", executionName))
	}

	podLog, err := e.kubernetesClient.GetPodLogs(ctx, pod)
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(podLog)
	scanner.Split(bufio.ScanLines)

	var buffer bytes.Buffer
	for scanner.Scan() {
		buffer.WriteString(scanner.Text() + "\n")
	}

	gzip.
		log.Info("Execution Output Produced " + executionName + " with length ")
	reader, err := gzip.NewReader(podLog)
	return
}

func (e *executorClient) GetJob() (*executorCPproto.Job, error) {
	start := &executorCPproto.Start{Id: e.id}
	return e.grpcClient.GetJob(start)
}

func (e *executorClient) StreamJobLog() {

	logs := []*executorCPproto.JobLog{
		{Log: "success log"},
		{Log: "failed log"},
	}

	logStream, err := e.grpcClient.StreamLog()
	if err != nil {
		log.Error(err, "error setting up job log stream")
		return
	}
	for _, jobLog := range logs {
		if err := logStream.Send(jobLog); err != nil {
			if err == io.EOF {
				break
			}
			log.Error(err, "error streaming log")
			return
		}
	}
	logSummary, _ := logStream.CloseAndRecv()
	fmt.Println(logSummary)
}
