package job

import (
	"context"
	"octavius/internal/pkg/protofiles"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client interface {
	ConnectClient(cpHost string, connectionTimeOut time.Duration) error
	FetchJob(start *protofiles.ExecutorID) (*protofiles.Job, error)
	SendExecutionContext(executionData *protofiles.ExecutionContext) (*protofiles.Acknowledgement, error)
}

type GrpcClient struct {
	client                protofiles.JobServiceClient
	connectionTimeoutSecs time.Duration
}

func (g *GrpcClient) ConnectClient(cpHost string, connectionTimeOut time.Duration) error {
	conn, err := grpc.Dial(cpHost, grpc.WithInsecure())
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	grpcClient := protofiles.NewJobServiceClient(conn)
	g.client = grpcClient
	g.connectionTimeoutSecs = connectionTimeOut
	return nil
}

func (g *GrpcClient) FetchJob(request *protofiles.ExecutorID) (*protofiles.Job, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()
	return g.client.Get(ctx, request)
}

func (g *GrpcClient) SendExecutionContext(executionData *protofiles.ExecutionContext) (*protofiles.Acknowledgement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()
	return g.client.PostExecutionData(ctx, executionData)
}
