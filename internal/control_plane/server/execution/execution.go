package execution

import (
	"context"
	"octavius/internal/control_plane/server/metadata/repository"
	"octavius/pkg/protobuf"
)

// Execution interface for methods related to execution
type Execution interface {
	SaveMetadataToDb(ctx context.Context, metadata *protobuf.Metadata) (*protobuf.MetadataName, error)
	ReadAllMetadata(ctx context.Context) (*protobuf.MetadataArray, error)
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

func (e *execution) SaveMetadataToDb(ctx context.Context, metadata *protobuf.Metadata) (*protobuf.MetadataName, error) {
	result, err := e.metadata.Save(ctx, metadata.Name, metadata)
	return result, err
}

func (e *execution) ReadAllMetadata(ctx context.Context) (*protobuf.MetadataArray, error) {
	result, err := e.metadata.GetAll(ctx)
	return result, err
}
