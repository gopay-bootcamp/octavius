package client

import (
	"context"
	"fmt"
	"net"
	"octavius/pkg/protobuf"

	"google.golang.org/grpc"
)

type server struct {
}

func (s server) CreateJob(ctx context.Context, metadata *protobuf.RequestForMetadataPost) (*protobuf.Response, error) {
	return &protobuf.Response{Status: "success"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8000")

	srvr := grpc.NewServer()

	protobuf.RegisterOctaviusServicesServer(srvr, &server{})

	if err == nil {
		fmt.Println("Server running successfully....")
		srvr.Serve(listener)
	}
}
