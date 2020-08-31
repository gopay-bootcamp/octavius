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
	StartStream()
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

func (e *executorClient) StartStream() {
	clientStream, err := e.grpcClient.Stream(&executorCPproto.Start{Id: e.id})
	if err != nil {
		log.Error(err, "error starting executor job stream")
		return
	}
	for {
		jobDetails, err := clientStream.Recv()
		if err == io.EOF {
			log.Error(err, "server stream closed")
			return
		}
		if err != nil {
			log.Error(err, "error in server stream")
			return
		}
		fmt.Println(jobDetails.JobName)
	}
}
