package scheduler

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"octavius/internal/pkg/db/etcd"
	"octavius/internal/pkg/idgen"
	"testing"
)

func TestAddToPendingList(t *testing.T) {
	mockEtcdClient := etcd.ClientMock{}
	mockRandomIdGenerator := idgen.IdGeneratorMock{}
	scheduler := NewScheduler(&mockEtcdClient, &mockRandomIdGenerator)

	mockRandomIdGenerator.On("Generate").Return(uint64(123456789), nil)
	mockEtcdClient.On("PutValue", "jobs/pending/123456789", "11").Return(nil)

	err := scheduler.AddToPendingList(uint64(11))
	assert.Nil(t, err)
	mockEtcdClient.AssertExpectations(t)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestAddToPendingListForIdGeneratorFailure(t *testing.T) {
	mockEtcdClient := etcd.ClientMock{}
	mockRandomIdGenerator := idgen.IdGeneratorMock{}
	scheduler := NewScheduler(&mockEtcdClient, &mockRandomIdGenerator)

	mockRandomIdGenerator.On("Generate").Return(uint64(0), errors.New("failed to generate randomId"))
	mockEtcdClient.On("PutValue", "jobs/pending/123456789", "11").Return(nil)

	err := scheduler.AddToPendingList(uint64(11))
	assert.Equal(t, err.Error(), "failed to generate randomId")
	mockEtcdClient.AssertNotCalled(t, "PutValue", "jobs/pending/123456789", "11")
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestAddToPendingListForEtcdClientFailure(t *testing.T) {
	mockEtcdClient := etcd.ClientMock{}
	mockRandomIdGenerator := idgen.IdGeneratorMock{}
	scheduler := NewScheduler(&mockEtcdClient, &mockRandomIdGenerator)

	mockRandomIdGenerator.On("Generate").Return(uint64(123456789), nil)
	mockEtcdClient.On("PutValue", "jobs/pending/123456789", "11").Return(errors.New("failed to put value in etcd"))

	err := scheduler.AddToPendingList(uint64(11))
	assert.Equal(t, err.Error(), "failed to put value in etcd")
	mockEtcdClient.AssertExpectations(t)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestRemoveFromPendingList(t *testing.T) {
	mockEtcdClient := etcd.ClientMock{}
	mockRandomIdGenerator := idgen.IdGeneratorMock{}
	scheduler := NewScheduler(&mockEtcdClient, &mockRandomIdGenerator)

	mockEtcdClient.On("DeleteKey").Return(true, nil)

	err := scheduler.RemoveFromPendingList("jobs/pending/123456789")
	assert.Nil(t, err)
	mockEtcdClient.AssertExpectations(t)
}

func TestRemoveFromPendingListForEtcdClientFailure(t *testing.T) {
	mockEtcdClient := etcd.ClientMock{}
	mockRandomIdGenerator := idgen.IdGeneratorMock{}
	scheduler := NewScheduler(&mockEtcdClient, &mockRandomIdGenerator)

	mockEtcdClient.On("DeleteKey").Return(false, errors.New("failed to delete key from etcd"))

	err := scheduler.RemoveFromPendingList("jobs/pending/123456789")
	assert.Equal(t, err.Error(), "failed to delete key from etcd")
	mockEtcdClient.AssertExpectations(t)
}

func TestFetchJob(t *testing.T) {
	mockEtcdClient := etcd.ClientMock{}
	mockRandomIdGenerator := idgen.IdGeneratorMock{}
	scheduler := NewScheduler(&mockEtcdClient, &mockRandomIdGenerator)

	var keys, values []string
	keys = append(keys, "key1")
	keys = append(keys, "key2")
	values = append(values, "value1")
	values = append(values, "value2")

	mockEtcdClient.On("GetAllKeyAndValues", "jobs/pending/").Return(keys, values, nil)
	mockEtcdClient.On("DeleteKey").Return(true, nil)

	value, err := scheduler.FetchJob()
	assert.Equal(t, value, "value1")
	assert.Nil(t, err)
	mockEtcdClient.AssertExpectations(t)
}

func TestFetchJobForEtcdClientFailure(t *testing.T) {
	mockEtcdClient := etcd.ClientMock{}
	mockRandomIdGenerator := idgen.IdGeneratorMock{}
	scheduler := NewScheduler(&mockEtcdClient, &mockRandomIdGenerator)

	var keys, values []string

	mockEtcdClient.On("GetAllKeyAndValues", "jobs/pending/").Return(keys, values, errors.New("failed to get all keys and values from etcd"))
	mockEtcdClient.On("DeleteKey").Return(true, nil)

	value, err := scheduler.FetchJob()
	assert.Equal(t, value, "")
	assert.Equal(t, err.Error(), "failed to get all keys and values from etcd")
	mockEtcdClient.AssertNotCalled(t, "DeleteKey")
	mockEtcdClient.AssertCalled(t, "GetAllKeyAndValues", "jobs/pending/")
}

func TestFetchJobForNoPendingJob(t *testing.T) {
	mockEtcdClient := etcd.ClientMock{}
	mockRandomIdGenerator := idgen.IdGeneratorMock{}
	scheduler := NewScheduler(&mockEtcdClient, &mockRandomIdGenerator)

	var keys, values []string

	mockEtcdClient.On("GetAllKeyAndValues", "jobs/pending/").Return(keys, values, nil)
	mockEtcdClient.On("DeleteKey").Return(true, nil)

	value, err := scheduler.FetchJob()
	assert.Equal(t, value, "")
	assert.Equal(t, err.Error(), "no pending job in pending job list")
	mockEtcdClient.AssertNotCalled(t, "DeleteKey")
	mockEtcdClient.AssertCalled(t, "GetAllKeyAndValues", "jobs/pending/")
}
