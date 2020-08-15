package client

import (
	"context"
	"fmt"
	"octavius/pkg/protobuf"
	"time"
)

type Client interface {
	CreateJob(*protobuf.RequestForMetadataPost) error
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

func (g *grpcClient) CreateJob(metadataPostRequest *protobuf.RequestForMetadataPost) error {
	fmt.Println("Creating job at: ")
	res, err := g.client.CreateJob(context.Background(), metadataPostRequest)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
