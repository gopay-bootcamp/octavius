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

func (s *server) Describe(ctx context.Context, describe *protofiles.RequestForDescribe) (*protofiles.Metadata, error) {
	return &protofiles.Metadata{
		Name: "test image",
	}, nil
}

func (s *server) Post(context.Context, *protofiles.RequestToPostMetadata) (*protofiles.MetadataName, error) {
	return &protofiles.MetadataName{
		Name: "name",
	}, nil
}

// TestCreateMetadata used to test CreateMetadata
func TestCreateMetadata(t *testing.T) {
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

func TestDescribeJob(t *testing.T) {

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

	testDescribeRequest := &protofiles.RequestForDescribe{}
	actual, err := testClient.Describe(testDescribeRequest)
	assert.Nil(t, err)
	assert.Equal(t, "test image", actual.Name)

}
