package execution

import (
	"context"
	"octavius/internal/control_plane/server/metadata/repository"
	clientCPproto "octavius/internal/pkg/protofiles/client_CP"
)

// Execution interface for methods related to execution
type Execution interface {
	SaveMetadataToDb(ctx context.Context, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error)
	ReadAllMetadata(ctx context.Context) (*clientCPproto.MetadataArray, error)
}

type execution struct {
	metadata repository.MetadataRepository
}

// NewExec creates a new instance of metadata respository
func NewExec(metadataRepo repository.MetadataRepository) Execution {
	return &execution{
		metadata: metadataRepo,
	}
}

//SaveMetadataToDb calls the metadata/repository Save() function and returns MetadataName
func (e *execution) SaveMetadataToDb(ctx context.Context, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error) {
	result, err := e.metadata.Save(ctx, metadata.Name, metadata)
	return result, err
}

//ReadAllMetadata calls the metadata/repository GetAll() and returns MetadataArray
func (e *execution) ReadAllMetadata(ctx context.Context) (*clientCPproto.MetadataArray, error) {
	result, err := e.metadata.GetAll(ctx)
	return result, err
}
