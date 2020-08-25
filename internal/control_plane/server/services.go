package server

import (
	"context"
	"errors"
	"fmt"
	"octavius/internal/control_plane/id_generator"
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
	uid := id_generator.NextID()
	ctx = context.WithValue(ctx, "uid", uid)
	name, err := s.procExec.SaveMetadataToDb(ctx, request.Metadata)
	logger.Error(err, fmt.Sprintf("%v Job Create Request Received - Posting Metadata to etcd", uid))
	return name, err
}

func (s *clientCPServicesServer) GetAllMetadata(ctx context.Context, request *protobuf.RequestToGetAllMetadata) (*protobuf.MetadataArray, error) {
	dataList, err := s.procExec.ReadAllMetadata(ctx)
	logger.Error(err, "Getting Metadata")
	return dataList, err
}

func (s *clientCPServicesServer) GetStreamLogs(request *protobuf.RequestForStreamLog, stream protobuf.ClientCPServices_GetStreamLogsServer) error {
	uid := id_generator.NextID()
	logString := &protobuf.Log{RequestId: uid, Log: "lorem ipsum logger logger logger dumb"}
	err := stream.Send(logString)
	logger.Error(err, fmt.Sprintf("%v GetStream Request Received - Sending stream to client", uid))
	return err
}

func (s *clientCPServicesServer) ExecuteJob(ctx context.Context, execute *protobuf.RequestForExecute) (*protobuf.Response, error) {
	//will be utilized after implementation
	//uid := id_generator.NextID()
	return nil, errors.New("not implemented yet")
}
