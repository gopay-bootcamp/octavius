package server

import (
	"context"
	"fmt"
	"octavius/internal/controller/server/execution/metadata"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/util"

	"octavius/internal/pkg/protofiles"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type metadataServicesServer struct {
	procExec metadata.MetadataExecution
	idgen    idgen.RandomIdGenerator
}

// NewMetadataServiceServer used to create a new execution context
func NewMetadataServiceServer(exec metadata.MetadataExecution, idgen idgen.RandomIdGenerator) protofiles.MetadataServicesServer {
	return &metadataServicesServer{
		procExec: exec,
		idgen:    idgen,
	}
}

// Post service is used to save metadata to repository
func (s *metadataServicesServer) Post(ctx context.Context, request *protofiles.RequestToPostMetadata) (*protofiles.MetadataName, error) {
	uuid, err := s.idgen.Generate()
	if err != nil {
		log.Error(err, "error while assigning id to the request")
		return nil, status.Error(codes.Internal, err.Error())
	}

	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	log.Info(fmt.Sprintf("request id: %v, PostMetadata request received with metadata %v", uuid, request.Metadata))

	name, err := s.procExec.SaveMetadata(ctx, request.Metadata)
	if err != nil {
		log.Error(err, fmt.Sprintf("request id: %v, error in saving to etcd", uuid))
		return nil, err
	}
	return name, nil
}

// Describe service is used to fetch metadata for given job
func (s *metadataServicesServer) Describe(ctx context.Context, descriptionData *protofiles.RequestToDescribe) (*protofiles.Metadata, error) {
	uuid, err := s.idgen.Generate()
	if err != nil {
		log.Error(err, "error while assigning id to the request")
		return nil, status.Error(codes.Internal, err.Error())
	}

	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	log.Info(fmt.Sprintf("request ID: %v, DescribeJob request received with name %+v", uuid, descriptionData))
	metadata, err := s.procExec.GetMetadata(ctx, descriptionData)
	if err != nil {
		log.Error(err, "error in fetching metadata of job")
	}
	return metadata, err
}

// List service is used to return list of available jobs
func (s *metadataServicesServer) List(ctx context.Context, request *protofiles.RequestToGetJobList) (*protofiles.JobList, error) {
	uuid, err := s.idgen.Generate()
	if err != nil {
		log.Error(err, "error while assigning id to the request")
		return nil, status.Error(codes.Internal, err.Error())
	}

	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	log.Info(fmt.Sprintf("request ID: %v, GetJobList request received with clientInfo %+v", uuid, request))
	return s.procExec.GetJobList(ctx)
}
