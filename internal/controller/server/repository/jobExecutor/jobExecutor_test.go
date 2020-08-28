package repository

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"octavius/internal/controller/server/scheduler"
	"octavius/internal/pkg/db/etcd"
	"octavius/internal/pkg/randomIdGenerator"
	"testing"
)

func TestExecuteJob(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(randomIdGenerator.IdGeneratorMock)
	jobExecutionRepository := NewJobExecutionRepository(mockClient, mockScheduler,mockRandomIdGenerator)

	 testJobData:= map[string]string {
		"env1": "envValue1",
		"env2": "envValue2",
	}

	mockClient.On("PutValue", "jobs/11/metadataKeyName", "metadata/testJob").Return(nil)
	mockClient.On("PutValue", "jobs/11/env/env1", "envValue1").Return(nil)
	mockClient.On("PutValue", "jobs/11/env/env2", "envValue2").Return(nil)
	mockScheduler.On("AddToPendingList", uint64(11)).Return(nil)
	mockRandomIdGenerator.On("Generate").Return(uint64(11),nil)

	jobId,err:= jobExecutionRepository.ExecuteJob(context.Background(),"testJob", testJobData)
	assert.Nil(t,err)
	assert.Equal(t,jobId,uint64(11))
	mockRandomIdGenerator.AssertExpectations(t)
	mockScheduler.AssertExpectations(t)
	mockClient.AssertExpectations(t)

}


func TestExecuteJobForEtcdClientFailure(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(randomIdGenerator.IdGeneratorMock)
	jobExecutionRepository := NewJobExecutionRepository(mockClient, mockScheduler,mockRandomIdGenerator)

	testJobData:= map[string]string {
		"env1": "envValue1",
		"env2": "envValue2",
	}

	mockClient.On("PutValue", "jobs/11/metadataKeyName", "metadata/testJob").Return(errors.New("failed to put value in Etcd"))
	mockClient.On("PutValue", "jobs/11/env/env1", "envValue1").Return(nil)
	mockClient.On("PutValue", "jobs/11/env/env2", "envValue2").Return(nil)
	mockScheduler.On("AddToPendingList", uint64(11)).Return(nil)
	mockRandomIdGenerator.On("Generate").Return(uint64(11),nil)

	jobId,err:= jobExecutionRepository.ExecuteJob(context.Background(),"testJob", testJobData)
	assert.Equal(t,err.Error(), "failed to put value in Etcd")
	assert.Equal(t,jobId,uint64(0))
	mockRandomIdGenerator.AssertExpectations(t)
	mockScheduler.AssertExpectations(t)
	mockClient.AssertCalled(t,"PutValue","jobs/11/metadataKeyName", "metadata/testJob")
	mockClient.AssertNotCalled(t,"PutValue", "jobs/11/env/env1", "envValue1")
	mockClient.AssertNotCalled(t,"PutValue", "jobs/11/env/env2", "envValue2")
}

func TestExecuteJobForRandomNumberGeneratorFailure(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(randomIdGenerator.IdGeneratorMock)
	jobExecutionRepository := NewJobExecutionRepository(mockClient, mockScheduler,mockRandomIdGenerator)

	testJobData:= map[string]string {
		"env1": "envValue1",
		"env2": "envValue2",
	}

	mockClient.On("PutValue", "jobs/11/metadataKeyName", "metadata/testJob").Return(nil)
	mockClient.On("PutValue", "jobs/11/env/env1", "envValue1").Return(nil)
	mockClient.On("PutValue", "jobs/11/env/env2", "envValue2").Return(nil)
	mockScheduler.On("AddToPendingList", uint64(11)).Return(nil)
	mockRandomIdGenerator.On("Generate").Return(uint64(0),errors.New("failed to generate random number"))

	jobId,err:= jobExecutionRepository.ExecuteJob(context.Background(),"testJob", testJobData)
	assert.Equal(t,err.Error(),"failed to generate random number")
	assert.Equal(t,jobId,uint64(0))
	mockRandomIdGenerator.AssertExpectations(t)
	mockScheduler.AssertNotCalled(t,"AddToPendingList", uint64(11))
	mockClient.AssertNotCalled(t,"PutValue","jobs/11/metadataKeyName", "metadata/testJob")
	mockClient.AssertNotCalled(t,"PutValue", "jobs/11/env/env1", "envValue1")
	mockClient.AssertNotCalled(t,"PutValue", "jobs/11/env/env2", "envValue2")


}


func TestExecuteJobForSchedulerFailure(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(randomIdGenerator.IdGeneratorMock)
	jobExecutionRepository := NewJobExecutionRepository(mockClient, mockScheduler,mockRandomIdGenerator)

	testJobData:= map[string]string {
		"env1": "envValue1",
		"env2": "envValue2",
	}

	mockClient.On("PutValue", "jobs/11/metadataKeyName", "metadata/testJob").Return(nil)
	mockClient.On("PutValue", "jobs/11/env/env1", "envValue1").Return(nil)
	mockClient.On("PutValue", "jobs/11/env/env2", "envValue2").Return(nil)
	mockScheduler.On("AddToPendingList", uint64(11)).Return(errors.New("failed to add job in pending list"))
	mockRandomIdGenerator.On("Generate").Return(uint64(11),nil)

	jobId,err:= jobExecutionRepository.ExecuteJob(context.Background(),"testJob", testJobData)
	assert.Equal(t,err.Error(),"failed to add job in pending list")
	assert.Equal(t,jobId,uint64(0))
	mockRandomIdGenerator.AssertExpectations(t)
	mockScheduler.AssertExpectations(t)
	mockClient.AssertNotCalled(t,"PutValue","jobs/11/metadataKeyName", "metadata/testJob")
	mockClient.AssertNotCalled(t,"PutValue", "jobs/11/env/env1", "envValue1")
	mockClient.AssertNotCalled(t,"PutValue", "jobs/11/env/env2", "envValue2")


}