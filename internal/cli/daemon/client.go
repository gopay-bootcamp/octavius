package daemon

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"octavius/pkg/protobuf"

	"io"
	"octavius/internal/cli/client"
	"octavius/internal/cli/config"
	"time"
)

type Client interface {
	ConfigureClient() error
	CreateMetadata(io.Reader,client.Client) (*protobuf.Response, error)
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

func (c *octaviusClient) ConfigureClient() error {
	err := c.loadOctaviusConfig()
	if err != nil {
		return err
	}


	return nil
}

func (c *octaviusClient) loadOctaviusConfig() error {

	octaveConfig, err := c.octaviusConfigLoader.Load()
	if err != (config.ConfigError{}) {
		return errors.New(err.Message)
	}

	c.CPHost = octaveConfig.Host
	c.emailId = octaveConfig.Email
	c.accessToken = octaveConfig.AccessToken
	c.connectionTimeoutSecs = octaveConfig.ConnectionTimeoutSecs
	return nil
}



func (c *octaviusClient) CreateMetadata(metadataFileHandler io.Reader,client client.Client) (*protobuf.Response, error) {




	metadata := protobuf.Metadata{}
	err := jsonpb.Unmarshal(metadataFileHandler, &metadata)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("Error unmarshalling metadata.json file: ", err))
	}


	err = c.ConfigureClient()
	if err != nil {
		fmt.Printf("%v",err)
		return nil, err
	}

	c.grpcClient,err=client.NewGrpcClient(c.CPHost)
	if err != nil {
		fmt.Printf("%v",err)
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
