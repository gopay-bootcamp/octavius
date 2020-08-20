package etcd

import (
	"context"

	"github.com/coreos/etcd/clientv3"
	"github.com/stretchr/testify/mock"
)

// ClientMock struct
type ClientMock struct {
	mock.Mock
}

// DeleteKey mock
func (m *ClientMock) DeleteKey(ctx context.Context, key string) (bool, error) {
	args := m.Called()
	return args.Get(0).(bool), args.Error(1)
}

// PutValue mock
func (m *ClientMock) PutValue(ctx context.Context, key string, value string) (error) {
	args := m.Called(key, value)
	return args.Error(0)
}

// GetValue mock
func (m *ClientMock) GetValue(ctx context.Context, key string) (string, error) {
	args := m.Called(key)
	return args.Get(0).(string), args.Error(1)
}

// GetAllValues mock
func (m *ClientMock) GetAllValues(ctx context.Context, prefix string) ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

// GetValueWithRevision mock
func (m *ClientMock) GetValueWithRevision(ctx context.Context, key string, header int64) (string, error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}

// GetProcRevisionById mock
func (m *ClientMock) GetProcRevisionByID(ctx context.Context, id string) (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

// SetWatchOnPrefix mock
func (m *ClientMock) SetWatchOnPrefix(ctx context.Context, prefix string) clientv3.WatchChan {
	args := m.Called()
	return args.Get(0).(clientv3.WatchChan)
}

// Close mock
func (m *ClientMock) Close() {

}
