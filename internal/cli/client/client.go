package client

import (
	"context"
	"fmt"
	"io"
	"octavius/pkg/protobuf"
	"time"
)

type Client interface {
	CreateJob(*protobuf.RequestForMetadataPost) error
	GetStreamLog(*protobuf.RequestForStreamLog) error
	ExecuteJob(*protobuf.RequestForExecute) error
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

func (g *grpcClient) ExecuteJob(requestForExecute *protobuf.RequestForExecute) error {
	res, err := g.client.ExecuteJob(context.Background(), requestForExecute)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
