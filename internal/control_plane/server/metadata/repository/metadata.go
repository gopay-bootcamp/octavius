package repository

import (
	"context"
	"octavius/internal/control_plane/db/etcd"
	"octavius/pkg/protobuf"

	"github.com/gogo/protobuf/proto"
)

const prefix = "metadata/"

//MetadataRepository interface for functions related to metadata repository
type MetadataRepository interface {
	Save(ctx context.Context, key string, metadata *protobuf.Metadata) (*protobuf.MetadataName, error)
	GetAll(ctx context.Context) (*protobuf.MetadataArray, error)
}

type metadataRepository struct {
	etcdClient etcd.EtcdClient
}

//NewMetadataRepository initializes metadataRepository with the given etcdClient
func NewMetadataRepository(client etcd.EtcdClient) MetadataRepository {
	return &metadataRepository{
		etcdClient: client,
	}
}

//Save marshals metadata and saves the value in etcd database with the given key
func (c *metadataRepository) Save(ctx context.Context, key string, metadata *protobuf.Metadata) (*protobuf.MetadataName, error) {
	val, err := proto.Marshal(metadata)

	if err != nil {
		errMsg := &protobuf.Error{ErrorCode: 2, ErrorMessage: "error in marshalling metadata"}
		res := &protobuf.MetadataName{Err: errMsg, Name: ""}
		return res, err
	}

	dbKey := prefix + key
	_, err = c.etcdClient.PutValue(ctx, dbKey, string(val))
	if err != nil {
		errMsg := &protobuf.Error{ErrorCode: 3, ErrorMessage: "error in saving to etcd"}
		res := &protobuf.MetadataName{Err: errMsg, Name: ""}
		return res, err
	}

	errMsg := &protobuf.Error{ErrorCode: 0, ErrorMessage: "no error"}
	res := &protobuf.MetadataName{Err: errMsg, Name: key}
	return res, nil
}

//GetAll returns array of metadata
func (c *metadataRepository) GetAll(ctx context.Context) (*protobuf.MetadataArray, error) {
	res, err := c.etcdClient.GetAllValues(ctx, prefix)
	if err != nil {
		errMsg := &protobuf.Error{ErrorCode: 3, ErrorMessage: "error in saving to etcd"}
		var arr []*protobuf.Metadata
		res := &protobuf.MetadataArray{Err: errMsg, Value: arr}
		return res, err
	}

	errMsg := &protobuf.Error{ErrorCode: 0, ErrorMessage: "no error"}
	var resArr []*protobuf.Metadata
	for _, val := range res {
		metadata := &protobuf.Metadata{}
		proto.Unmarshal([]byte(val), metadata)
		resArr = append(resArr, metadata)
	}
	resp := &protobuf.MetadataArray{Err: errMsg, Value: resArr}
	return resp, nil
}
