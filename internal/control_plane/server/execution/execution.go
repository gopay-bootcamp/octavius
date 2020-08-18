package execution

import (
	"context"
	"octavius/internal/control_plane/server/metadata/repository"
	"octavius/pkg/protobuf"
)

type Execution interface {
	SaveMetadataToDb(ctx context.Context, metadata *protobuf.Metadata) *protobuf.MetadataName
	ReadAllMetadata(ctx context.Context) *protobuf.MetadataArray
}

type execution struct {
	metadata repository.MetadataRepository
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewExec creates a new instance of metadataRepo
func NewExec(metadataRepo repository.MetadataRepository) Execution {
	return &execution{
		metadata: metadataRepo,
	}
}

func (e *execution) SaveMetadataToDb(ctx context.Context, metadata *protobuf.Metadata) *protobuf.MetadataName {
	result := e.metadata.Save(ctx, metadata.Name, metadata)
	return result
}

func (e *execution) ReadAllMetadata(ctx context.Context) *protobuf.MetadataArray {
	res := e.metadata.GetAll(ctx)
	return res
}
