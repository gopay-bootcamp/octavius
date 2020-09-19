package job

import (
	"context"
	"octavius/internal/pkg/protofiles"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type Client interface {
	Logs(*protofiles.RequestToGetLogs) (*protofiles.Log, error)
	Execute(*protofiles.RequestToExecute) (*protofiles.Response, error)
	ConnectClient(cpHost string) error
}

type GrpcClient struct {
	client                protofiles.JobServiceClient
	connectionTimeoutSecs time.Duration
}

func (g *GrpcClient) ConnectClient(cpHost string) error {
	conn, err := grpc.Dial(cpHost, grpc.WithInsecure())
	if err != nil {
		return status.Error(codes.Unavailable, err.Error())
	}
	grpcClient := protofiles.NewJobServiceClient(conn)
	g.client = grpcClient
	g.connectionTimeoutSecs = time.Second
	return nil
}

func (g *GrpcClient) Logs(requestForLogs *protofiles.RequestToGetLogs) (*protofiles.Log, error) {
	return g.client.Logs(context.Background(), requestForLogs)
}

func (g *GrpcClient) Execute(requestForExecute *protofiles.RequestToExecute) (*protofiles.Response, error) {
	return g.client.Execute(context.Background(), requestForExecute)
}
