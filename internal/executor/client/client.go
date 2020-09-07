package client

import (
	"context"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"time"

	"google.golang.org/grpc"
)

type Client interface {
	Ping(ping *executorCPproto.Ping) (*executorCPproto.HealthResponse, error)
	Register(request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error)
	ConnectClient(cpHost string) error
}

type GrpcClient struct {
	client                executorCPproto.ExecutorCPServicesClient
	connectionTimeoutSecs time.Duration
}

func (g *GrpcClient) ConnectClient(cpHost string) error {
	conn, err := grpc.Dial(cpHost, grpc.WithInsecure())
	if err != nil {
		return err
	}
	grpcClient := executorCPproto.NewExecutorCPServicesClient(conn)
	g.client = grpcClient
	g.connectionTimeoutSecs = time.Second
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
