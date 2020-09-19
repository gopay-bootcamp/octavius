package metadata

import (
	"context"
	"octavius/internal/pkg/protofiles"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type Client interface {
	Post(*protofiles.RequestToPostMetadata) (*protofiles.MetadataName, error)
	ConnectClient(cpHost string) error
	Describe(*protofiles.RequestToDescribe) (*protofiles.Metadata, error)
	List(*protofiles.RequestToGetJobList) (*protofiles.JobList, error)
}

type GrpcClient struct {
	client                protofiles.MetadataServicesClient
	connectionTimeoutSecs time.Duration
}

func (g *GrpcClient) ConnectClient(cpHost string) error {
	conn, err := grpc.Dial(cpHost, grpc.WithInsecure())
	if err != nil {
		return status.Error(codes.Unavailable, err.Error())
	}
	grpcClient := protofiles.NewMetadataServicesClient(conn)
	g.client = grpcClient
	g.connectionTimeoutSecs = time.Second
	return nil
}

func (g *GrpcClient) Post(metadataPostRequest *protofiles.RequestToPostMetadata) (*protofiles.MetadataName, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()

	res, err := g.client.Post(ctx, metadataPostRequest)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (g *GrpcClient) Describe(requestForDescribe *protofiles.RequestToDescribe) (*protofiles.Metadata, error) {
	return g.client.Describe(context.Background(), requestForDescribe)
}

func (g *GrpcClient) List(requestForGetJobList *protofiles.RequestToGetJobList) (*protofiles.JobList, error) {
	return g.client.List(context.Background(), requestForGetJobList)
}
