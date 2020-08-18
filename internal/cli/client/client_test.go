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

var testPostResponse = &protobuf.Response{
	Status: "success",
}

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

func (s *server) CreateJob(context.Context, *protobuf.RequestForMetadataPost) (*protobuf.Response, error) {
	return testPostResponse, nil
}

func TestCreateJob(t *testing.T) {
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
	testPostRequest := &protobuf.RequestForMetadataPost{}
	res, err := testClient.CreateJob(testPostRequest)
	assert.Nil(t, err)
	assert.Equal(t, testPostResponse.Status, res.Status)
}
