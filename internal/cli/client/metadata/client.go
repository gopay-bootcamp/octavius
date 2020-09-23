//Package metadata implements methods to send metadata related gRPC requests to controller
package metadata

import (
	"context"
	"octavius/internal/pkg/protofiles"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

//Client interface defines metadata related methods
type Client interface {
	Post(*protofiles.RequestToPostMetadata) (*protofiles.MetadataName, error)
	ConnectClient(cpHost string) error
	Describe(*protofiles.RequestToDescribe) (*protofiles.Metadata, error)
	List(*protofiles.RequestToGetJobList) (*protofiles.JobList, error)
}

//GrpcClient structure represents metadata related gRPC client
type GrpcClient struct {
	client                protofiles.MetadataServicesClient
	connectionTimeoutSecs time.Duration
}

//ConnectClient function dials a connection provided host between controller and client
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

//Post function sends metadata to controller and returns metadata name as response
func (g *GrpcClient) Post(metadataPostRequest *protofiles.RequestToPostMetadata) (*protofiles.MetadataName, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()

	res, err := g.client.Post(ctx, metadataPostRequest)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//Describe function sends request to controller to describe a specific job to controller and returns metadata of that job
func (g *GrpcClient) Describe(requestForDescribe *protofiles.RequestToDescribe) (*protofiles.Metadata, error) {
	return g.client.Describe(context.Background(), requestForDescribe)
}

//List function sends request to controller to get list of all available jobs and returns job list
func (g *GrpcClient) List(requestForGetJobList *protofiles.RequestToGetJobList) (*protofiles.JobList, error) {
	return g.client.List(context.Background(), requestForGetJobList)
}
