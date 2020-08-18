package server

import (
	"context"
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

func (s *octaviusServiceServer) PostMetadata(ctx context.Context, request *procProto.RequestToPostMetadata) (*procProto.MetadataName,error) {
	id := s.procExec.SaveMetadataToDb(ctx, request.Metadata)
	return id,nil
}

func (s *octaviusServiceServer) GetAllMetadata(ctx context.Context, request *procProto.RequestToGetAllMetadata) (*procProto.MetadataArray,error) {
	dataList:= s.procExec.ReadAllMetadata(ctx)
	return dataList,nil
}
