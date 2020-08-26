package etcd

import (
	"context"
	"errors"
	"fmt"
	"octavius/internal/control_plane/logger"
	"octavius/internal/pkg/constant"
	octerr "octavius/internal/pkg/errors"
	"time"

	"github.com/coreos/etcd/clientv3"
)

//EtcdClient is exported to be used in server/execution
type EtcdClient interface {
	DeleteKey(ctx context.Context, key string) (bool, error)
	GetValue(ctx context.Context, key string) (string, error)
	PutValue(ctx context.Context, key string, value string) error
	GetAllValues(ctx context.Context, prefix string) ([]string, error)
	GetValueWithRevision(ctx context.Context, key string, header int64) (string, error)
	Close()
	SetWatchOnPrefix(ctx context.Context, prefix string) clientv3.WatchChan
	GetProcRevisionByID(ctx context.Context, id string) (int64, error)
}

type etcdClient struct {
	db *clientv3.Client
}

//NewClient returns a new client of etcd database
func NewClient(dialTimeout time.Duration, etcdHost string) EtcdClient {
	db, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{etcdHost},
	})
	return &etcdClient{
		db: db,
	}
}

//DeleteKey deltes the key-value pair with the given key
func (client *etcdClient) DeleteKey(ctx context.Context, id string) (bool, error) {
	_, err := client.db.Delete(ctx, id)
	if err != nil {
		return false, octerr.New(3, err)
	}

	return true, nil
}

//PutValue puts the given key-value pair in etcd database
func (client *etcdClient) PutValue(ctx context.Context, key string, value string) error {
	_, err := client.db.Put(ctx, key, value)
	if err != nil {
		return octerr.New(3, err)
	}
	return nil

}

//GetValue gets the value of the given key
func (client *etcdClient) GetValue(ctx context.Context, id string) (string, error) {
	res, err := client.db.Get(ctx, id)
	if err != nil {
		return "", octerr.New(3, err)
	}
	gr := res.OpResponse().Get()
	if len(gr.Kvs) == 0 {
		return "", octerr.New(3, errors.New(constant.NoValueFound))
	}
	return string(gr.Kvs[0].Value), nil
}

//GetProcRevisionById returns revision of the key-value pair of the given key
func (client *etcdClient) GetProcRevisionByID(ctx context.Context, id string) (int64, error) {
	res, err := client.db.Get(ctx, id)
	if err != nil {
		return -1, octerr.New(3, err)
	}
	gr := res.OpResponse().Get()
	if len(gr.Kvs) == 0 {
		return -1, octerr.New(3, errors.New(constant.NoValueFound))
	}
	return gr.Header.Revision, nil
}

//GetAllValues return all values with keys starting with the given prefix
func (client *etcdClient) GetAllValues(ctx context.Context, prefix string) ([]string, error) {
	res, err := client.db.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		return nil, octerr.New(3, err)
	}
	var procs []string
	gr := res.OpResponse().Get()
	for _, kv := range gr.Kvs {
		str := string(kv.Value)
		procs = append(procs, str)
	}
	return procs, nil
}

//GetValueWithRevision returns value with revision
func (client *etcdClient) GetValueWithRevision(ctx context.Context, id string, header int64) (string, error) {

	res, err := client.db.Get(ctx, id, clientv3.WithRev(header))
	if err != nil {
		return "", octerr.New(3, err)
	}
	gr := res.OpResponse().Get()
	if len(gr.Kvs) == 0 {
		return "", octerr.New(3, errors.New(constant.NoValueFound))
	}
	return string(gr.Kvs[0].Value), nil
}

//SetWatchOnPrefix returns a watch channel on the given prefix
func (client *etcdClient) SetWatchOnPrefix(ctx context.Context, prefix string) clientv3.WatchChan {
	watchChan := client.db.Watch(ctx, prefix, clientv3.WithPrefix())
	fmt.Println("set WATCH on " + prefix)
	return watchChan

}

//Close closes connection to etcd database
func (client *etcdClient) Close() {
	logger.Info("Closing connections to db")
	defer client.db.Close()
}
