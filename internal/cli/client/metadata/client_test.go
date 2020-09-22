package metadata

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

type server struct{}

var lis *bufconn.Listener

func createFakeServer() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	protofiles.RegisterMetadataServicesServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func (s *server) Describe(ctx context.Context, describe *protofiles.RequestToDescribe) (*protofiles.Metadata, error) {
	return &protofiles.Metadata{
		Name: "test image",
	}, nil
}

func (s *server) Post(context.Context, *protofiles.RequestToPostMetadata) (*protofiles.MetadataName, error) {
	return &protofiles.MetadataName{
		Name: "name",
	}, nil
}

func (s *server) List(context.Context, *protofiles.RequestToGetJobList) (*protofiles.JobList, error) {
	var jobList []string
	jobList = append(jobList, "demo-image-name")
	jobList = append(jobList, "demo-image-name-1")

	response := &protofiles.JobList{
		Jobs: jobList,
	}
	return response, nil
}

// TestPost used to test Post
func TestPost(t *testing.T) {
	createFakeServer()

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	client := protofiles.NewMetadataServicesClient(conn)
	testClient := GrpcClient{
		client:                client,
		connectionTimeoutSecs: 10 * time.Second,
	}
	testPostRequest := &protofiles.RequestToPostMetadata{}
	res, err := testClient.Post(testPostRequest)
	assert.Nil(t, err)
	assert.Equal(t, "name", res.Name)
}

func TestDescribe(t *testing.T) {

	createFakeServer()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	client := protofiles.NewMetadataServicesClient(conn)
	testClient := GrpcClient{
		client:                client,
		connectionTimeoutSecs: 10 * time.Second,
	}

	testDescribeRequest := &protofiles.RequestToDescribe{}
	actual, err := testClient.Describe(testDescribeRequest)
	assert.Nil(t, err)
	assert.Equal(t, "test image", actual.Name)

}
func TestList(t *testing.T) {
	createFakeServer()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	client := protofiles.NewMetadataServicesClient(conn)
	testClient := GrpcClient{
		client:                client,
		connectionTimeoutSecs: 10 * time.Second,
	}

	testGetJobListRequest := &protofiles.RequestToGetJobList{}
	res, err := testClient.List(testGetJobListRequest)
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

func TestConnectClient(t *testing.T) {
	createFakeServer()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	client := protofiles.NewMetadataServicesClient(conn)
	testClient := GrpcClient{
		client:                client,
		connectionTimeoutSecs: 10 * time.Second,
	}

	err = testClient.ConnectClient(lis.Addr().String())
	assert.Nil(t, err)
}
