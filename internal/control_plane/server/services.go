package server

import (
	"context"
	"octavius/internal/control_plane/logger"
	"octavius/internal/control_plane/server/execution"
	procProto "octavius/pkg/protobuf"
)

type octaviusServiceServer struct {
	procExec execution.Execution
}

// NewProcServiceServer used to create a new execution context
func NewProcServiceServer(exec execution.Execution) procProto.OctaviusServicesServer {
	return &octaviusServiceServer{
		procExec: exec,
	}
}

func (s *octaviusServiceServer) PostMetadata(ctx context.Context, request *procProto.RequestToPostMetadata) (*procProto.MetadataName, error) {
	name, err := s.procExec.SaveMetadataToDb(ctx, request.Metadata)
	logger.Error(err, "Job Create Request Received - Posting Metadata to etcd")
	return name, err
}

func (s *octaviusServiceServer) GetAllMetadata(ctx context.Context, request *procProto.RequestToGetAllMetadata) (*procProto.MetadataArray, error) {
	dataList, err := s.procExec.ReadAllMetadata(ctx)
	logger.Error(err, "Getting Metadata")
	return dataList, err
}

func (s *octaviusServiceServer) GetStreamLogs(request *procProto.RequestForStreamLog, stream procProto.OctaviusServices_GetStreamLogsServer) error {
	logString := &procProto.Log{Log: "lorem ipsum logger logger logger dumb"}
	err := stream.Send(logString)
	logger.Error(err, "GetStream Request Received - Sending stream to client")
	return err
}

func (s *octaviusServiceServer) ExecuteJob(ctx context.Context, execute *procProto.RequestForExecute) (*procProto.Response, error) {
	logger.Fatal("Execution is yet to be implemented")
	return nil, nil
}
