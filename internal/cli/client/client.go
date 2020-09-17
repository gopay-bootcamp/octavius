package client

import (
	"context"
	protobuf "octavius/internal/pkg/protofiles/client_cp"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type Client interface {
	GetLogs(*protobuf.RequestForLogs) (*protobuf.Log, error)
	ExecuteJob(*protobuf.RequestForExecute) (*protobuf.Response, error)
	CreateMetadata(*protobuf.RequestToPostMetadata) (*protobuf.MetadataName, error)
	ConnectClient(cpHost string) error
	GetJobList(*protobuf.RequestForGetJobList) (*protobuf.JobList, error)
	DescribeJob(*protobuf.RequestForDescribe) (*protobuf.Metadata, error)
}

type GrpcClient struct {
	metadataClient        protobuf.MetadataServicesClient
	jobClient             protobuf.JobServicesClient
	connectionTimeoutSecs time.Duration
}

func (g *GrpcClient) ConnectClient(cpHost string) error {
	conn, err := grpc.Dial(cpHost, grpc.WithInsecure())
	if err != nil {
		return status.Error(codes.Unavailable, err.Error())
	}
	grpcMetadataClient := protobuf.NewMetadataServicesClient(conn)
	grpcJobClient := protobuf.NewJobServicesClient(conn)
	g.jobClient = grpcJobClient
	g.metadataClient = grpcMetadataClient
	g.connectionTimeoutSecs = time.Second
	return nil
}

func (g *GrpcClient) CreateMetadata(metadataPostRequest *protobuf.RequestToPostMetadata) (*protobuf.MetadataName, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()

	res, err := g.metadataClient.PostMetadata(ctx, metadataPostRequest)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (g *GrpcClient) GetLogs(requestForLogs *protobuf.RequestForLogs) (*protobuf.Log, error) {
	return g.jobClient.GetLogs(context.Background(), requestForLogs)
}

func (g *GrpcClient) ExecuteJob(requestForExecute *protobuf.RequestForExecute) (*protobuf.Response, error) {
	return g.jobClient.ExecuteJob(context.Background(), requestForExecute)
}

func (g *GrpcClient) GetJobList(requestForGetJobList *protobuf.RequestForGetJobList) (*protobuf.JobList, error) {
	return g.jobClient.GetJobList(context.Background(), requestForGetJobList)
}

func (g *GrpcClient) DescribeJob(requestForDescribe *protobuf.RequestForDescribe) (*protobuf.Metadata, error) {
	return g.jobClient.DescribeJob(context.Background(), requestForDescribe)
}
