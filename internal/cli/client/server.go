package daemon

import (
	"context"
	"fmt"
	"net"
	"octavius/pkg/protobuf"

	"google.golang.org/grpc"
)

type server struct {
}

func (s server) CreateJob(ctx context.Context, metadata *protobuf.Metadata) (*protobuf.Response, error) {
	return &protobuf.Response{Status: "success"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8000")

	srvr := grpc.NewServer()

	protobuf.RegisterProcServiceServer(srvr, &server{})

	if err == nil {
		fmt.Println("Server running successfully....")
		srvr.Serve(listener)
	}
}
