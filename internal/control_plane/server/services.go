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
	if err != nil {
		logger.Log.Err(err).Msg("error in posting metadata")
		return name, err
	}
	logger.Log.Info().Msg("Metadata posted")
	return name, nil
}

func (s *octaviusServiceServer) GetAllMetadata(ctx context.Context, request *procProto.RequestToGetAllMetadata) (*procProto.MetadataArray, error) {
	dataList, err := s.procExec.ReadAllMetadata(ctx)
	if err != nil {
		logger.Log.Err(err).Msg("error in getting metadata list")
		return dataList, err
	}
	logger.Log.Info().Msgf("Getting Metadata")
	return dataList, nil
}
