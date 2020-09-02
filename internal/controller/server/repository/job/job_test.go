package job

import (
	"context"
	"errors"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/db/etcd"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	var testJobContext = &clientCPproto.RequestForExecute{
		JobName: "testJobName",
		JobData: map[string]string{
			"env1": "envValue1",
			"env2": "envValue2",
		},
	}

	testJobContextAsByte, err := proto.Marshal(testJobContext)
	testJobContextAsString := string(testJobContextAsByte)

	mockClient.On("PutValue", "jobs/pending/12345678", testJobContextAsString).Return(nil)

	err = jobRepository.Save(context.Background(), uint64(12345678), testJobContext)
	assert.Nil(t, err)
	mockClient.AssertExpectations(t)
}

func TestSaveForEtcdClientFailure(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	var testJobContext = &clientCPproto.RequestForExecute{
		JobName: "testJobName",
		JobData: map[string]string{
			"env1": "envValue1",
			"env2": "envValue2",
		},
	}

	testJobContextAsByte, err := proto.Marshal(testJobContext)
	testJobContextAsString := string(testJobContextAsByte)

	mockClient.On("PutValue", "jobs/pending/12345678", testJobContextAsString).Return(errors.New("failed to put value in database"))

	err = jobRepository.Save(context.Background(), uint64(12345678), testJobContext)
	assert.Equal(t, "failed to put value in database", err.Error())
	mockClient.AssertExpectations(t)
}

func TestCheckJobMetadataIsAvailableForJobAvailable(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	mockClient.On("GetValue", "metadata/testJobName").Return("metadata of testJobName", nil)

	jobAvailabilityStatus, err := jobRepository.CheckJobMetadataIsAvailable(context.Background(), "testJobName")
	assert.Nil(t, err)
	assert.True(t, jobAvailabilityStatus)
	mockClient.AssertExpectations(t)
}

func TestCheckJobMetadataIsAvailableForJobNotAvailable(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	mockClient.On("GetValue", "metadata/testJobName").Return("", errors.New(constant.NoValueFound))

	jobAvailabilityStatus, err := jobRepository.CheckJobMetadataIsAvailable(context.Background(), "testJobName")
	assert.Equal(t, "job with given name not found", err.Error())
	assert.False(t, jobAvailabilityStatus)
	mockClient.AssertExpectations(t)
}

func TestCheckJobMetadataIsAvailableForEtcdClientFailure(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	mockClient.On("GetValue", "metadata/testJobName").Return("", errors.New("failed to get value from database"))

	jobAvailabilityStatus, err := jobRepository.CheckJobMetadataIsAvailable(context.Background(), "testJobName")
	assert.Equal(t, "failed to get value from database", err.Error())
	assert.False(t, jobAvailabilityStatus)
	mockClient.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	mockClient.On("DeleteKey").Return(true, nil)
	err := jobRepository.Delete(context.Background(), "12345")
	assert.Nil(t, err)
	mockClient.AssertExpectations(t)

}

func TestDeleteForEtcdClientFailure(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	mockClient.On("DeleteKey").Return(false, errors.New("failed to delete key from database"))
	err := jobRepository.Delete(context.Background(), "12345")
	assert.Equal(t, "failed to delete key from database", err.Error())
	mockClient.AssertExpectations(t)

}

func TestFetchNextJob(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	var keys []string
	keys = append(keys, "jobs/pending/123")
	keys = append(keys, "jobs/pending/234")

	var values []string

	jobContext1 := &clientCPproto.RequestForExecute{
		JobName: "testJobName1",
		JobData: map[string]string{
			"env1": "envValue1",
			"env2": "envValue2",
		},
	}

	jobContext2 := &clientCPproto.RequestForExecute{
		JobName: "testJobName2",
		JobData: map[string]string{
			"env1": "envValue1",
			"env2": "envValue2",
		},
	}

	jobContext1AsByte, err := proto.Marshal(jobContext1)
	jobContext1AsString := string(jobContext1AsByte)
	jobContext2AsByte, err := proto.Marshal(jobContext2)
	jobContext2AsString := string(jobContext2AsByte)

	values = append(values, jobContext1AsString)
	values = append(values, jobContext2AsString)
	var nextJobContext *clientCPproto.RequestForExecute
	nextJobContext = &clientCPproto.RequestForExecute{}
	err = proto.Unmarshal([]byte(values[0]), nextJobContext)
	mockClient.On("GetAllKeyAndValues", "jobs/pending/").Return(keys, values, nil)
	nextJobID, nextJobContext, err := jobRepository.FetchNextJob(context.Background())
	assert.Equal(t, "123", nextJobID)
	assert.Equal(t, jobContext1.JobName, nextJobContext.JobName)
	assert.Equal(t, jobContext1.JobData, nextJobContext.JobData)
	assert.Nil(t, err)
	mockClient.AssertExpectations(t)

}

func TestFetchNextJobForEtcdClientFailure(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	var keys []string

	var values []string

	mockClient.On("GetAllKeyAndValues", "jobs/pending/").Return(keys, values, errors.New("failed to get keys and values from database"))
	nextJobID, nextJobContext, err := jobRepository.FetchNextJob(context.Background())

	assert.Nil(t, nextJobContext)
	assert.Equal(t, "", nextJobID)
	assert.Equal(t, err.Error(), "failed to get keys and values from database")
	mockClient.AssertExpectations(t)

}

func TestFetchNextJobForJobNotAvailable(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	var keys []string

	var values []string

	mockClient.On("GetAllKeyAndValues", "jobs/pending/").Return(keys, values, nil)
	nextJobID, nextJobContext, err := jobRepository.FetchNextJob(context.Background())

	assert.Nil(t, nextJobContext)
	assert.Equal(t, "", nextJobID)
	assert.Equal(t, err.Error(), "no pending job in pending job list")
	mockClient.AssertExpectations(t)

}
