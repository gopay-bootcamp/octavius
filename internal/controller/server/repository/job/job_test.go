package job

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"octavius/internal/pkg/db/etcd"
	"testing"
)

func TestExecuteJob(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobExecutionRepository := NewJobExecutionRepository(mockClient)

	testJobData := map[string]string{
		"env1": "envValue1",
		"env2": "envValue2",
	}

	mockClient.On("PutValue", "jobs/11/metadataKeyName", "metadata/testJob").Return(nil)
	mockClient.On("PutValue", "jobs/11/env/env1", "envValue1").Return(nil)
	mockClient.On("PutValue", "jobs/11/env/env2", "envValue2").Return(nil)

	err := jobExecutionRepository.ExecuteJob(context.Background(), "11", "testJob", testJobData)
	assert.Nil(t, err)
	mockClient.AssertExpectations(t)
}

func TestExecuteJobForEtcdClientFailure(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobExecutionRepository := NewJobExecutionRepository(mockClient)

	testJobData := map[string]string{
		"env1": "envValue1",
		"env2": "envValue2",
	}

	mockClient.On("PutValue", "jobs/11/metadataKeyName", "metadata/testJob").Return(errors.New("failed to put value in Etcd"))
	mockClient.On("PutValue", "jobs/11/env/env1", "envValue1").Return(nil)
	mockClient.On("PutValue", "jobs/11/env/env2", "envValue2").Return(nil)

	err := jobExecutionRepository.ExecuteJob(context.Background(), "11", "testJob", testJobData)
	assert.Equal(t, err.Error(), "failed to put value in Etcd")

	mockClient.AssertCalled(t, "PutValue", "jobs/11/metadataKeyName", "metadata/testJob")
	mockClient.AssertNotCalled(t, "PutValue", "jobs/11/env/env1", "envValue1")
	mockClient.AssertNotCalled(t, "PutValue", "jobs/11/env/env2", "envValue2")
}

func TestCheckJobMetadataIsAvailable(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobExecutionRepository := NewJobExecutionRepository(mockClient)

	var testKeys []string
	testKeys = append(testKeys, "metadata/testJob1")
	testKeys = append(testKeys, "metadata/testJob2")

	var testValues []string

	mockClient.On("GetAllKeyAndValues", "metadata/").Return(testKeys, testValues, nil)

	availabilityStatus, err := jobExecutionRepository.CheckJobMetadataIsAvailable(context.Background(), "testJob1")
	assert.True(t, availabilityStatus)
	assert.Nil(t, err)

	availabilityStatus, err = jobExecutionRepository.CheckJobMetadataIsAvailable(context.Background(), "testJob2")
	assert.True(t, availabilityStatus)
	assert.Nil(t, err)

	availabilityStatus, err = jobExecutionRepository.CheckJobMetadataIsAvailable(context.Background(), "testJob3")
	assert.False(t, availabilityStatus)
	assert.Nil(t, err)
}

func TestCheckJobMetadataIsAvailableForEtcdClientFailure(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobExecutionRepository := NewJobExecutionRepository(mockClient)

	var testKeys []string
	var testValues []string

	mockClient.On("GetAllKeyAndValues", "metadata/").Return(testKeys, testValues, errors.New("failed to read database"))

	availabilityStatus, err := jobExecutionRepository.CheckJobMetadataIsAvailable(context.Background(), "testJob1")
	assert.False(t, availabilityStatus)
	assert.Equal(t, "failed to read database", err.Error())
}
