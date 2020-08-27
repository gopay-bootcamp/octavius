package daemon

import (
	"fmt"
	"octavius/internal/executor/client"
	"octavius/internal/executor/config"
	"octavius/internal/executor/logger"
	executorCPproto "octavius/internal/pkg/protofiles/executor_CP"
	"time"
)

type Client interface {
	RegisterClient() (bool, error)
	StartClient() error
	StartPing()
}

type executorClient struct {
	id                    string
	grpcClient            client.Client
	cpHost                string
	accessToken           string
	connectionTimeoutSecs time.Duration
	pingInterval          time.Duration
}

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
	regsiterRequest := &executorCPproto.RegisterRequest{
		ID:           e.id,
		ExecutorInfo: executorInfo,
	}
	res, err := e.grpcClient.Register(regsiterRequest)
	if err != nil {
		return false, err
	}
	return res.Registered, nil
}

func (e *executorClient) StartPing() {
	logger.Info("starting ping")
	for {
		res, err := e.grpcClient.Ping(&executorCPproto.Ping{ID: e.id, State: "stale"})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v \n", res)
		time.Sleep(config.Config().PingInterval)
	}
}
