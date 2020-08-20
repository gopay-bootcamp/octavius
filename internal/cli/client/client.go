package client

import (
	"context"

	"errors"
	"fmt"
	"io"

	"octavius/pkg/protobuf"
	"time"

	"google.golang.org/grpc"
)

type Client interface {
	GetStreamLog(*protobuf.RequestForStreamLog) error
	ExecuteJob(*protobuf.RequestForExecute) error

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

func (g *GrpcClient) GetStreamLog(requestForStreamLog *protobuf.RequestForStreamLog) error {
	responseStream, err := g.client.GetStreamLogs(context.Background(), requestForStreamLog)
	if err != nil {
		return err
	}
	for {
		log, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Printf("log %v", log.Log)
	}
	return nil
}

func (g *GrpcClient) ExecuteJob(requestForExecute *protobuf.RequestForExecute) error {
	res, err := g.client.ExecuteJob(context.Background(), requestForExecute)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
