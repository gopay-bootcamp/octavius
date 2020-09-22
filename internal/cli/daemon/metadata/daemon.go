package metadata

import (
	"errors"
	"io"
	"net/http"
	"octavius/internal/cli/client/metadata"
	"octavius/internal/cli/config"
	"octavius/internal/pkg/protofiles"
	"strings"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client interface {
	Post(io.Reader, metadata.Client) (*protofiles.MetadataName, error)
	Describe(string, metadata.Client) (*protofiles.Metadata, error)
	List(metadata.Client) (*protofiles.JobList, error)
}

type octaviusClient struct {
	octaviusConfigLoader  config.Loader
	grpcClient            metadata.Client
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

func (c *octaviusClient) startOctaviusClient(grpcClient metadata.Client) error {
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

func validateImageNameDocker(imageName string) (bool, error) {
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
	} else if response.StatusCode == 404 {
		return false, status.Error(codes.NotFound, "image not found in docker hub")
	}
	return false, nil
}

// CreateMetadata take metadata file handler and grpc client
func (c *octaviusClient) Post(metadataFileHandler io.Reader, grpcClient metadata.Client) (*protofiles.MetadataName, error) {
	metadata := protofiles.Metadata{}
	err := jsonpb.Unmarshal(metadataFileHandler, &metadata)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if metadata.Name == "" {
		return nil, status.Error(codes.Internal, "job-name should not be empty in metadata")
	}

	isImageExist, err := validateImageNameDocker(metadata.ImageName)
	if err != nil {
		return nil, err
	}
	if metadata.ImageName == "" || !isImageExist {
		return nil, status.Error(codes.Internal, "image-name should be valid in metadata")
	}
	if metadata.EnvVars == nil {
		return nil, status.Error(codes.Internal, "there should be envVars field in metadata")
	}

	err = c.startOctaviusClient(grpcClient)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	postRequestHeader := protofiles.ClientInfo{
		ClientEmail: c.emailId,
		AccessToken: c.accessToken,
	}
	metadataPostRequest := protofiles.RequestToPostMetadata{
		Metadata:   &metadata,
		ClientInfo: &postRequestHeader,
	}

	res, err := c.grpcClient.Post(&metadataPostRequest)
	return res, err
}

func (c *octaviusClient) Describe(jobName string, grpcClient metadata.Client) (*protofiles.Metadata, error) {
	err := c.startOctaviusClient(grpcClient)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	postRequestHeader := protofiles.ClientInfo{
		ClientEmail: c.emailId,
		AccessToken: c.accessToken,
	}
	descriptionPostRequest := protofiles.RequestToDescribe{
		ClientInfo: &postRequestHeader,
		JobName:    jobName,
	}

	return c.grpcClient.Describe(&descriptionPostRequest)
}

// GetJobList takes jobGrpcClient as argument and returns list of available jobs
func (c *octaviusClient) List(grpcClient metadata.Client) (*protofiles.JobList, error) {
	err := c.startOctaviusClient(grpcClient)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	postRequestHeader := protofiles.ClientInfo{
		ClientEmail: c.emailId,
		AccessToken: c.accessToken,
	}
	listJobPostRequest := protofiles.RequestToGetJobList{
		ClientInfo: &postRequestHeader,
	}
	return c.grpcClient.List(&listJobPostRequest)
}
