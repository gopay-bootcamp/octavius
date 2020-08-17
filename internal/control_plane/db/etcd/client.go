package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"octavius/internal/config"
	"octavius/pkg/protobuf"
	"time"

	"github.com/coreos/etcd/clientv3"
)

//EtcdClient is exported to be used in server/execution
type EtcdClient interface {
	DeleteKey(ctx context.Context, key string) error
	GetValue(ctx context.Context, key string) (*protobuf.Proc, error)
	PutValue(ctx context.Context, key string, value *protobuf.Proc) (string, error)
	GetAllValues(ctx context.Context) ([]protobuf.Proc, error)
	GetValueWithRevision(ctx context.Context, key string, header int64) (*protobuf.Proc, error)
	Close()
	SetWatchOnPrefix(ctx context.Context, prefix string) clientv3.WatchChan
	GetProcRevisionById(ctx context.Context, id string) (int64, error)
}

type etcdClient struct {
	db *clientv3.Client
}

var (
	dialTimeout    = 2 * time.Second
	requestTimeout = 10 * time.Second
	etcdHost       = "localhost:" + config.Config().EtcdPort
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
func (client *etcdClient) DeleteKey(ctx context.Context, id string) error {
	_, err := client.db.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (client *etcdClient) PutValue(ctx context.Context, key string, proc *protobuf.Proc) (string, error) {
	value, err := json.Marshal(proc)
	if err != nil {
		return "", err
	}
	_, err = client.db.Put(ctx, key, string(value))
	if err != nil {
		return "", err
	}
	return proc.Name, nil
}

func (client *etcdClient) GetValue(ctx context.Context, id string) (*protobuf.Proc, error) {
	res, err := client.db.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	gr := res.OpResponse().Get()
	if len(gr.Kvs) == 0 {
		return nil, errors.New("No proc found")
	}
	var proc *protobuf.Proc
	json.Unmarshal(gr.Kvs[0].Value, &proc)
	return proc, nil
}

func (client *etcdClient) GetProcRevisionById(ctx context.Context, id string) (int64, error) {
	res, err := client.db.Get(ctx, id)
	if err != nil {
		return -1, err
	}
	gr := res.OpResponse().Get()
	if len(gr.Kvs) == 0 {
		return -1, errors.New("No proc found")
	}
	return gr.Header.Revision, nil
}

func (client *etcdClient) GetAllValues(ctx context.Context) ([]protobuf.Proc, error) {
	prefix := "key_"
	res, err := client.db.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		return nil, err
	}
	gr := res.OpResponse().Get()
	var procs []protobuf.Proc
	for _, kv := range gr.Kvs {
		proc := protobuf.Proc{}
		str := string(kv.Value)
		json.Unmarshal([]byte(str), &proc)
		procs = append(procs, proc)
	}
	return procs, nil
}

func (client *etcdClient) GetValueWithRevision(ctx context.Context, id string, header int64) (*protobuf.Proc, error) {
	res, err := client.db.Get(ctx, id, clientv3.WithRev(header))
	if err != nil {
		return nil, err
	}
	gr := res.OpResponse().Get()
	if len(gr.Kvs) == 0 {
		return nil, errors.New("No proc found")
	}
	var proc *protobuf.Proc
	json.Unmarshal(gr.Kvs[0].Value, &proc)
	return proc, nil
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
