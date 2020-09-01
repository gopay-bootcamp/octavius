package daemon

import (
	"errors"
	"fmt"
	"io"
	"octavius/internal/executor/client"
	"octavius/internal/executor/config"
	"octavius/internal/pkg/log"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"time"
)

// Client executor client interface
type Client interface {
	RegisterClient() (bool, error)
	StartClient() error
	StartPing()
	GetJob() (*executorCPproto.Job, error)
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
}

//NewExecutorClient returns new empty executor client
func NewExecutorClient(grpcClient client.Client) Client {
	return &executorClient{
		grpcClient: grpcClient,
	}
}

func (e *executorClient) StartClient() error {
	e.id = config.Config().ID
	e.cpHost = config.Config().CPHost
	e.accessToken = config.Config().AccessToken
	e.connectionTimeoutSecs = config.Config().ConnTimeOutSec
	e.pingInterval = config.Config().PingInterval
	e.jobChan = make(chan *executorCPproto.Job)
	err := e.grpcClient.ConnectClient(e.cpHost)
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

func (e *executorClient) GetJob() (*executorCPproto.Job, error) {
	start := &executorCPproto.Start{Id: e.id}
	return e.grpcClient.GetJob(start)
}

func (e *executorClient) StreamJobLog() {

	logs := []*executorCPproto.JobLog{
		{Log: "success log"},
		{Log: "failed log"},
	}

	for {
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
		time.Sleep(5 * time.Second)
	}
}
