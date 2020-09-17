package daemon

import (
	"errors"
	"io"
	"net/http"
	"octavius/internal/cli/client"
	"octavius/internal/cli/config"
	protobuf "octavius/internal/pkg/protofiles/client_cp"
	"strings"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client interface {
	CreateMetadata(io.Reader, client.Client) (*protobuf.MetadataName, error)
	GetLogs(string, client.Client) (*protobuf.Log, error)
	ExecuteJob(string, map[string]string, client.Client) (*protobuf.Response, error)
	GetJobList(client.Client) (*protobuf.JobList, error)
	DescribeJob(string, client.Client) (*protobuf.Metadata, error)
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

func validateImageName(imageName string) (bool, error) {
	splitedImageName := strings.Split(imageName, "/")
	splitedimageTag := strings.Split(imageName, ":")
	imageNameWithoutTag := splitedimageTag[0]
	var tag string

	if len(splitedimageTag) == 1 {
		tag = "latest"
	} else if len(splitedimageTag) == 2 {
		tag = splitedimageTag[1]
	} else {
		return false, status.Error(codes.Internal, "invalid image tag")
	}

	var url string
	if len(splitedImageName) == 1 {
		url = "https://hub.docker.com/v2/repositories/library/" + imageNameWithoutTag + "/tags/" + tag
	} else if len(splitedImageName) == 2 {
		url = "https://hub.docker.com/v2/repositories/" + imageNameWithoutTag + "/tags/" + tag
	} else {
		return false, status.Error(codes.Internal, "invalid image name")
	}

	response, err := http.Get(url)
	if err != nil {
		return false, status.Error(codes.Internal, "error in validating image name")
	}
	if response.StatusCode == 200 {
		return true, nil
	}
	return false, nil
}

// CreateMetadata take metadata file handler and grpc client
func (c *octaviusClient) CreateMetadata(metadataFileHandler io.Reader, grpcClient client.Client) (*protobuf.MetadataName, error) {
	metadata := protobuf.Metadata{}
	err := jsonpb.Unmarshal(metadataFileHandler, &metadata)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if metadata.Name == "" {
		return nil, status.Error(codes.Internal, "job-name should not be empty in metadata")
	}

	doesExist, err := validateImageName(metadata.ImageName)
	if err != nil {
		return nil, err
	}
	if metadata.ImageName == "" || doesExist == false {
		return nil, status.Error(codes.Internal, "image-name should be valid in metadata")
	}
	if metadata.EnvVars == nil {
		return nil, status.Error(codes.Internal, "there should be envVars field in metadata")
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

func (c *octaviusClient) GetLogs(jobID string, grpcClient client.Client) (*protobuf.Log, error) {
	err := c.startOctaviusClient(grpcClient)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	postRequestHeader := protobuf.ClientInfo{
		ClientEmail: c.emailId,
		AccessToken: c.accessToken,
	}
	logRequest := protobuf.RequestForLogs{
		ClientInfo: &postRequestHeader,
		JobName:    jobID,
	}
	logResponse, err := c.grpcClient.GetLogs(&logRequest)
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

func (c *octaviusClient) DescribeJob(jobName string, grpcClient client.Client) (*protobuf.Metadata, error) {
	err := c.startOctaviusClient(grpcClient)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	postRequestHeader := protobuf.ClientInfo{
		ClientEmail: c.emailId,
		AccessToken: c.accessToken,
	}
	descriptionPostRequest := protobuf.RequestForDescribe{
		ClientInfo: &postRequestHeader,
		JobName:    jobName,
	}

	return c.grpcClient.DescribeJob(&descriptionPostRequest)
}
