package execution

import (
	"context"
	"errors"
	"octavius/internal/controller/server/repository/executor"
	executorRepo "octavius/internal/controller/server/repository/executor"
	job "octavius/internal/controller/server/repository/job"
	"octavius/internal/controller/server/repository/metadata"
	metadataRepo "octavius/internal/controller/server/repository/metadata"
	"octavius/internal/controller/server/scheduler"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"sync"
	"testing"
	"time"

	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	log.Init("info", "", true)
}

func TestStartExecutorHealthCheck(t *testing.T) {
	healthChan := make(chan string)
	sessionID := uint64(1234)
	clock := clockwork.NewFakeClock()
	newActiveExecutor := activeExecutor{
		healthChan: healthChan,
		sessionID:  sessionID,
		timer:      clock.After(10),
	}

	testExecutorMap := &activeExecutorMap{
		execMap: new(sync.Map),
	}
	testExecutorMap.Put("exec 1", &newActiveExecutor)

	testMetadataRepo := new(metadataRepo.MetadataMock)
	testExecRepo := new(executorRepo.ExecutorMock)

	testExecution := &execution{
		metadataRepo:      testMetadataRepo,
		executorRepo:      testExecRepo,
		activeExecutorMap: testExecutorMap,
	}

	testExecRepo.On("UpdateStatus", "exec 1", "free").Return(nil)
	testExecRepo.On("UpdateStatus", "exec 1", "expired").Return(nil)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		startExecutorHealthCheck(testExecution, testExecutorMap, "exec 1")
		wg.Done()
	}()
	//Block for asserting normal condition
	clock.BlockUntil(1)

	// Advance the FakeClock forward in time
	clock.Advance(40 * time.Second)

	// Wait until the function completes
	wg.Wait()

	//assert exit condition
	_, exists := testExecutorMap.Get("exec 1")
	assert.Equal(t, false, exists)
	testExecRepo.AssertExpectations(t)

}

func TestUpdateExecutorStatusNotRegistered(t *testing.T) {
	jobRepoMock := new(job.JobExecutorMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)

	ctx := context.Background()
	request := executorCPproto.Ping{
		ID:    "exec 1",
		State: "healthy",
	}
	executorRepoMock.On("Get", "exec 1").Return(&executorCPproto.ExecutorInfo{}, errors.New(constant.NoValueFound))
	res, err := testExec.UpdateExecutorStatus(ctx, &request)
	executorRepoMock.AssertExpectations(t)
	assert.Nil(t, res)
	assert.Equal(t, err, status.Error(codes.PermissionDenied, "executor not registered"))
}

func TestUpdateExecutorStatus(t *testing.T) {
	jobRepoMock := new(job.JobExecutorMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)

	ctx := context.Background()
	request := executorCPproto.Ping{
		ID:    "exec 1",
		State: "free",
	}
	executorRepoMock.On("Get", "exec 1").Return(&executorCPproto.ExecutorInfo{}, nil)
	executorRepoMock.On("UpdateStatus", "exec 1", "free").Return(nil)
	res, err := testExec.UpdateExecutorStatus(ctx, &request)
	_, ok := getActiveExecutorMap(testExec.(*execution)).Get("exec 1")
	assert.Equal(t, res.Recieved, true)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)

	res, err = testExec.UpdateExecutorStatus(ctx, &request)
	assert.Equal(t, res.Recieved, true)
	assert.Nil(t, err)
}

func TestExecuteJob(t *testing.T) {
	jobRepoMock := new(job.JobExecutorMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testJobData := map[string]string{
		"env1": "envValue1",
		"env2": "envValue2",
	}

	mockScheduler.On("AddToPendingList", uint64(11)).Return(nil)
	mockRandomIdGenerator.On("Generate").Return(uint64(11), nil)
	jobRepoMock.On("ExecuteJob", "11", "testJob", testJobData).Return(nil)
	jobRepoMock.On("CheckJobMetadataIsAvailable", "testJob").Return(true, nil)

	jobId, err := testExec.ExecuteJob(context.Background(), "testJob", testJobData)
	assert.Nil(t, err)
	assert.Equal(t, uint64(11), jobId)
	jobRepoMock.AssertExpectations(t)
	mockScheduler.AssertExpectations(t)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestExecuteJobForJobExecutorRepoFailure(t *testing.T) {
	jobRepoMock := new(job.JobExecutorMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testJobData := map[string]string{
		"env1": "envValue1",
		"env2": "envValue2",
	}
	mockScheduler.On("AddToPendingList", uint64(11)).Return(nil)
	mockRandomIdGenerator.On("Generate").Return(uint64(11), nil)
	jobRepoMock.On("ExecuteJob", "11", "testJob", testJobData).Return(errors.New("failed to execute job"))
	jobRepoMock.On("CheckJobMetadataIsAvailable", "testJob").Return(true, nil)

	jobId, err := testExec.ExecuteJob(context.Background(), "testJob", testJobData)
	assert.Equal(t, "failed to execute job", err.Error())
	assert.Equal(t, uint64(11), jobId)
	jobRepoMock.AssertExpectations(t)
	mockScheduler.AssertExpectations(t)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestExecuteJobForSchedulerFailure(t *testing.T) {
	jobRepoMock := new(job.JobExecutorMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testJobData := map[string]string{
		"env1": "envValue1",
		"env2": "envValue2",
	}
	mockScheduler.On("AddToPendingList", uint64(11)).Return(errors.New("failed to add job in pending list"))
	mockRandomIdGenerator.On("Generate").Return(uint64(11), nil)
	jobRepoMock.On("ExecuteJob", "11", "testJob", testJobData).Return(nil)
	jobRepoMock.On("CheckJobMetadataIsAvailable", "testJob").Return(true, nil)

	jobId, err := testExec.ExecuteJob(context.Background(), "testJob", testJobData)
	assert.Equal(t, err.Error(), "failed to add job in pending list")
	assert.Equal(t, uint64(0), jobId)
	jobRepoMock.AssertNotCalled(t, "ExecuteJob", "11", "testJob", testJobData)
	mockScheduler.AssertExpectations(t)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestExecuteJobForRandomIdGeneratorFailure(t *testing.T) {
	jobRepoMock := new(job.JobExecutorMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testJobData := map[string]string{
		"env1": "envValue1",
		"env2": "envValue2",
	}
	mockScheduler.On("AddToPendingList", uint64(11)).Return(nil)
	mockRandomIdGenerator.On("Generate").Return(uint64(0), errors.New("failed to generate random id"))
	jobRepoMock.On("ExecuteJob", "11", "testJob", testJobData).Return(nil)
	jobRepoMock.On("CheckJobMetadataIsAvailable", "testJob").Return(true, nil)

	jobId, err := testExec.ExecuteJob(context.Background(), "testJob", testJobData)
	assert.Equal(t, err.Error(), "failed to generate random id")
	assert.Equal(t, uint64(0), jobId)
	jobRepoMock.AssertNotCalled(t, "ExecuteJob", "11", "testJob", testJobData)
	mockScheduler.AssertNotCalled(t, "AddToPendingList", uint64(11))
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestExecuteJobForJobNameNotAvailable(t *testing.T) {
	jobRepoMock := new(job.JobExecutorMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testJobData := map[string]string{
		"env1": "envValue1",
		"env2": "envValue2",
	}
	mockScheduler.On("AddToPendingList", uint64(11)).Return(nil)
	mockRandomIdGenerator.On("Generate").Return(uint64(11), nil)
	jobRepoMock.On("ExecuteJob", "11", "testJob", testJobData).Return(nil)
	jobRepoMock.On("CheckJobMetadataIsAvailable", "testJob").Return(false, nil)

	jobId, err := testExec.ExecuteJob(context.Background(), "testJob", testJobData)
	assert.Equal(t, err.Error(), "job with given name not available")
	assert.Equal(t, uint64(0), jobId)
	jobRepoMock.AssertNotCalled(t, "ExecuteJob", "11", "testJob", testJobData)
	mockScheduler.AssertNotCalled(t, "AddToPendingList", uint64(11))
	mockRandomIdGenerator.AssertNotCalled(t, "Generate")
}
