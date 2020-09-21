package registration

import (
	client "octavius/internal/executor/client/registration"
	"octavius/internal/executor/config"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/kubernetes"
	"octavius/internal/pkg/protofiles"
	"sync"
	"time"
)

// Client executor client interface
type RegistrationServicesClient interface {
	RegisterClient(executorConfig config.OctaviusExecutorConfig) (bool, error)
	connectClient(executorConfig config.OctaviusExecutorConfig) error
}

type registrationServicesClient struct {
	id                    string
	cpHost                string
	accessToken           string
	grpcClient            client.Client
	connectionTimeoutSecs time.Duration
	pingInterval          time.Duration
	kubernetesClient      kubernetes.KubeClient
	kubeLogWaitTime       time.Duration
	state                 string
	statusLock            sync.RWMutex
}

//NewregistrationServicesClient returns new empty executor client
func NewRegistrationServicesClient(grpcClient client.Client) RegistrationServicesClient {
	return &registrationServicesClient{
		grpcClient: grpcClient,
	}
}

func (e *registrationServicesClient) connectClient(executorConfig config.OctaviusExecutorConfig) error {
	e.id = executorConfig.ID
	e.cpHost = executorConfig.CPHost
	e.accessToken = executorConfig.AccessToken
	e.connectionTimeoutSecs = executorConfig.ConnTimeOutSec
	e.pingInterval = executorConfig.PingInterval
	e.state = constant.IdleState
	err := e.grpcClient.ConnectClient(e.cpHost, e.connectionTimeoutSecs)
	if err != nil {
		return err
	}
	return nil
}

func (e *registrationServicesClient) RegisterClient(executorConfig config.OctaviusExecutorConfig) (bool, error) {
	err := e.connectClient(executorConfig)
	if err != nil {
		return false, err
	}
	executorInfo := &protofiles.ExecutorInfo{
		Info: e.accessToken,
	}
	registerRequest := &protofiles.RegisterRequest{
		ID:           e.id,
		ExecutorInfo: executorInfo,
	}
	res, err := e.grpcClient.Register(registerRequest)
	if err != nil {
		return false, err
	}
	return res.Registered, nil
}
