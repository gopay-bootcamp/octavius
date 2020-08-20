package client

import (
	"context"
	"errors"
	"octavius/pkg/protobuf"
	"time"

	"google.golang.org/grpc"
)

type Client interface {
	CreateMetadata(*protobuf.RequestToPostMetadata) (*protobuf.MetadataName, error)
	ConnectClient(cpHost string) error
}

type GrpcClient struct {
	client                protobuf.OctaviusServicesClient
	connectionTimeoutSecs time.Duration
}

func (g *GrpcClient) ConnectClient(cpHost string) error {
	conn, err := grpc.Dial(cpHost, grpc.WithInsecure())
	if err != nil {
		return errors.New("error dialing to CP host server")
	}
	grpcClient := protobuf.NewOctaviusServicesClient(conn)
	g.client = grpcClient
	g.connectionTimeoutSecs = time.Second
	return nil
}

func (g *GrpcClient) CreateMetadata(metadataPostRequest *protobuf.RequestToPostMetadata) (*protobuf.MetadataName, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()
	res, err := g.client.PostMetadata(ctx, metadataPostRequest)
	if err != nil {
		return nil, err
	}
	return res, nil
}
