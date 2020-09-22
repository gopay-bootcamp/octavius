package health

import (
	"context"
	"octavius/internal/pkg/protofiles"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//Client interface declares the executor's health functions
type Client interface {
	Ping(ping *protofiles.Ping) (*protofiles.HealthResponse, error)
	ConnectClient(cpHost string, connectionTimeOut time.Duration) error
}

//GrpcClient struct holds the data required to invoke executor's health related functions
type GrpcClient struct {
	client                protofiles.HealthServicesClient
	connectionTimeoutSecs time.Duration
}

//ConnectClient establishes connection between HealthServicesClient and controller
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

//Ping sends executor's Ping to the controller
func (g *GrpcClient) Ping(ping *protofiles.Ping) (*protofiles.HealthResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()

	return g.client.Check(ctx, ping)
}
