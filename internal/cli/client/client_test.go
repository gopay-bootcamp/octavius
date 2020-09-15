package client

import (
	"context"
	"log"
	"net"
	protobuf "octavius/internal/pkg/protofiles/client_cp"
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
	protobuf.RegisterClientCPServicesServer(s, &server{})
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

func (s *server) GetLogs(ctx context.Context, GetLogs *protobuf.RequestForLogs) (*protobuf.Log, error) {
	return &protobuf.Log{
		Log: "sample log 1",
	}, nil
}

func (s *server) DescribeJob(ctx context.Context, describe *protobuf.RequestForDescribe) (*protobuf.Metadata, error) {
	return &protobuf.Metadata{
		Name: "test image",
	}, nil
}

func (s *server) ExecuteJob(ctx context.Context, execute *protobuf.RequestForExecute) (*protobuf.Response, error) {
	return &protobuf.Response{
		Status: "success",
	}, nil
}

func (s *server) PostMetadata(context.Context, *protobuf.RequestToPostMetadata) (*protobuf.MetadataName, error) {
	return &protobuf.MetadataName{
		Name: "name",
	}, nil
}

func (s *server) GetAllMetadata(context.Context, *protobuf.RequestToGetAllMetadata) (*protobuf.MetadataArray, error) {
	return nil, nil
}

func (s *server) GetJobList(context.Context, *protobuf.RequestForGetJobList) (*protobuf.JobList, error) {
	var jobList []string
	jobList = append(jobList, "demo-image-name")
	jobList = append(jobList, "demo-image-name-1")

	response := &protobuf.JobList{
		Jobs: jobList,
	}
	return response, nil
}

// TestCreateMetadata used to test CreateMetadata
func TestCreateMetadata(t *testing.T) {
	createFakeServer()

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	client := protobuf.NewClientCPServicesClient(conn)
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

	client := protobuf.NewClientCPServicesClient(conn)
	testClient := GrpcClient{
		client:                client,
		connectionTimeoutSecs: 10 * time.Second,
	}
	testExecuteRequest := &protobuf.RequestForExecute{}
	res, err := testClient.ExecuteJob(testExecuteRequest)
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

	client := protobuf.NewClientCPServicesClient(conn)
	testClient := GrpcClient{
		client:                client,
		connectionTimeoutSecs: 10 * time.Second,
	}
	testGetRequest := &protobuf.RequestForLogs{}
	res, err := testClient.GetLogs(testGetRequest)

	assert.Nil(t, err)
	assert.Equal(t, res.Log, "sample log 1")
}

func TestGetJobList(t *testing.T) {
	createFakeServer()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	client := protobuf.NewClientCPServicesClient(conn)
	testClient := GrpcClient{
		client:                client,
		connectionTimeoutSecs: 10 * time.Second,
	}

	testGetJobListRequest := &protobuf.RequestForGetJobList{}
	res, err := testClient.GetJobList(testGetJobListRequest)
	assert.Nil(t, err)
	var actual [2]string
	for index, value := range res.Jobs {
		actual[index] = value
	}
	var expected [2]string
	expected[0] = "demo-image-name"
	expected[1] = "demo-image-name-1"

	assert.Equal(t, actual, expected)

}

func TestDescribeJob(t *testing.T) {

	createFakeServer()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	client := protobuf.NewClientCPServicesClient(conn)
	testClient := GrpcClient{
		client:                client,
		connectionTimeoutSecs: 10 * time.Second,
	}

	testDescribeRequest := &protobuf.RequestForDescribe{}
	actual, err := testClient.DescribeJob(testDescribeRequest)
	assert.Nil(t, err)
	assert.Equal(t, "test image", actual.Name)

}
