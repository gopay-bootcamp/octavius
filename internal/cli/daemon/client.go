package daemon

import (
	"errors"
	"fmt"
	"io"
	"octavius/internal/cli/client"
	"octavius/internal/cli/config"
	clientCPproto "octavius/internal/pkg/protofiles/client_CP"
	"time"

	"github.com/gogo/protobuf/jsonpb"
)

type Client interface {
	CreateMetadata(io.Reader, client.Client) (*clientCPproto.MetadataName, error)
	GetStreamLog(string, client.Client) (*[]clientCPproto.Log, error)
	ExecuteJob(string, map[string]string, client.Client) (*clientCPproto.Response, error)
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
func (c *octaviusClient) CreateMetadata(metadataFileHandler io.Reader, grpcClient client.Client) (*clientCPproto.MetadataName, error) {
	metadata := clientCPproto.Metadata{}
	err := jsonpb.Unmarshal(metadataFileHandler, &metadata)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("Error unmarshalling metadata.json file: ", err))
	}

	err = c.startOctaviusClient(grpcClient)
	if err != nil {
		return nil, err
	}

	postRequestHeader := clientCPproto.ClientInfo{
		ClientEmail: c.emailId,
		AccessToken: c.accessToken,
	}
	metadataPostRequest := clientCPproto.RequestToPostMetadata{
		Metadata:   &metadata,
		ClientInfo: &postRequestHeader,
	}

	res, err := c.grpcClient.CreateMetadata(&metadataPostRequest)
	return res, err
}

func (c *octaviusClient) GetStreamLog(jobName string, grpcClient client.Client) (*[]clientCPproto.Log, error) {
	err := c.startOctaviusClient(grpcClient)
	if err != nil {
		return nil, err
	}

	postRequestHeader := clientCPproto.ClientInfo{
		ClientEmail: c.emailId,
		AccessToken: c.accessToken,
	}
	getStreamPostRequest := clientCPproto.RequestForStreamLog{
		ClientInfo: &postRequestHeader,
		JobName:    jobName,
	}
	logResponse, err := c.grpcClient.GetStreamLog(&getStreamPostRequest)
	if err != nil {
		return nil, errors.New("error occured when sending the grpc request. Check your CPHost")
	}
	return logResponse, nil
}
func (c *octaviusClient) ExecuteJob(jobName string, jobData map[string]string, grpcClient client.Client) (*clientCPproto.Response, error) {
	err := c.startOctaviusClient(grpcClient)
	if err != nil {
		return nil, err
	}
	postRequestHeader := clientCPproto.ClientInfo{
		ClientEmail: c.emailId,
		AccessToken: c.accessToken,
	}
	executePostRequest := clientCPproto.RequestForExecute{
		ClientInfo: &postRequestHeader,
		JobName:    jobName,
		JobData:    jobData,
	}
	response, err := c.grpcClient.ExecuteJob(&executePostRequest)
	if err != nil {
		return nil, errors.New("error occured when sending the grpc request. Check your CPHost")

	}
	return response, nil
}
