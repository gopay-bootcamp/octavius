package client

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"octavius/pkg/protobuf"
	"time"
)

type MockGrpcClient struct {
	mock.Mock
}



func (m *MockGrpcClient) CreateJob(metadataPostRequest *protobuf.RequestForMetadataPost) (*protobuf.Response, error) {
	args := m.Called(metadataPostRequest)
	return args.Get(0).(*protobuf.Response), args.Error(1)
}

func (m *MockGrpcClient) NewGrpcClient(CPHost string) (Client,error) {
	m.Called(CPHost)
	fmt.Printf("Mock called")
	conn, err := grpc.Dial(CPHost, grpc.WithInsecure())
	if err != nil {
		return nil,err
	}
	grpcClient := protobuf.NewOctaviusServicesClient(conn)


	return &GrpcClient{
		client:                grpcClient,
		connectionTimeoutSecs: time.Second,
	},nil
}


