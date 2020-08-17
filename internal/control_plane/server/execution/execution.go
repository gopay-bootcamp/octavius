package execution

import (
	"octavius/internal/control_plane/server/metadata/repository"
	"context"
	"octavius/internal/control_plane/db/etcd"
	"octavius/pkg/protobuf"
)

type Execution interface {
	SaveMetadataToDb(ctx context.Context, proc *protobuf.Proc) (string, error)
	ReadAllMetadata(ctx context.Context) ([]protobuf.Proc, error)
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

func (e *execution) SaveMetadataToDb(ctx context.Context, proc *protobuf.Proc) (string, error) {

	result, err := e.metadata.Save(ctx,proc.Name,proc)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (e *execution) ReadAllMetadata(ctx context.Context) ([]protobuf.Proc, error) {
	procs, err := e.client.GetAllValues(ctx)
	if err != nil {
		return nil, err
	}
	return procs, nil
}
