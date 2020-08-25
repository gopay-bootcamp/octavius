package metadata

import (
	"context"
	"errors"
	"octavius/internal/control_plane/db/etcd"
	clientCPproto "octavius/internal/pkg/protofiles/client_CP"

	"github.com/gogo/protobuf/proto"
)

const prefix = "metadata/"

//MetadataRepository interface for functions related to metadata repository
type MetadataRepository interface {
	Save(ctx context.Context, key string, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error)
	GetAll(ctx context.Context) (*clientCPproto.MetadataArray, error)
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
func (c *metadataRepository) Save(ctx context.Context, key string, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error) {
	val, err := proto.Marshal(metadata)

	if err != nil {
		errMsg := &clientCPproto.Error{ErrorCode: 2, ErrorMessage: "error in marshalling metadata"}
		res := &clientCPproto.MetadataName{Err: errMsg, Name: ""}
		return res, err
	}
	dbKey := prefix + key

	gr, err := c.etcdClient.GetValue(ctx, dbKey)
	if gr != "" {
		errMsg := &clientCPproto.Error{ErrorCode: 3, ErrorMessage: "key already present"}
		val, _ = proto.Marshal(errMsg)
		res := &clientCPproto.MetadataName{Err: errMsg, Name: ""}
		return res, errors.New(string(val))
	}

	if err != nil {
		if err.Error() != "no value found" {
			errMsg := &clientCPproto.Error{ErrorCode: 3, ErrorMessage: "error in getting from etcd"}
			res := &clientCPproto.MetadataName{Err: errMsg, Name: ""}
			return res, err
		}
	}

	err = c.etcdClient.PutValue(ctx, dbKey, string(val))
	if err != nil {
		errMsg := &clientCPproto.Error{ErrorCode: 3, ErrorMessage: "error in saving to etcd"}
		res := &clientCPproto.MetadataName{Err: errMsg, Name: ""}
		return res, err
	}

	errMsg := &clientCPproto.Error{ErrorCode: 0, ErrorMessage: "no error"}
	res := &clientCPproto.MetadataName{Err: errMsg, Name: key}
	return res, nil
}

//GetAll returns array of metadata
func (c *metadataRepository) GetAll(ctx context.Context) (*clientCPproto.MetadataArray, error) {
	res, err := c.etcdClient.GetAllValues(ctx, prefix)
	if err != nil {
		errMsg := &clientCPproto.Error{ErrorCode: 3, ErrorMessage: "error in saving to etcd"}
		var arr []*clientCPproto.Metadata
		res := &clientCPproto.MetadataArray{Err: errMsg, Values: arr}
		return res, err
	}

	errMsg := &clientCPproto.Error{ErrorCode: 0, ErrorMessage: "no error"}
	var resArr []*clientCPproto.Metadata
	for _, val := range res {
		metadata := &clientCPproto.Metadata{}
		proto.Unmarshal([]byte(val), metadata)
		resArr = append(resArr, metadata)
	}
	resp := &clientCPproto.MetadataArray{Err: errMsg, Values: resArr}
	return resp, nil
}
