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
	return nil, nil
}

func (s *server) Execute(ctx context.Context, execute *protofiles.RequestToExecute) (*protofiles.Response, error) {
	return nil, nil
}

func (s *server) Get(context.Context, *protofiles.ExecutorID) (*protofiles.Job, error) {
	return &protofiles.Job{
		HasJob:    true,
		JobID:     "test-id",
		ImageName: "test-image-name",
	}, nil
}
func (s *server) PostExecutionData(context.Context, *protofiles.ExecutionContext) (*protofiles.Acknowledgement, error) {
	return &protofiles.Acknowledgement{
		Recieved: true,
	}, nil
}

func (s *server) PostExecutorStatus(context.Context, *protofiles.Status) (*protofiles.Acknowledgement, error) {
	return &protofiles.Acknowledgement{
		Recieved: true,
	}, nil
}

func init() {
	createFakeServer()
}

func TestConnectClient(t *testing.T) {

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

	err = testClient.ConnectClient(lis.Addr().String(), 10*time.Second)
	assert.Nil(t, err)
}
func TestFetchJob(t *testing.T) {

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
	testExecutorID := &protofiles.ExecutorID{}
	res, err := testClient.FetchJob(testExecutorID)
	assert.Nil(t, err)
	assert.Equal(t, res.HasJob, true)
	assert.Equal(t, res.JobID, "test-id")
}

func TestSendExecutionContext(t *testing.T) {

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
	testExecutionContext := &protofiles.ExecutionContext{}
	res, err := testClient.SendExecutionContext(testExecutionContext)
	assert.Nil(t, err)
	assert.Equal(t, res.Recieved, true)
}
