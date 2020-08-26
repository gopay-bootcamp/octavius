package daemon

import (
	"octavius/internal/executor/client"
	"octavius/internal/executor/config"
	executorCPproto "octavius/internal/pkg/protofiles/executor_CP"
	"time"
)

type Client interface {
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

func (e *executorClient) StartPing() {
	for {
		e.grpcClient.Ping(&executorCPproto.Ping{ID: e.id, State: "stale"})
		time.Sleep(config.Config().PingInterval)
	}

}
