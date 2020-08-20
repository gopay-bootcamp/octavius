package server

import (
	"context"
	"github.com/rs/zerolog"
	"octavius/internal/control_plane/server/execution"
	procProto "octavius/pkg/protobuf"
)

type octaviusServiceServer struct {
	procExec execution.Execution
	logger zerolog.Logger
}

// NewProcServiceServer used to create a new execution context
func NewProcServiceServer(exec execution.Execution, logger zerolog.Logger) procProto.OctaviusServicesServer {
	return &octaviusServiceServer{
		procExec: exec,
		logger: logger,
	}
}

func (s *octaviusServiceServer) PostMetadata(ctx context.Context, request *procProto.RequestToPostMetadata) (*procProto.MetadataName, error) {
	name, err := s.procExec.SaveMetadataToDb(ctx, request.Metadata)
	if err != nil {
		s.logger.Err(err).Msg("error in posting metadata")
		return name, err
	}
	s.logger.Info().Msg("Metadata posted")
	return name, nil
}

func (s *octaviusServiceServer) GetAllMetadata(ctx context.Context, request *procProto.RequestToGetAllMetadata) (*procProto.MetadataArray, error) {
	dataList, err := s.procExec.ReadAllMetadata(ctx)
	if err != nil {
		s.logger.Err(err).Msg("error in getting metadata list")
		return dataList, err
	}
	s.logger.Info().Msgf("Getting Metadata")
	return dataList, nil
}
