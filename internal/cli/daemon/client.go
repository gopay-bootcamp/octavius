package daemon

import (
	"octavius/internal/cli/client"
	"octavius/pkg/protobuf"

	"google.golang.org/grpc"
)

var octaviusDClient client.Client

func StartClient() {
	conn, _ := grpc.Dial("localhost:8000", grpc.WithInsecure())
	grpcClient := protobuf.NewOctaviusServicesClient(conn)
	octaviusDClient = client.NewGrpcClient(grpcClient)
}

func CreateJob() {
	octaviusDClient.CreateJob()
}
