package metadata

import (
	"context"

	metadataRepo "octavius/internal/controller/server/repository/metadata"
	"octavius/internal/controller/server/scheduler"
	"octavius/internal/pkg/protofiles"
)

// Execution interface for methods related to metadata
type MetadataExecution interface {
	SaveMetadata(ctx context.Context, metadata *protofiles.Metadata) (*protofiles.MetadataName, error)
	GetMetadata(ctx context.Context, request *protofiles.RequestToDescribe) (*protofiles.Metadata, error)
	ReadAllMetadata(ctx context.Context) (*protofiles.MetadataArray, error)
	GetJobList(ctx context.Context) (*protofiles.JobList, error)
}

type metadataExecution struct {
	metadataRepo metadataRepo.Repository
	scheduler    scheduler.Scheduler
}

// NewExec creates a new instance of metadata respository
func NewMetadataExec(metadataRepo metadataRepo.Repository, scheduler scheduler.Scheduler) *metadataExecution {
	return &metadataExecution{
		metadataRepo: metadataRepo,
		scheduler:    scheduler,
	}
}

//SaveMetadata calls the repository/metadata Save() function and returns MetadataName
func (e *metadataExecution) SaveMetadata(ctx context.Context, metadata *protofiles.Metadata) (*protofiles.MetadataName, error) {
	return e.metadataRepo.Save(ctx, metadata.Name, metadata)
}

//GetMetadata calls the repository/metadata GetValue() and returns Metadata
func (e *metadataExecution) GetMetadata(ctx context.Context, request *protofiles.RequestToDescribe) (*protofiles.Metadata, error) {
	return e.metadataRepo.GetValue(ctx, request.JobName)
}

//ReadAllMetadata calls the repository/metadata GetAll() and returns MetadataArray
func (e *metadataExecution) ReadAllMetadata(ctx context.Context) (*protofiles.MetadataArray, error) {
	return e.metadataRepo.GetAll(ctx)
}

// GetJobList function will call metadata repository and return list of available jobs
func (e *metadataExecution) GetJobList(ctx context.Context) (*protofiles.JobList, error) {
	return e.metadataRepo.GetAvailableJobList(ctx)
}
