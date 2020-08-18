package execution

import (
	"octavius/internal/control_plane/server/metadata/repository"
	"context"
	"octavius/pkg/protobuf"
)

type Execution interface {
	SaveMetadataToDb(ctx context.Context, metadata *protobuf.Metadata) (*protobuf.MetadataID)
	ReadAllMetadata(ctx context.Context) (string, error)
}

type execution struct {
	metadata repository.MetadataRepository
	ctx    context.Context
	cancel context.CancelFunc
}

func NewExec(metadataRepo repository.MetadataRepository) Execution {
	return &execution{
		metadata: metadataRepo,
	}
}

func (e *execution) SaveMetadataToDb(ctx context.Context, metadata *protobuf.Metadata) (*protobuf.MetadataID) {

	result := e.metadata.Save(ctx,metadata.Name,metadata)
	return result
}

func (e *execution) ReadAllMetadata(ctx context.Context) (string, error) {
	procs, err := e.metadata.GetAll(ctx)
	if err != nil {
		return "", err
	}
	return procs, nil
}
