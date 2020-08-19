package etcd

import (
	"context"
	"testing"
	"time"
)

var requestTimeout = 10 * time.Second

func TestNewClient(t *testing.T) {
	client := NewClient()
	defer client.Close()
	if client == nil {
		t.Fatal("client returned nil")
	}
}

func TestEtcdClient_PutValue(t *testing.T) {
	client := NewClient()
	defer client.Close()
	if client == nil {
		t.Fatal("client returned nil")
	}
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err := client.PutValue(ctx, "test_key", "test value")
	cancel()
	if err != nil {
		t.Error("Put value returned error", err)
	}
}

func TestEtcdClient_DeleteKey(t *testing.T) {
	client := NewClient()
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err := client.GetValue(ctx, "test_key")

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
	client := NewClient()
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err := client.PutValue(ctx, "test_key", "test value")
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
	client := NewClient()
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)

	_, err := client.PutValue(ctx, "test_key", "test value")
	if err != nil {
		t.Error("error in put value", err)
	}

	header1, err := client.GetProcRevisionById(ctx, "test_key")
	if err != nil {
		t.Error("error in getting revision number", err)
	}

	_, err = client.PutValue(ctx, "test_key", "new value")
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
