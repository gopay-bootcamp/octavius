package metadata

import (
	"context"
	"errors"
	"fmt"
	"octavius/internal/control_plane/logger"

	"octavius/internal/control_plane/db/etcd"
	"octavius/internal/pkg/constant"
	octerr "octavius/internal/pkg/errors"

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
		errMsg := octerr.New(2, err)
		return nil, errMsg
	}
	dbKey := prefix + key

	gr, err := c.etcdClient.GetValue(ctx, dbKey)
	if gr != "" {
		errMsg := octerr.New(2, errors.New(constant.KeyAlreadyPresent))
		return nil, errMsg
	}

	if err != nil {
		if err.Error() != constant.NoValueFound {
			return nil, err
		}
	}

	logger.Info(fmt.Sprintf("Request ID: %v, saving metadata to etcd", ctx.Value("uid")))
	err = c.etcdClient.PutValue(ctx, dbKey, string(val))
	if err != nil {
		return nil, err
	}

	res := &clientCPproto.MetadataName{Name: key}
	return res, nil
}

//GetAll returns array of metadata
func (c *metadataRepository) GetAll(ctx context.Context) (*clientCPproto.MetadataArray, error) {
	res, err := c.etcdClient.GetAllValues(ctx, prefix)
	if err != nil {
		errMsg := octerr.New(3, errors.New(constant.EtcdSaveError))
		var arr []*clientCPproto.Metadata
		res := &clientCPproto.MetadataArray{Values: arr}
		return res, errMsg
	}

	var resArr []*clientCPproto.Metadata
	for _, val := range res {
		metadata := &clientCPproto.Metadata{}
		proto.Unmarshal([]byte(val), metadata)
		resArr = append(resArr, metadata)
	}
	resp := &clientCPproto.MetadataArray{Values: resArr}
	return resp, nil
}
