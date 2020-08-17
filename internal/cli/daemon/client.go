package daemon

import (
	"errors"
	"fmt"
	"io/ioutil"
	"octavius/internal/cli/client"
	"octavius/internal/cli/config"
	"octavius/pkg/protobuf"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"google.golang.org/grpc"
)

type Client interface {
	StartClient() error
	CreateMetadata(metadataFile string) error
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

func (c *octaviusClient) StartClient() error {
	err := c.loadOctaviusConfig()
	if err != nil {
		return err
	}

	conn, err := grpc.Dial(c.CPHost, grpc.WithInsecure())
	if err != nil {
		return err
	}
	grpcClient := protobuf.NewOctaviusServicesClient(conn)
	client := client.NewGrpcClient(grpcClient)
	c.grpcClient = client
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

func (c *octaviusClient) CreateMetadata(metadataFile string) error {
	metadataJson, err := ioutil.ReadFile(metadataFile)
	if err != nil {
		return errors.New(fmt.Sprintln("Error reading metadata file: ", metadataFile))
	}

	metadata := protobuf.Metadata{}
	//find a better method for umarshalling using io reader
	err = jsonpb.UnmarshalString(string(metadataJson), &metadata)
	if err != nil {
		return errors.New(fmt.Sprintln("Error unmarshalling metadata.json file: ", err))
	}

	err = c.StartClient()
	if err != nil {
		return err
	}

	postRequestHeader := protobuf.ClientInfo{
		ClientEmail: c.emailId,
		AcessToken: c.accessToken,
	}
	metadataPostRequest := protobuf.RequestForMetadataPost{
		Metadata:   &metadata,
		ClientInfo: &postRequestHeader,
	}

	err = c.grpcClient.CreateJob(&metadataPostRequest)
	if err != nil {
		return errors.New("Error occured when sending the grpc request. Check your CPHost")
	}
	return nil
}
