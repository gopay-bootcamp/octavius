package daemon

import (
	"errors"
	"octavius/internal/executor/client"
	"octavius/internal/executor/config"
	"octavius/internal/pkg/kubernetes"
	"octavius/internal/pkg/log"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
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
	kubernetesClient      kubernetes.KubeClient
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
	e.pingInterval = config.Config().PingInterval
	err := e.grpcClient.ConnectClient(e.cpHost)
	if err != nil {
		return err
	}

	var kubeConfig = config.OctaviusExecutorConfig{
		KubeConfig:                   config.Config().KubeConfig,
		KubeContext:                  config.Config().KubeContext,
		DefaultNamespace:             config.Config().DefaultNamespace,
		KubeServiceAccountName:       config.Config().KubeServiceAccountName,
		JobPodAnnotations:            config.Config().JobPodAnnotations,
		KubeJobActiveDeadlineSeconds: config.Config().KubeJobActiveDeadlineSeconds,
		KubeJobRetries:               config.Config().KubeJobRetries,
		KubeWaitForResourcePollCount: config.Config().KubeWaitForResourcePollCount,
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
