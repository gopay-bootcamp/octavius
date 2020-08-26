package client

import (
	"context"
	"io"

	octerr "octavius/internal/pkg/errors"
	clientCPproto "octavius/internal/pkg/protofiles/client_CP"
	"time"

	"google.golang.org/grpc"
)

type Client interface {
	GetStreamLog(*clientCPproto.RequestForStreamLog) (*[]clientCPproto.Log, error)
	ExecuteJob(*clientCPproto.RequestForExecute) (*clientCPproto.Response, error)

	CreateMetadata(*clientCPproto.RequestToPostMetadata) (*clientCPproto.MetadataName, error)
	ConnectClient(cpHost string) error
}

type GrpcClient struct {
	client                clientCPproto.ClientCPServicesClient
	connectionTimeoutSecs time.Duration
}

func (g *GrpcClient) ConnectClient(cpHost string) error {
	conn, err := grpc.Dial(cpHost, grpc.WithInsecure())
	if err != nil {
		return octerr.New(2, err)
	}
	grpcClient := clientCPproto.NewClientCPServicesClient(conn)
	g.client = grpcClient
	g.connectionTimeoutSecs = time.Second
	return nil
}

func (g *GrpcClient) CreateMetadata(metadataPostRequest *clientCPproto.RequestToPostMetadata) (*clientCPproto.MetadataName, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()
	res, err := g.client.PostMetadata(ctx, metadataPostRequest)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (g *GrpcClient) GetStreamLog(requestForStreamLog *clientCPproto.RequestForStreamLog) (*[]clientCPproto.Log, error) {
	responseStream, err := g.client.GetStreamLogs(context.Background(), requestForStreamLog)
	if err != nil {
		return nil, err
	}
	var logResponse []clientCPproto.Log
	for {
		log, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		logResponse = append(logResponse, *log)
	}
	return &logResponse, nil
}

func (g *GrpcClient) ExecuteJob(requestForExecute *clientCPproto.RequestForExecute) (*clientCPproto.Response, error) {
	res, err := g.client.ExecuteJob(context.Background(), requestForExecute)
	if err != nil {
		return nil, err
	}
	return res, nil
}
