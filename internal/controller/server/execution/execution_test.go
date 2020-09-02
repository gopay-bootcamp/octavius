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
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
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
	jobRepoMock := new(job.JobMock)
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
	jobRepoMock := new(job.JobMock)
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
	jobRepoMock := new(job.JobMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testJobContext := &clientCPproto.RequestForExecute{
		JobName: "testJobName1",
		JobData: map[string]string{
			"env1": "envValue1",
			"env2": "envValue2",
		},
	}
	testJobID := uint64(12345)

	jobRepoMock.On("CheckJobMetadataIsAvailable", testJobContext.JobName).Return(true, nil)
	mockRandomIdGenerator.On("Generate").Return(testJobID, nil)
	mockScheduler.On("AddToPendingList", testJobID, testJobContext).Return(nil)

	jobId, err := testExec.ExecuteJob(context.Background(), testJobContext)
	assert.Nil(t, err)
	assert.Equal(t, testJobID, jobId)
	jobRepoMock.AssertExpectations(t)
	mockScheduler.AssertExpectations(t)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestExecuteJobForRandomIDGeneratorFailure(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testJobContext := &clientCPproto.RequestForExecute{
		JobName: "testJobName1",
		JobData: map[string]string{
			"env1": "envValue1",
			"env2": "envValue2",
		},
	}
	testJobID := uint64(12345)

	jobRepoMock.On("CheckJobMetadataIsAvailable", testJobContext.JobName).Return(true, nil)
	mockRandomIdGenerator.On("Generate").Return(testJobID, errors.New("failed to generate random ID"))
	mockScheduler.On("AddToPendingList", testJobID, testJobContext).Return(nil)

	jobId, err := testExec.ExecuteJob(context.Background(), testJobContext)
	assert.Equal(t, "failed to generate random ID", err.Error())
	assert.Equal(t, uint64(0), jobId)
	jobRepoMock.AssertExpectations(t)
	mockScheduler.AssertNotCalled(t, "AddToPendingList", testJobID, testJobContext)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestExecuteJobForJobRepoMockFailure(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testJobContext := &clientCPproto.RequestForExecute{
		JobName: "testJobName1",
		JobData: map[string]string{
			"env1": "envValue1",
			"env2": "envValue2",
		},
	}
	testJobID := uint64(12345)

	jobRepoMock.On("CheckJobMetadataIsAvailable", testJobContext.JobName).Return(false, errors.New("failed to check jobMetadata in job repo"))
	mockRandomIdGenerator.On("Generate").Return(testJobID, nil)
	mockScheduler.On("AddToPendingList", testJobID, testJobContext).Return(nil)

	jobId, err := testExec.ExecuteJob(context.Background(), testJobContext)
	assert.Equal(t, "failed to check jobMetadata in job repo", err.Error())
	assert.Equal(t, uint64(0), jobId)
	jobRepoMock.AssertExpectations(t)
	mockScheduler.AssertNotCalled(t, "AddToPendingList", testJobID, testJobContext)
	mockRandomIdGenerator.AssertNotCalled(t, "Generate")
}

func TestExecuteJobForSchedulerFailure(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testJobContext := &clientCPproto.RequestForExecute{
		JobName: "testJobName1",
		JobData: map[string]string{
			"env1": "envValue1",
			"env2": "envValue2",
		},
	}
	testJobID := uint64(12345)

	jobRepoMock.On("CheckJobMetadataIsAvailable", testJobContext.JobName).Return(true, nil)
	mockRandomIdGenerator.On("Generate").Return(testJobID, nil)
	mockScheduler.On("AddToPendingList", testJobID, testJobContext).Return(errors.New("failed to add job in pending list in scheduler"))

	jobId, err := testExec.ExecuteJob(context.Background(), testJobContext)
	assert.Equal(t, "failed to add job in pending list in scheduler", err.Error())
	assert.Equal(t, uint64(0), jobId)
	jobRepoMock.AssertExpectations(t)
	mockScheduler.AssertExpectations(t)
	mockRandomIdGenerator.AssertExpectations(t)
}
