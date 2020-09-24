package health

import (
	"errors"
	client "octavius/internal/executor/client/health"
	"octavius/internal/executor/config"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/protofiles"
	"sync"
	"time"
)

// Client executor client interface
type HealthServicesClient interface {
	connectClient(executorConfig config.OctaviusExecutorConfig) error
	StartPing(executorConfig config.OctaviusExecutorConfig)
}

type healthServicesClient struct {
	id                    string
	grpcClient            client.Client
	cpHost                string
	accessToken           string
	connectionTimeoutSecs time.Duration
	pingInterval          time.Duration
	statusLock            sync.RWMutex
}

//NewHealthServicesClient returns new empty executor client
func NewHealthServicesClient(grpcClient client.Client) HealthServicesClient {
	return &healthServicesClient{
		grpcClient: grpcClient,
	}
}

func (e *healthServicesClient) connectClient(executorConfig config.OctaviusExecutorConfig) error {
	e.id = executorConfig.ID
	e.cpHost = executorConfig.CPHost
	e.accessToken = executorConfig.AccessToken
	e.connectionTimeoutSecs = executorConfig.ConnTimeOutSec
	e.pingInterval = executorConfig.PingInterval
	err := e.grpcClient.ConnectClient(e.cpHost, e.connectionTimeoutSecs)
	if err != nil {
		return err
	}
	return err
}

func (e *healthServicesClient) StartPing(executorConfig config.OctaviusExecutorConfig) {
	err := e.connectClient(executorConfig)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	ticker := time.NewTicker(e.pingInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				e.statusLock.RLock()
				res, err := e.grpcClient.Ping(&protofiles.Ping{ID: e.id})
				e.statusLock.RUnlock()
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
