// Package job implements job related functions
package job

import (
	"context"
	"errors"
	"fmt"
	"octavius/internal/controller/server/repository/job"
	"octavius/internal/controller/server/scheduler"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/protofiles"
	"testing"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stretchr/testify/assert"
)

func init() {
	log.Init("info", "", true, 1)
}

func TestExecuteJob(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)

	testExec := NewJobExec(jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testExecutionData := &protofiles.RequestToExecute{
		JobName: "testJobName1",
		JobData: map[string]string{
			"env1": "envValue1",
		},
	}
	testJobID := uint64(12345)

	testMetadata := protofiles.Metadata{
		Author:      "adbusa67",
		ImageName:   "demo image",
		Name:        "test data",
		Description: "sample test metadata",
	}
	jobRepoMock.On("GetMetadata", "testJobName1").Return(&testMetadata, nil).Once()
	jobRepoMock.On("ValidateJob", testExecutionData).Return(true, nil)
	mockRandomIdGenerator.On("Generate").Return(testJobID, nil)
	mockScheduler.On("AddToPendingList", testJobID, testExecutionData).Return(nil)

	jobID, err := testExec.ExecuteJob(context.Background(), testExecutionData)
	assert.Nil(t, err)
	assert.Equal(t, testJobID, jobID)
	jobRepoMock.AssertExpectations(t)
	mockScheduler.AssertExpectations(t)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestExecuteJobForRandomIDGeneratorFailure(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)

	testExec := NewJobExec(jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testExecutionData := &protofiles.RequestToExecute{
		JobName: "testJobName1",
		JobData: map[string]string{
			"env1": "envValue1",
		},
	}
	testJobID := uint64(12345)

	testMetadata := protofiles.Metadata{
		Author:      "adbusa67",
		ImageName:   "demo image",
		Name:        "test data",
		Description: "sample test metadata",
	}
	jobRepoMock.On("GetMetadata", "testJobName1").Return(&testMetadata, nil).Once()
	jobRepoMock.On("ValidateJob", testExecutionData).Return(true, nil)
	mockRandomIdGenerator.On("Generate").Return(testJobID, errors.New("failed to generate random ID"))
	mockScheduler.On("AddToPendingList", testJobID, testExecutionData).Return(nil)

	jobId, err := testExec.ExecuteJob(context.Background(), testExecutionData)
	assert.Equal(t, "failed to generate random ID", err.Error())
	assert.Equal(t, uint64(0), jobId)
	jobRepoMock.AssertExpectations(t)
	mockScheduler.AssertNotCalled(t, "AddToPendingList", testJobID, testExecutionData)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestExecuteJobForJobRepoMockFailure(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)

	testExec := NewJobExec(jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testExecutionData := &protofiles.RequestToExecute{
		JobName: "testJobName1",
		JobData: map[string]string{
			"env1": "envValue1",
		},
	}
	testJobID := uint64(12345)
	testMetadata := protofiles.Metadata{
		Author:      "adbusa67",
		ImageName:   "demo image",
		Name:        "test data",
		Description: "sample test metadata",
	}
	jobRepoMock.On("GetMetadata", "testJobName1").Return(&testMetadata, status.Error(codes.NotFound, constant.NoValueFound)).Once()
	mockRandomIdGenerator.On("Generate").Return(testJobID, nil)
	mockScheduler.On("AddToPendingList", testJobID, testExecutionData).Return(nil)

	jobId, err := testExec.ExecuteJob(context.Background(), testExecutionData)
	assert.Equal(t, status.Error(codes.NotFound, constant.Etcd+fmt.Sprintf("job with testJobName1 name not found")).Error(), err.Error())
	assert.Equal(t, uint64(0), jobId)
	jobRepoMock.AssertExpectations(t)
	mockScheduler.AssertNotCalled(t, "AddToPendingList", testJobID, testExecutionData)
	mockRandomIdGenerator.AssertNotCalled(t, "Generate")
}

func TestExecuteJobForSchedulerFailure(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)

	testExec := NewJobExec(jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testExecutionData := &protofiles.RequestToExecute{
		JobName: "testJobName1",
		JobData: map[string]string{
			"env1": "envValue1",
		},
	}
	testJobID := uint64(12345)

	testMetadata := protofiles.Metadata{
		Author:      "adbusa67",
		ImageName:   "demo image",
		Name:        "test data",
		Description: "sample test metadata",
	}
	jobRepoMock.On("GetMetadata", "testJobName1").Return(&testMetadata, nil).Once()
	jobRepoMock.On("ValidateJob", testExecutionData).Return(true, nil)
	mockRandomIdGenerator.On("Generate").Return(testJobID, nil)
	mockScheduler.On("AddToPendingList", testJobID, testExecutionData).Return(errors.New("failed to add job in pending list in scheduler"))

	jobId, err := testExec.ExecuteJob(context.Background(), testExecutionData)
	assert.Equal(t, "failed to add job in pending list in scheduler", err.Error())
	assert.Equal(t, uint64(0), jobId)
	jobRepoMock.AssertExpectations(t)
	mockScheduler.AssertExpectations(t)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestGetJobLogs(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	testArgs := map[string]string{"data": "test data"}
	testExecutionContext := &protofiles.ExecutionContext{
		JobK8SName: "test execution",
		JobID:      "123",
		ImageName:  "test image",
		ExecutorID: "test id",
		Status:     "CREATED",
		EnvArgs:    testArgs,
		Output:     "here are the logs",
	}

	val, err := proto.Marshal(testExecutionContext)
	assert.Nil(t, err)
	testExec := NewJobExec(jobRepoMock, mockRandomIdGenerator, mockScheduler)
	jobRepoMock.On("GetLogs", testExecutionContext.JobK8SName).Return(string(val), nil)
	res, err := testExec.GetJobLogs(context.TODO(), testExecutionContext.JobK8SName)
	assert.Nil(t, err)
	assert.Equal(t, res, string(val))
	jobRepoMock.AssertExpectations(t)
}

func TestSaveJobExecutionData(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	testArgs := map[string]string{"data": "test data"}
	testExecutionContext := &protofiles.ExecutionContext{
		JobK8SName: "demo-jobID",
		JobID:      "123",
		ImageName:  "test image",
		ExecutorID: "test id",
		Status:     "CREATED",
		EnvArgs:    testArgs,
		Output:     "here are the logs",
	}

	testExec := NewJobExec(jobRepoMock, mockRandomIdGenerator, mockScheduler)
	jobRepoMock.On("SaveJobExecutionData", "demo-jobID", testExecutionContext).Return(nil)
	err := testExec.SaveJobExecutionData(context.Background(), testExecutionContext)
	jobRepoMock.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestGetJob(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	testExec := NewJobExec(jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testMetadata := protofiles.Metadata{
		Author:      "adbusa67",
		ImageName:   "demo image",
		Name:        "test data",
		Description: "sample test metadata",
	}
	envArg := map[string]string{
		"name": "akshay",
	}
	testRequestToExecute := protofiles.RequestToExecute{
		JobName: "demo-jobName",
		JobData: envArg,
	}

	jobRepoMock.On("GetMetadata", "demo-jobName").Return(&testMetadata, nil).Once()
	mockScheduler.On("FetchJob").Return("demo-jobID", &testRequestToExecute, nil).Once()
	actualJob, err := testExec.GetJob(context.Background())
	expectedJob := &protofiles.Job{
		HasJob:    true,
		JobID:     "demo-jobID",
		ImageName: "demo image",
		JobData:   envArg,
	}
	assert.Nil(t, err)
	assert.Equal(t, expectedJob, actualJob)
	jobRepoMock.AssertExpectations(t)
	mockScheduler.AssertExpectations(t)
}
