package client

import (
	"context"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type Client interface {
	Ping(ping *executorCPproto.Ping) (*executorCPproto.HealthResponse, error)
	Register(request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error)
	ConnectClient(cpHost string, connectionTimeOut time.Duration) error
	FetchJob(start *executorCPproto.ExecutorID) (*executorCPproto.Job, error)
	SendExecutionContext(executionData *executorCPproto.ExecutionContext) (*executorCPproto.Acknowledgement, error)
}

type GrpcClient struct {
	client                executorCPproto.ExecutorCPServicesClient
	connectionTimeoutSecs time.Duration
}

func (g *GrpcClient) ConnectClient(cpHost string, connectionTimeOut time.Duration) error {
	conn, err := grpc.Dial(cpHost, grpc.WithInsecure())
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	grpcClient := executorCPproto.NewExecutorCPServicesClient(conn)
	g.client = grpcClient
	g.connectionTimeoutSecs = connectionTimeOut
	return nil
}

func (g *GrpcClient) Ping(ping *executorCPproto.Ping) (*executorCPproto.HealthResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()

	return g.client.HealthCheck(ctx, ping)
}

func (g *GrpcClient) Register(request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()

	return g.client.Register(ctx, request)
}

func (g *GrpcClient) FetchJob(request *executorCPproto.ExecutorID) (*executorCPproto.Job, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()
	return g.client.FetchJob(ctx, request)
}

func (g *GrpcClient) SendExecutionContext(executionData *executorCPproto.ExecutionContext) (*executorCPproto.Acknowledgement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()
	return g.client.SendExecutionContext(ctx, executionData)
}
