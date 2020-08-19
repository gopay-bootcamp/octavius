package daemon

import (
	"errors"
	"fmt"
	"io"
	"octavius/internal/cli/client"
	"octavius/internal/cli/config"
	"octavius/pkg/protobuf"
	"time"

	"github.com/golang/protobuf/jsonpb"
)

type Client interface {
	CreateMetadata(io.Reader, client.Client) (*protobuf.Response, error)
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

func (c *octaviusClient) CreateMetadata(metadataFileHandler io.Reader, grpcClient client.Client) (*protobuf.Response, error) {
	metadata := protobuf.Metadata{}
	err := jsonpb.Unmarshal(metadataFileHandler, &metadata)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("Error unmarshalling metadata.json file: ", err))
	}

	err = c.startOctaviusClient(grpcClient)
	if err != nil {
		return nil, err
	}

	postRequestHeader := protobuf.ClientInfo{
		ClientEmail: c.emailId,
		AccessToken: c.accessToken,
	}
	metadataPostRequest := protobuf.RequestForMetadataPost{
		Metadata:   &metadata,
		ClientInfo: &postRequestHeader,
	}

	res, err := c.grpcClient.CreateJob(&metadataPostRequest)
	if err != nil {
		return nil, errors.New("Error occured when sending the grpc request. Check your CPHost")
	}
	return res, nil
}
