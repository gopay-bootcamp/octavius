package client

import (
	"context"
	"fmt"
	"octavius/pkg/protobuf"
	"time"
)

type Client interface {
	CreateJob(*protobuf.RequestForMetadataPost) (*protobuf.Response, error)
}

type grpcClient struct {
	client                protobuf.OctaviusServicesClient
	connectionTimeoutSecs time.Duration
}

func NewGrpcClient(client protobuf.OctaviusServicesClient) Client {
	return &grpcClient{
		client:                client,
		connectionTimeoutSecs: time.Second,
	}
}

func (g *grpcClient) CreateJob(metadataPostRequest *protobuf.RequestForMetadataPost) (*protobuf.Response, error) {
	res, err := g.client.CreateJob(context.Background(), metadataPostRequest)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}
