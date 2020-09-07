package metadata

import (
	"context"
	"errors"
	"fmt"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/db/etcd"
	"octavius/internal/pkg/log"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
	"octavius/internal/pkg/util"

	"github.com/golang/protobuf/proto"
)

//Repository interface for functions related to metadata repository
type Repository interface {
	Save(ctx context.Context, key string, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error)
	GetAll(ctx context.Context) (*clientCPproto.MetadataArray, error)
	Get(ctx context.Context, name string) (*clientCPproto.Metadata, error)
}

type metadataRepository struct {
	etcdClient etcd.Client
}

//NewMetadataRepository initializes metadataRepository with the given etcdClient
func NewMetadataRepository(client etcd.Client) Repository {
	return &metadataRepository{
		etcdClient: client,
	}
}

//Save marshals metadata and saves the value in etcd database with the given key
func (c *metadataRepository) Save(ctx context.Context, key string, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error) {

	val, err := proto.Marshal(metadata)

	if err != nil {
		return nil, err
	}
	dbKey := constant.MetadataPrefix + key

	gr, err := c.etcdClient.GetValue(ctx, dbKey)
	if gr != "" {
		return nil, errors.New(constant.KeyAlreadyPresent)
	}

	if err != nil {
		if err.Error() != constant.NoValueFound {
			return nil, err
		}
	}

	log.Info(fmt.Sprintf("Request ID: %v, saving metadata to etcd with value %s", ctx.Value(util.ContextKeyUUID), metadata.String()))
	err = c.etcdClient.PutValue(ctx, dbKey, string(val))
	if err != nil {
		return nil, err
	}

	res := &clientCPproto.MetadataName{Name: key}
	return res, nil
}

//GetAll returns array of metadata
func (c *metadataRepository) GetAll(ctx context.Context) (*clientCPproto.MetadataArray, error) {
	res, err := c.etcdClient.GetAllValues(ctx, constant.MetadataPrefix)
	if err != nil {
		var arr []*clientCPproto.Metadata
		res := &clientCPproto.MetadataArray{Values: arr}
		return res, err
	}

	var resArr []*clientCPproto.Metadata
	for _, val := range res {
		metadata := &clientCPproto.Metadata{}
		err := proto.Unmarshal([]byte(val), metadata)
		if err != nil {
			return nil, err
		}
		resArr = append(resArr, metadata)
	}
	resp := &clientCPproto.MetadataArray{Values: resArr}
	return resp, nil
}

func (c *metadataRepository) Get(ctx context.Context, name string) (*clientCPproto.Metadata, error) {
	res, err := c.etcdClient.GetValue(ctx, name)
	if err != nil {
		return nil, err
	}

	metadata := &clientCPproto.Metadata{}
	err = proto.Unmarshal([]byte(res), metadata)
	return metadata, err
}
