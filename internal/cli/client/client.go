package client

import (
	"context"
	"fmt"
	"octavius/pkg/protobuf"
	"time"
)

type Client interface {
	CreateJob(*protobuf.RequestForMetadataPost) error
	GetStreamLog(*protobuf.RequestForStreamLog) error
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
	res, err := g.client.CreateJob(context.Background(), metadataPostRequest)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (g *grpcClient) GetStreamLog(requestForStreamLog *protobuf.RequestForStreamLog) error {
	res, err := g.client.
	if err != nil {
		return err
	}
	fmt.Println(res)

	return nil
}
