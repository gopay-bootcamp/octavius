package client

import (
	"octavius/pkg/protobuf"
	"time"
)
type Client interface {
	CreateJob() ()
}

type grpcClient struct {
	client protobuf.OctaviusServicesClient
	connectionTimeoutSecs time.Duration
}


func NewGrpcClient(client protobuf.OctaviusServicesClient) Client{
	return &grpcClient{
		client: client,
		connectionTimeoutSecs: time.Second,
	}
}

func (g *grpcClient) CreateJob() {
	panic("implement me")
}




