package etcd

import (
	"context"
	"errors"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/log"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// Client is exported to be used in server/execution
// type name will be used as etcd.EtcdClient by other packages, and that stutters; consider calling this Client
type Client interface {
	DeleteKey(ctx context.Context, key string) (bool, error)
	GetValue(ctx context.Context, key string) (string, error)
	PutValue(ctx context.Context, key string, value string) error
	GetAllValues(ctx context.Context, prefix string) ([]string, error)
	GetAllKeyAndValues(ctx context.Context, prefix string) ([]string, []string,  error)
	GetValueWithRevision(ctx context.Context, key string, header int64) (string, error)
	Close()
	SetWatchOnPrefix(ctx context.Context, prefix string) clientv3.WatchChan
	GetProcRevisionByID(ctx context.Context, id string) (int64, error)
}

type etcdClient struct {
	db *clientv3.Client
}

//NewClient returns a new client of etcd database
func NewClient(dialTimeout time.Duration, etcdHost string) (Client, error) {
	// handle the error
	db, err := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{etcdHost},
	})
	if err != nil {
		return &etcdClient{}, err
	}
	return &etcdClient{
		db: db,
	}, nil
}

//DeleteKey deltes the key-value pair with the given key
func (client *etcdClient) DeleteKey(ctx context.Context, id string) (bool, error) {
	_, err := client.db.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

//PutValue puts the given key-value pair in etcd database
func (client *etcdClient) PutValue(ctx context.Context, key string, value string) error {
	_, err := client.db.Put(ctx, key, value)
	if err != nil {
		return err
	}
	return nil

}

//GetValue gets the value of the given key
func (client *etcdClient) GetValue(ctx context.Context, id string) (string, error) {
	res, err := client.db.Get(ctx, id)
	if err != nil {
		return "", err
	}
	gr := res.OpResponse().Get()
	if len(gr.Kvs) == 0 {
		return "", errors.New(constant.NoValueFound)
	}
	return string(gr.Kvs[0].Value), nil
}

//GetAllValues return all values with keys starting with the given prefix
func (client *etcdClient) GetAllKeyAndValues(ctx context.Context, prefix string) ([]string, []string, error) {
	res, err := client.db.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		return nil,nil , err
	}
	var keys []string
	var values []string
	gr := res.OpResponse().Get()
	for _, kv := range gr.Kvs {
		key := string(kv.Key)
		keys = append(keys, key)
		value := string(kv.Value)
		values = append(values, value)
	}
	return keys,values,nil
}

//GetProcRevisionById returns revision of the key-value pair of the given key
func (client *etcdClient) GetProcRevisionByID(ctx context.Context, id string) (int64, error) {
	res, err := client.db.Get(ctx, id)
	if err != nil {
		return constant.NullRevision, err
	}
	gr := res.OpResponse().Get()
	if len(gr.Kvs) == 0 {
		return constant.NullRevision, errors.New(constant.NoValueFound)
	}
	return gr.Header.Revision, nil
}

//GetAllValues return all values with keys starting with the given prefix
func (client *etcdClient) GetAllValues(ctx context.Context, prefix string) ([]string, error) {
	res, err := client.db.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		return nil, err
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
		return "", err
	}
	gr := res.OpResponse().Get()
	if len(gr.Kvs) == 0 {
		return "", errors.New(constant.NoValueFound)
	}
	return string(gr.Kvs[0].Value), nil
}

//SetWatchOnPrefix returns a watch channel on the given prefix
func (client *etcdClient) SetWatchOnPrefix(ctx context.Context, prefix string) clientv3.WatchChan {
	watchChan := client.db.Watch(ctx, prefix, clientv3.WithPrefix())
	return watchChan

}

//Close closes connection to etcd database
func (client *etcdClient) Close() {
	log.Info("Closing connections to db")
	client.db.Close()
}
