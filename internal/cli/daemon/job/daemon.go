package job

import (
	"errors"
	"octavius/internal/cli/client/job"
	"octavius/internal/cli/config"
	"octavius/internal/pkg/protofiles"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client interface {
	Logs(string, job.Client) (*protofiles.Log, error)
	Execute(string, map[string]string, job.Client) (*protofiles.Response, error)
}

type octaviusClient struct {
	octaviusConfigLoader  config.Loader
	jobGrpcClient         job.Client
	CPHost                string
	emailId               string
	accessToken           string
	connectionTimeoutSecs time.Duration
}

func NewClient(clientConfigLoader config.Loader) Client {
	return &octaviusClient{
		octaviusConfigLoader: clientConfigLoader,
	}
}

func (c *octaviusClient) startOctaviusClient(jobGrpcClient job.Client) error {
	octaveConfig, configErr := c.octaviusConfigLoader.Load()
	if configErr != (config.ConfigError{}) {
		return errors.New(configErr.Message)
	}

	c.CPHost = octaveConfig.Host
	c.emailId = octaveConfig.Email
	c.accessToken = octaveConfig.AccessToken
	c.connectionTimeoutSecs = octaveConfig.ConnectionTimeoutSecs
	c.jobGrpcClient = jobGrpcClient

	err := c.jobGrpcClient.ConnectClient(c.CPHost)
	if err != nil {
		return err
	}
	return nil
}

func (c *octaviusClient) Logs(jobID string, jobGrpcClient job.Client) (*protofiles.Log, error) {
	err := c.startOctaviusClient(jobGrpcClient)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	postRequestHeader := protofiles.ClientInfo{
		ClientEmail: c.emailId,
		AccessToken: c.accessToken,
	}
	logRequest := protofiles.RequestToGetLogs{
		ClientInfo: &postRequestHeader,
		JobName:    jobID,
	}
	logResponse, err := c.jobGrpcClient.Logs(&logRequest)
	return logResponse, err
}

func (c *octaviusClient) Execute(jobName string, jobData map[string]string, jobGrpcClient job.Client) (*protofiles.Response, error) {
	err := c.startOctaviusClient(jobGrpcClient)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	postRequestHeader := protofiles.ClientInfo{
		ClientEmail: c.emailId,
		AccessToken: c.accessToken,
	}
	executePostRequest := protofiles.RequestToExecute{
		ClientInfo: &postRequestHeader,
		JobName:    jobName,
		JobData:    jobData,
	}
	return c.jobGrpcClient.Execute(&executePostRequest)
}
