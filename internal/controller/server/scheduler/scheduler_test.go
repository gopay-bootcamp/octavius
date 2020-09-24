// Package scheduler implements scheduling related functions
package scheduler

import (
	"context"
	"errors"
	"octavius/internal/controller/server/repository/job"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/protofiles"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToPendingList(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	mockRandomIdGenerator := idgen.IdGeneratorMock{}
	scheduler := NewScheduler(&mockRandomIdGenerator, jobRepoMock)

	testExecutionData := &protofiles.RequestToExecute{
		JobName: "testJobName1",
		JobData: map[string]string{
			"env1": "envValue1",
		},
	}
	testJobID := uint64(12345)

	jobRepoMock.On("Save", testJobID, testExecutionData).Return(nil)
	err := scheduler.AddToPendingList(context.Background(), testJobID, testExecutionData)

	assert.Nil(t, err)
	jobRepoMock.AssertExpectations(t)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestAddToPendingListForJobRepoFailure(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	mockRandomIdGenerator := idgen.IdGeneratorMock{}
	scheduler := NewScheduler(&mockRandomIdGenerator, jobRepoMock)

	testExecutionData := &protofiles.RequestToExecute{
		JobName: "testJobName1",
		JobData: map[string]string{
			"env1": "envValue1",
		},
	}
	testJobID := uint64(12345)

	jobRepoMock.On("Save", testJobID, testExecutionData).Return(errors.New("failed to save job in jobRepo"))
	err := scheduler.AddToPendingList(context.Background(), testJobID, testExecutionData)

	assert.Equal(t, "failed to save job in jobRepo", err.Error())
	jobRepoMock.AssertExpectations(t)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestRemoveFromPendingList(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	mockRandomIdGenerator := idgen.IdGeneratorMock{}
	scheduler := NewScheduler(&mockRandomIdGenerator, jobRepoMock)

	testJobID := "12345"

	jobRepoMock.On("Delete", testJobID).Return(nil)
	err := scheduler.RemoveFromPendingList(context.Background(), testJobID)

	assert.Nil(t, err)
	jobRepoMock.AssertExpectations(t)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestRemoveFromPendingListForJobRepoFailure(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	mockRandomIdGenerator := idgen.IdGeneratorMock{}
	scheduler := NewScheduler(&mockRandomIdGenerator, jobRepoMock)

	testJobID := "12345"

	jobRepoMock.On("Delete", testJobID).Return(errors.New("failed to delete job in jobRepo"))
	err := scheduler.RemoveFromPendingList(context.Background(), testJobID)

	assert.Equal(t, "failed to delete job in jobRepo", err.Error())
	jobRepoMock.AssertExpectations(t)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestFetchJob(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	mockRandomIdGenerator := idgen.IdGeneratorMock{}
	scheduler := NewScheduler(&mockRandomIdGenerator, jobRepoMock)
	expectedJobID := "demo-jobID"
	envArg := map[string]string{
		"name": "akshay",
	}
	testRequestToExecute := protofiles.RequestToExecute{
		JobName: "octavius-job",
		JobData: envArg,
	}

	jobRepoMock.On("FetchNextJob").Return(expectedJobID, &testRequestToExecute, nil).Once()
	jobRepoMock.On("Delete", expectedJobID).Return(nil).Once()
	actualJobID, jobData, err := scheduler.FetchJob(context.Background())
	jobRepoMock.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, jobData, &testRequestToExecute)
	assert.Equal(t, expectedJobID, actualJobID)
}

func TestFetchJobForJobRepoFailure(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	mockRandomIdGenerator := idgen.IdGeneratorMock{}
	scheduler := NewScheduler(&mockRandomIdGenerator, jobRepoMock)
	testRequestToExecute := protofiles.RequestToExecute{}

	jobRepoMock.On("FetchNextJob").Return("", &testRequestToExecute, errors.New("job repository failure")).Once()
	jobRepoMock.On("Delete", "").Return(nil).Once()
	actualJobID, jobData, err := scheduler.FetchJob(context.Background())
	jobRepoMock.AssertCalled(t, "FetchNextJob")
	jobRepoMock.AssertNotCalled(t, "Delete")
	assert.Equal(t, err.Error(), "job repository failure")
	assert.Equal(t, (*protofiles.RequestToExecute)(nil), jobData)
	assert.Equal(t, "", actualJobID)
}
