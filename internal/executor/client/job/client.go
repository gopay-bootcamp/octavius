package job

import (
	"context"
	"octavius/internal/pkg/protofiles"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//Client interface declares the executor's job functions
type Client interface {
	ConnectClient(cpHost string, connectionTimeOut time.Duration) error
	FetchJob(start *protofiles.ExecutorID) (*protofiles.Job, error)
	SendExecutionContext(executionData *protofiles.ExecutionContext) (*protofiles.Acknowledgement, error)
	PostExecutorStatus(stat *protofiles.Status) (*protofiles.Acknowledgement, error)
}

//GrpcClient struct holds the data required to invoke executor's job functions
type GrpcClient struct {
	client                protofiles.JobServiceClient
	connectionTimeoutSecs time.Duration
}

//ConnectClient establishes connection between JobServiceClient and controller
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

//FetchJob fetches pending job for the executor
func (g *GrpcClient) FetchJob(request *protofiles.ExecutorID) (*protofiles.Job, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()
	return g.client.Get(ctx, request)
}

//SendExecutionContext sends the job execution data over to the controller and gets acknowledgement fromt eh controller
func (g *GrpcClient) SendExecutionContext(executionData *protofiles.ExecutionContext) (*protofiles.Acknowledgement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()
	return g.client.PostExecutionData(ctx, executionData)
}

// PostExecutorStatus sends the executor status to controller
func (g *GrpcClient) PostExecutorStatus(stat *protofiles.Status) (*protofiles.Acknowledgement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), g.connectionTimeoutSecs)
	defer cancel()
	return g.client.PostExecutorStatus(ctx, stat)
}
