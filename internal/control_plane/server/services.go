package server

import (
	"context"
	"octavius/internal/control_plane/logger"
	"octavius/internal/control_plane/server/execution"
	protobuf "octavius/internal/pkg/protofiles/client_CP"
)

type clientCPServicesServer struct {
	procExec execution.Execution
}

// NewProcServiceServer used to create a new execution context
func NewProcServiceServer(exec execution.Execution) protobuf.ClientCPServicesServer {
	return &clientCPServicesServer{
		procExec: exec,
	}
}

func (s *clientCPServicesServer) PostMetadata(ctx context.Context, request *protobuf.RequestToPostMetadata) (*protobuf.MetadataName, error) {
	name, err := s.procExec.SaveMetadataToDb(ctx, request.Metadata)
	logger.Error(err, "Job Create Request Received - Posting Metadata to etcd")
	return name, err
}

func (s *clientCPServicesServer) GetAllMetadata(ctx context.Context, request *protobuf.RequestToGetAllMetadata) (*protobuf.MetadataArray, error) {
	dataList, err := s.procExec.ReadAllMetadata(ctx)
	logger.Error(err, "Getting Metadata")
	return dataList, err
}

func (s *clientCPServicesServer) GetStreamLogs(request *protobuf.RequestForStreamLog, stream protobuf.ClientCPServices_GetStreamLogsServer) error {
	logString := &protobuf.Log{Log: "lorem ipsum logger logger logger dumb"}
	err := stream.Send(logString)
	logger.Error(err, "GetStream Request Received - Sending stream to client")
	return err
}

func (s *clientCPServicesServer) ExecuteJob(ctx context.Context, execute *protobuf.RequestForExecute) (*protobuf.Response, error) {
	logger.Fatal("Execution is yet to be implemented")
	return nil, nil

}
