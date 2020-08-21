package server

import (
	"context"
	"octavius/internal/control_plane/server/execution"
	"octavius/internal/logger"
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
	if err != nil {
		logger.Error("error in posting metadata", err)
		return name, err
	}
	return name, nil
}

func (s *octaviusServiceServer) GetAllMetadata(ctx context.Context, request *procProto.RequestToGetAllMetadata) (*procProto.MetadataArray, error) {
	dataList, err := s.procExec.ReadAllMetadata(ctx)
	if err != nil {
		logger.Error("error in getting metadata list", err)
		return dataList, err
	}
	return dataList, nil
}

func (s *octaviusServiceServer) GetStreamLogs( request *procProto.RequestForStreamLog, server procProto.OctaviusServices_GetStreamLogsServer) (error) {
	return nil
}


func (s *octaviusServiceServer) ExecuteJob(ctx context.Context, execute *procProto.RequestForExecute) (*procProto.Response, error) {
	panic("implement me")
}
