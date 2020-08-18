package repository

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"octavius/internal/control_plane/db/etcd"
	"octavius/pkg/protobuf"
)

const prefix = "metadata/"

type MetadataRepository interface {
	Save(ctx context.Context, key string, metadata *protobuf.Metadata) *protobuf.MetadataID
	GetAll(ctx context.Context) *protobuf.MetadataArray
}

type metadataRepository struct {
	etcdClient etcd.EtcdClient
}

func NewMetadataRepository(client etcd.EtcdClient) MetadataRepository {
	return &metadataRepository{
		etcdClient: client,
	}
}

func (c *metadataRepository) Save(ctx context.Context, key string, metadata *protobuf.Metadata) *protobuf.MetadataID {
	val, err := proto.Marshal(metadata)

	if err != nil {
		errMsg := &protobuf.Error{ErrorCode: 2, ErrorMessage: "error in marshalling metadata"}
		res := &protobuf.MetadataID{Error: errMsg, ID: ""}
		return res
	}

	key = prefix + key
	pr, err := c.etcdClient.PutValue(ctx, key, string(val))
	if err != nil {
		errMsg := &protobuf.Error{ErrorCode: 3, ErrorMessage: "error in saving to etcd"}
		res := &protobuf.MetadataID{Error: errMsg, ID: ""}
		return res
	}

	errMsg := &protobuf.Error{ErrorCode: 0, ErrorMessage: "no error"}
	res := &protobuf.MetadataID{Error: errMsg, ID: pr}
	return res
}

func (c *metadataRepository) GetAll(ctx context.Context) *protobuf.MetadataArray {
	res,err := c.etcdClient.GetAllValues(ctx,prefix)
	if err != nil {
		errMsg := &protobuf.Error{ErrorCode: 3, ErrorMessage: "error in saving to etcd"}
		var arr []*protobuf.Metadata
		res := &protobuf.MetadataArray{Error: errMsg, Value:arr}
		return res
	}

	errMsg := &protobuf.Error{ErrorCode: 0, ErrorMessage: "no error"}
	var resArr []*protobuf.Metadata
	for _,val := range res{
		metadata := &protobuf.Metadata{}
		proto.Unmarshal([]byte(val),metadata)
		resArr = append(resArr,metadata)
	}
	resp := &protobuf.MetadataArray{Error: errMsg, Value:resArr}
	return resp
}
