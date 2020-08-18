package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"octavius/pkg/protobuf"
	"time"
)

type Client interface {
	CreateJob(*protobuf.RequestForMetadataPost) (*protobuf.Response, error)
	NewGrpcClient(string2 string) (Client,error)
}



type GrpcClient struct {
	client                protobuf.OctaviusServicesClient
	connectionTimeoutSecs time.Duration
}

func NewClient() Client {
	return &GrpcClient{}
}



func (g *GrpcClient) NewGrpcClient(CPHost string) (Client,error){

	conn, err := grpc.Dial(CPHost, grpc.WithInsecure())
	if err != nil {
		return nil,err
	}
	grpcClient := protobuf.NewOctaviusServicesClient(conn)


	return &GrpcClient{
		client:                grpcClient,
		connectionTimeoutSecs: time.Second,
	},nil

}

func (g *GrpcClient) CreateJob(metadataPostRequest *protobuf.RequestForMetadataPost) (*protobuf.Response, error) {
	res, err := g.client.CreateJob(context.Background(), metadataPostRequest)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}
