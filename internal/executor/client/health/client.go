package health

import (
	"context"
	"octavius/internal/pkg/protofiles"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client interface {
	Ping(ping *protofiles.Ping) (*protofiles.HealthResponse, error)
	ConnectClient(cpHost string, connectionTimeOut time.Duration) error
}

type GrpcClient struct {
	client                protofiles.HealthServicesClient
	connectionTimeoutSecs time.Duration
}

func (g *GrpcClient) ConnectClient(cpHost string, connectionTimeOut time.Duration) error {
	conn, err := grpc.Dial(cpHost, grpc.WithInsecure())
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	grpcClient := protofiles.NewHealthServicesClient(conn)
	g.client = grpcClient
	g.connectionTimeoutSecs = connectionTimeOut
	return nil
}

func (g *GrpcClient) Ping(ping *protofiles.Ping) (*protofiles.HealthResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()

	return g.client.Check(ctx, ping)
}
