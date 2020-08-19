package etcd

import (
	"context"
	"errors"
	"fmt"
	"octavius/internal/config"
	"time"

	"github.com/coreos/etcd/clientv3"
)

//EtcdClient is exported to be used in server/execution
type EtcdClient interface {
	DeleteKey(ctx context.Context, key string) (bool, error)
	GetValue(ctx context.Context, key string) (string, error)
	PutValue(ctx context.Context, key string, value string) (string, error)
	GetAllValues(ctx context.Context, prefix string) ([]string, error)
	GetValueWithRevision(ctx context.Context, key string, header int64) (string, error)
	Close()
	SetWatchOnPrefix(ctx context.Context, prefix string) clientv3.WatchChan
	GetProcRevisionById(ctx context.Context, id string) (int64, error)
}

type etcdClient struct {
	db *clientv3.Client
}

var (
	dialTimeout = 2 * time.Second
	etcdHost    = "localhost:" + config.Config().EtcdPort
)

// function to create new client of etcd database
func NewClient() EtcdClient {

	db, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{etcdHost},
	})
	return &etcdClient{
		db: db,
	}
}

// function to delete the key provided
func (client *etcdClient) DeleteKey(ctx context.Context, id string) (bool, error) {
	_, err := client.db.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (client *etcdClient) PutValue(ctx context.Context, key string, value string) (string, error) {
	_, err := client.db.Put(ctx, key, value)
	if err != nil {
		return "", err
	}
	return key, nil
}

func (client *etcdClient) GetValue(ctx context.Context, id string) (string, error) {
	res, err := client.db.Get(ctx, id)
	if err != nil {
		return "", err
	}
	gr := res.OpResponse().Get()
	if len(gr.Kvs) == 0 {
		return "", errors.New("no value found")
	}
	return string(gr.Kvs[0].Value), nil
}

func (client *etcdClient) GetProcRevisionById(ctx context.Context, id string) (int64, error) {
	res, err := client.db.Get(ctx, id)
	if err != nil {
		return -1, err
	}
	gr := res.OpResponse().Get()
	if len(gr.Kvs) == 0 {
		return -1, errors.New("no value found")
	}
	return gr.Header.Revision, nil
}

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

func (client *etcdClient) GetValueWithRevision(ctx context.Context, id string, header int64) (string, error) {
	res, err := client.db.Get(ctx, id, clientv3.WithRev(header))
	if err != nil {
		return "", err
	}
	gr := res.OpResponse().Get()
	if len(gr.Kvs) == 0 {
		return "", errors.New("no value found")
	}
	return string(gr.Kvs[0].Value), nil
}

func (client *etcdClient) SetWatchOnPrefix(ctx context.Context, prefix string) clientv3.WatchChan {
	watchChan := client.db.Watch(ctx, prefix, clientv3.WithPrefix())
	fmt.Println("set WATCH on " + prefix)
	return watchChan

}

func (client *etcdClient) Close() {
	fmt.Println("Closing connections to db")
	defer client.db.Close()
}
