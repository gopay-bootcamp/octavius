package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/stretchr/testify/mock"
)

type ClientMock struct {
	mock.Mock
}

func (m *ClientMock) DeleteKey(ctx context.Context, key string) (bool,error) {
	args := m.Called()
	return args.Get(0).(bool),args.Error(1)
}

func (m *ClientMock) PutValue(ctx context.Context, key string, value string) (string, error) {
	args := m.Called(key,value)
	return args.Get(0).(string), args.Error(1)
}

func (m *ClientMock) GetValue(ctx context.Context, key string) (string, error) {
	args := m.Called(key)
	return args.Get(0).(string), args.Error(1)
}

func (m *ClientMock) GetAllValues(ctx context.Context,prefix string) ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *ClientMock) GetValueWithRevision(ctx context.Context, key string, header int64) (string, error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}

func (m *ClientMock) GetProcRevisionById(ctx context.Context, id string) (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}
func (m *ClientMock) SetWatchOnPrefix(ctx context.Context, prefix string) clientv3.WatchChan {
	args := m.Called()
	return args.Get(0).(clientv3.WatchChan)
}
func (m *ClientMock) Close() {

}
