package client

import (
	"context"
	"log"
	"net"
	"octavius/pkg/protobuf"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const (
	bufSize = 1024 * 1024
)

var testMetadataArray = &protobuf.MetadataArray{}

var lis *bufconn.Listener


func createFakeServer() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	protobuf.RegisterOctaviusServicesServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

type server struct{}

func (s *server) GetStreamLogs(streamLog *protobuf.RequestForStreamLog, logsServer protobuf.OctaviusServices_GetStreamLogsServer) (error) {
	logsServer.Send(&protobuf.Log{Log: "Test log 1"})
	logsServer.Send(&protobuf.Log{Log: "Test log 2"})
	return nil
}

func (s *server) ExecuteJob(ctx context.Context, execute *protobuf.RequestForExecute) (*protobuf.Response, error) {
	return &protobuf.Response{
		Status: "success",
	},nil
}

func (s *server) PostMetadata(context.Context, *protobuf.RequestToPostMetadata) (*protobuf.MetadataName, error) {
	return &protobuf.MetadataName{
		Name: "name",
	}, nil
}

func (s *server) GetAllMetadata(context.Context, *protobuf.RequestToGetAllMetadata) (*protobuf.MetadataArray, error) {
	return nil, nil
}

// TestCreateMetadata used to test CreateMetadata
func TestCreateMetadata(t *testing.T) {
	createFakeServer()

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	client := protobuf.NewOctaviusServicesClient(conn)
	testClient := GrpcClient{
		client:                client,
		connectionTimeoutSecs: 10 * time.Second,
	}
	testPostRequest := &protobuf.RequestToPostMetadata{}
	res, err := testClient.CreateMetadata(testPostRequest)
	assert.Nil(t, err)
	assert.Equal(t, "name", res.Name)
}

func TestExecuteJob(t *testing.T) {
	createFakeServer()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	client := protobuf.NewOctaviusServicesClient(conn)
	testClient := GrpcClient{
		client:                client,
		connectionTimeoutSecs: 10 * time.Second,
	}
	testExecuteRequest := &protobuf.RequestForExecute{}
	res, err := testClient.ExecuteJob(testExecuteRequest)
	assert.Nil(t, err)
	assert.Equal(t, "success", res.Status)
}

func TestGetStream(t *testing.T) {
	createFakeServer()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	client := protobuf.NewOctaviusServicesClient(conn)
	testClient := GrpcClient{
		client:                client,
		connectionTimeoutSecs: 10 * time.Second,
	}
	testGetStreamRequest := &protobuf.RequestForStreamLog{}
	res, err := testClient.GetStreamLog(testGetStreamRequest)
	assert.Nil(t, err)
	var actual [2]string
	for index, value := range *res {
		actual[index] = value.Log
	}
	var expected [2]string
	expected[0] = "Test log 1"
	expected[1] = "Test log 2"

	assert.Equal(t, actual, expected)
}
