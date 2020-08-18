package repository

import (
	"github.com/gogo/protobuf/proto"
	"octavius/pkg/protobuf"
	"context"
	"octavius/internal/control_plane/db/etcd"
)

const prefix = "metadata/"

type MetadataRepository interface {
	Save(ctx context.Context, key string, metadata *protobuf.Metadata) (*protobuf.MetadataID)
	GetAll(ctx context.Context) (string, error)
}

type metadataRepository struct {
	etcdClient etcd.EtcdClient
}

func NewMetadataRepository(client etcd.EtcdClient) MetadataRepository{
	return &metadataRepository{
		etcdClient:client,
	}
}

func (c *metadataRepository) Save(ctx context.Context, key string, metadata *protobuf.Metadata) (*protobuf.MetadataID){
	val,err := proto.Marshal(metadata)
	errMsg := &protobuf.Error{ErrorCode:2,ErrorMessage:"error in marshalling metadata"}
	if err != nil {
		res := &protobuf.MetadataID{Error:errMsg,ID:""}
		return res
	}

	key = prefix + key
	pr,err := c.etcdClient.PutValue(ctx,key,string(val))
	errMsg = &protobuf.Error{ErrorCode:3,ErrorMessage:"error in saving to etcd"}
	if err != nil {
		res := &protobuf.MetadataID{Error:errMsg,ID:""}
		return res
	}

	errMsg = &protobuf.Error{ErrorCode:0,ErrorMessage:"no error"}
	res := &protobuf.MetadataID{Error:errMsg,ID:pr}
	return res
}

func (c *metadataRepository) GetAll(ctx context.Context) (string,error) {
	return "",nil
}