package client

import (
	"context"
	octerr "octavius/internal/pkg/errors"
	executorCPproto "octavius/internal/pkg/protofiles/executor_CP"
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
		return octerr.New(2, err)
	}
	grpcClient := executorCPproto.NewExecutorCPServicesClient(conn)
	g.client = grpcClient
	g.connectionTimeoutSecs = time.Second
	return nil
}

func (g *GrpcClient) Ping(ping *executorCPproto.Ping) (*executorCPproto.HealthResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()
	res, err := g.client.HealthCheck(ctx, ping)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (g *GrpcClient) Register(request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()
	res, err := g.client.Register(ctx, request)
	if err != nil {
		return nil, err
	}
	return res, err
}
