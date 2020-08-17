package repository

import (
	"octavius/pkg/protobuf"
	"context"
	"octavius/internal/control_plane/db/etcd"
)

const prefix = "metadata/"

type MetadataRepository interface {
	Save(ctx context.Context, key string, metadata string) error
	GetAll(ctx context.Context, key string) (string, error)
}

type metadataRepository struct {
	etcdClient etcd.EtcdClient
}

func NewMetadataRepository(client etcd.EtcdClient) MetadataRepository{
	return &metadataRepository{
		etcdClient:client,
	}
}

func (c *metadataRepository) Save(ctx context.Context, key string, metadata *protobuf.Proc) error{
	return nil
}

func (c *metadataRepository) GetAll(ctx context.Context, key string) (string,error) {
	return "",nil
}