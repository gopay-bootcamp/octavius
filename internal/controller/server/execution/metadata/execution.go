// Package metadata implements metadata related functions
package metadata

import (
	"context"

	metadataRepo "octavius/internal/controller/server/repository/metadata"
	"octavius/internal/pkg/protofiles"
)

// MetadataExecution interface for methods related to metadata
type MetadataExecution interface {
	SaveMetadata(ctx context.Context, metadata *protofiles.Metadata) (*protofiles.MetadataName, error)
	GetMetadata(ctx context.Context, request *protofiles.RequestToDescribe) (*protofiles.Metadata, error)
	GetJobList(ctx context.Context) (*protofiles.JobList, error)
}

type metadataExecution struct {
	metadataRepo metadataRepo.Repository
}

// NewMetadataExec creates a new instance of metadata respository
func NewMetadataExec(metadataRepo metadataRepo.Repository) *metadataExecution {
	return &metadataExecution{
		metadataRepo: metadataRepo,
	}
}

// SaveMetadata calls the repository/metadata Save() function and returns MetadataName
func (e *metadataExecution) SaveMetadata(ctx context.Context, metadata *protofiles.Metadata) (*protofiles.MetadataName, error) {
	return e.metadataRepo.SaveMetadata(ctx, metadata.Name, metadata)
}

// GetMetadata calls the repository/metadata GetValue() and returns Metadata
func (e *metadataExecution) GetMetadata(ctx context.Context, request *protofiles.RequestToDescribe) (*protofiles.Metadata, error) {
	return e.metadataRepo.GetMetadata(ctx, request.JobName)
}

// GetJobList function will call metadata repository and return list of available jobs
func (e *metadataExecution) GetJobList(ctx context.Context) (*protofiles.JobList, error) {
	return e.metadataRepo.GetAvailableJobs(ctx)
}
