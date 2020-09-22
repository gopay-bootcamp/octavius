package job

import (
	"context"
	"log"
	"net"
	"octavius/internal/pkg/protofiles"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const (
	bufSize = 1024 * 1024
)

var lis *bufconn.Listener

type server struct{}

func createFakeServer() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	protofiles.RegisterJobServiceServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func (s *server) Logs(ctx context.Context, GetLogs *protofiles.RequestToGetLogs) (*protofiles.Log, error) {
	return &protofiles.Log{
		Log: "sample log 1",
	}, nil
}

func (s *server) Execute(ctx context.Context, execute *protofiles.RequestToExecute) (*protofiles.Response, error) {
	return &protofiles.Response{
		Status: "success",
	}, nil
}

func (s *server) Get(context.Context, *protofiles.ExecutorID) (*protofiles.Job, error) {
	return nil, nil
}
func (s *server) PostExecutionData(context.Context, *protofiles.ExecutionContext) (*protofiles.Acknowledgement, error) {
	return nil, nil
}

func TestExecuteJob(t *testing.T) {
	createFakeServer()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	client := protofiles.NewJobServiceClient(conn)
	testClient := GrpcClient{
		client:                client,
		connectionTimeoutSecs: 10 * time.Second,
	}
	testExecuteRequest := &protofiles.RequestToExecute{}
	res, err := testClient.Execute(testExecuteRequest)
	assert.Nil(t, err)
	assert.Equal(t, "success", res.Status)
}

func TestGetLogs(t *testing.T) {
	createFakeServer()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	client := protofiles.NewJobServiceClient(conn)
	testClient := GrpcClient{
		client:                client,
		connectionTimeoutSecs: 10 * time.Second,
	}
	testGetRequest := &protofiles.RequestToGetLogs{}
	res, err := testClient.Logs(testGetRequest)

	assert.Nil(t, err)
	assert.Equal(t, res.Log, "sample log 1")
}

func TestConnectClient(t *testing.T) {
	createFakeServer()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	client := protofiles.NewJobServiceClient(conn)
	testClient := GrpcClient{
		client:                client,
		connectionTimeoutSecs: 10 * time.Second,
	}

	err = testClient.ConnectClient(lis.Addr().String())
	assert.Nil(t, err)
}
