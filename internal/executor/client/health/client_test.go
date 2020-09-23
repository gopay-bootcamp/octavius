package health

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
	protofiles.RegisterHealthServicesServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func (s *server) Check(ctx context.Context, ping *protofiles.Ping) (*protofiles.HealthResponse, error) {
	return &protofiles.HealthResponse{
		Recieved: true,
	}, nil
}

func TestConnectClient(t *testing.T) {
	createFakeServer()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	client := protofiles.NewHealthServicesClient(conn)
	testClient := GrpcClient{
		client:                client,
		connectionTimeoutSecs: 10 * time.Second,
	}

	err = testClient.ConnectClient(lis.Addr().String(), 10*time.Second)
	assert.Nil(t, err)
}

func TestPing(t *testing.T) {
	createFakeServer()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	client := protofiles.NewHealthServicesClient(conn)
	testClient := GrpcClient{
		client:                client,
		connectionTimeoutSecs: 10 * time.Second,
	}
	testPing := &protofiles.Ping{}
	res, err := testClient.Ping(testPing)
	assert.Nil(t, err)
	assert.Equal(t, res.Recieved, true)
}
