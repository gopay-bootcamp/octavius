package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"octavius/internal/pkg/log"
	"reflect"
	"testing"
	"time"
)

var (
	requestTimeout = 10 * time.Second
	dialTimeout    = 2 * time.Second
	etcdPort       = "localhost:2379"
)

func init() {
	log.Init("info", "")
}

func TestNewClient(t *testing.T) {
	// TODO: this test is fail because we haven't mock the etcd yet.
	t.Skip()

	client, err := NewClient(dialTimeout, etcdPort)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	if client == nil {
		t.Fatal("client returned nil")
	}
}

func TestEtcdClient_PutValue(t *testing.T) {
	// TODO: this test is fail because we haven't mock the etcd yet.
	t.Skip()

	client, err := NewClient(dialTimeout, etcdPort)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	if client == nil {
		t.Fatal("client returned nil")
	}
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	err = client.PutValue(ctx, "test_key", "test value")
	cancel()
	if err != nil {
		t.Error("Put value returned error", err)
	}
}

func TestEtcdClient_DeleteKey(t *testing.T) {
	// TODO: this test is fail because we haven't mock the etcd yet.
	t.Skip()

	client, err := NewClient(dialTimeout, etcdPort)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = client.GetValue(ctx, "test_key")
	if err != nil {
		t.Error("error in getting value")
	}
	status, err := client.DeleteKey(ctx, "test_key")
	if err != nil {
		t.Error("error in deleting key")
	}

	if status == false {
		t.Error("key not deleted")
	}

	val, err := client.GetValue(ctx, "test_key")
	if err == nil {
		t.Error("value still being retirieved")
	}
	if val != "" {
		t.Error("key not deleted", val)
	}
	cancel()
}

func TestEtcdClient_GetValue(t *testing.T) {
	// TODO: this test is fail because we haven't mock the etcd yet.
	t.Skip()

	client, err := NewClient(dialTimeout, etcdPort)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	err = client.PutValue(ctx, "test_key", "test value")
	if err != nil {
		t.Error("error in get value")
	}
	res, err := client.GetValue(ctx, "test_key")
	cancel()
	if err != nil {
		t.Error("error in get value", err)
	}
	if res != "test value" {
		t.Errorf("expected %s, returned %s", "test value", res)
	}
}
func TestEtcdClient_GetValueWithRevision(t *testing.T) {
	// TODO: this test is fail because we haven't mock the etcd yet.
	t.Skip()

	client, err := NewClient(dialTimeout, etcdPort)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)

	err = client.PutValue(ctx, "test_key", "test value")
	if err != nil {
		t.Error("error in put value", err)
	}

	header1, err := client.GetProcRevisionByID(ctx, "test_key")
	if err != nil {
		t.Error("error in getting revision number", err)
	}

	err = client.PutValue(ctx, "test_key", "new value")
	if err != nil {
		t.Error("error in put value", err)
	}

	grv, err := client.GetValueWithRevision(ctx, "test_key", header1)
	if err != nil {
		t.Error("error in get value", err)
	}
	if grv != "test value" {
		t.Errorf("expected %s, returned %s", "test value", grv)
	}
	cancel()
}

func Test_etcdClient_GetAllKeyAndValues(t *testing.T) {
	// TODO: this test is fail because we haven't mock the etcd yet.
	type fields struct {
		db *clientv3.Client
	}
	type args struct {
		ctx    context.Context
		prefix string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		want1   []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &etcdClient{
				db: tt.fields.db,
			}
			got, got1, err := client.GetAllKeyAndValues(tt.args.ctx, tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllKeyAndValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllKeyAndValues() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetAllKeyAndValues() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}