package client

import (
	"context"
	"fmt"
	"octavius/pkg/protobuf"
	"time"
)

type Client interface {
	PostMetadata(*protobuf.RequestToPostMetadata) error
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

func (g *grpcClient) PostMetadata(metadataPostRequest *protobuf.RequestToPostMetadata) error {
	res, err := g.client.PostMetadata(context.Background(), metadataPostRequest)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
