package daemon

import (
	"errors"
	"io"
	"octavius/internal/cli/client"
	"octavius/internal/cli/config"
	protobuf "octavius/internal/pkg/protofiles/client_cp"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client interface {
	CreateMetadata(io.Reader, client.Client) (*protobuf.MetadataName, error)
	GetStreamLog(string, client.Client) (*[]protobuf.Log, error)
	ExecuteJob(string, map[string]string, client.Client) (*protobuf.Response, error)
	GetJobList(client.Client) (*protobuf.JobList, error)
}

type octaviusClient struct {
	octaviusConfigLoader  config.Loader
	grpcClient            client.Client
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

func (c *octaviusClient) startOctaviusClient(grpcClient client.Client) error {
	octaveConfig, configErr := c.octaviusConfigLoader.Load()
	if configErr != (config.ConfigError{}) {
		return errors.New(configErr.Message)
	}

	c.CPHost = octaveConfig.Host
	c.emailId = octaveConfig.Email
	c.accessToken = octaveConfig.AccessToken
	c.connectionTimeoutSecs = octaveConfig.ConnectionTimeoutSecs
	c.grpcClient = grpcClient

	err := c.grpcClient.ConnectClient(c.CPHost)
	if err != nil {
		return err
	}
	return nil
}

// CreateMetadata take metadata file handler and grpc client
func (c *octaviusClient) CreateMetadata(metadataFileHandler io.Reader, grpcClient client.Client) (*protobuf.MetadataName, error) {
	metadata := protobuf.Metadata{}
	err := jsonpb.Unmarshal(metadataFileHandler, &metadata)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = c.startOctaviusClient(grpcClient)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	postRequestHeader := protobuf.ClientInfo{
		ClientEmail: c.emailId,
		AccessToken: c.accessToken,
	}
	metadataPostRequest := protobuf.RequestToPostMetadata{
		Metadata:   &metadata,
		ClientInfo: &postRequestHeader,
	}

	res, err := c.grpcClient.CreateMetadata(&metadataPostRequest)
	return res, err
}

func (c *octaviusClient) GetStreamLog(jobName string, grpcClient client.Client) (*[]protobuf.Log, error) {
	err := c.startOctaviusClient(grpcClient)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	postRequestHeader := protobuf.ClientInfo{
		ClientEmail: c.emailId,
		AccessToken: c.accessToken,
	}
	getStreamPostRequest := protobuf.RequestForStreamLog{
		ClientInfo: &postRequestHeader,
		JobName:    jobName,
	}
	logResponse, err := c.grpcClient.GetStreamLog(&getStreamPostRequest)
	return logResponse, err
}
func (c *octaviusClient) ExecuteJob(jobName string, jobData map[string]string, grpcClient client.Client) (*protobuf.Response, error) {
	err := c.startOctaviusClient(grpcClient)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	postRequestHeader := protobuf.ClientInfo{
		ClientEmail: c.emailId,
		AccessToken: c.accessToken,
	}
	executePostRequest := protobuf.RequestForExecute{
		ClientInfo: &postRequestHeader,
		JobName:    jobName,
		JobData:    jobData,
	}
	return c.grpcClient.ExecuteJob(&executePostRequest)
}

// GetJobList takes grpcClient as argument and returns list of available jobs
func (c *octaviusClient) GetJobList(grpcClient client.Client) (*protobuf.JobList, error) {
	err := c.startOctaviusClient(grpcClient)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	postRequestHeader := protobuf.ClientInfo{
		ClientEmail: c.emailId,
		AccessToken: c.accessToken,
	}
	listJobPostRequest := protobuf.RequestForGetJobList{
		ClientInfo: &postRequestHeader,
	}

	return c.grpcClient.GetJobList(&listJobPostRequest)

}
