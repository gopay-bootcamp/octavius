package client

import (
	"context"
	"errors"
	"fmt"
	"octavius/pkg/protobuf"
	"time"

	"google.golang.org/grpc"
)

type Client interface {
	CreateJob(*protobuf.RequestForMetadataPost) (*protobuf.Response, error)
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

func (g *GrpcClient) CreateJob(metadataPostRequest *protobuf.RequestForMetadataPost) (*protobuf.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()
	res, err := g.client.CreateJob(ctx, metadataPostRequest)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}
