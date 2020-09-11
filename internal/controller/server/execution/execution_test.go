package execution

import (
	"context"
	"errors"
	"octavius/internal/controller/server/repository/executor"
	executorRepo "octavius/internal/controller/server/repository/executor"
	"octavius/internal/controller/server/repository/job"
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

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	log.Init("info", "", true, 1)
}

func TestStartExecutorHealthCheck(t *testing.T) {
	statusChan := make(chan string)
	sessionID := uint64(1234)
	newActiveExecutor := activeExecutor{
		statusChan: statusChan,
		sessionID:  sessionID,
		timer:      time.NewTimer(10*time.Second),
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

	testExecRepo.On("UpdateStatus", "exec 1", "idle").Return(nil)
	testExecRepo.On("UpdateStatus", "exec 1", "expired").Return(nil)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		startExecutorHealthCheck(testExecution, testExecutorMap, "exec 1")
		wg.Done()
	}()

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
	executorRepoMock.On("Get", "exec 1").Return(&executorCPproto.ExecutorInfo{}, status.Error(codes.NotFound, constant.Etcd+constant.NoValueFound))
	pingTimeOut := 20 * time.Second
	res, err := testExec.UpdateExecutorStatus(ctx, &request, pingTimeOut)
	executorRepoMock.AssertExpectations(t)
	assert.Nil(t, res)
	assert.Equal(t, err.Error(), status.Error(codes.PermissionDenied, "executor not registered").Error())
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
		State: "idle",
	}
	executorRepoMock.On("Get", "exec 1").Return(&executorCPproto.ExecutorInfo{}, nil)
	executorRepoMock.On("UpdateStatus", "exec 1", "idle").Return(nil)
	res, err := testExec.UpdateExecutorStatus(ctx, &request, 20*time.Second)
	_, ok := getActiveExecutorMap(testExec.(*execution)).Get("exec 1")
	assert.Equal(t, res.Recieved, true)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	pingTimeOut := 20 * time.Second
	res, err = testExec.UpdateExecutorStatus(ctx, &request, pingTimeOut)
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

	testExecutionData := &clientCPproto.RequestForExecute{
		JobName: "testJobName1",
		JobData: map[string]string{
			"env1": "envValue1",
		},
	}
	testJobID := uint64(12345)

	jobRepoMock.On("CheckJobIsAvailable", testExecutionData.JobName).Return(true, nil)
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
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testExecutionData := &clientCPproto.RequestForExecute{
		JobName: "testJobName1",
		JobData: map[string]string{
			"env1": "envValue1",
		},
	}
	testJobID := uint64(12345)

	jobRepoMock.On("CheckJobIsAvailable", testExecutionData.JobName).Return(true, nil)
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
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testExecutionData := &clientCPproto.RequestForExecute{
		JobName: "testJobName1",
		JobData: map[string]string{
			"env1": "envValue1",
		},
	}
	testJobID := uint64(12345)

	jobRepoMock.On("CheckJobIsAvailable", testExecutionData.JobName).Return(false, errors.New("failed to check jobMetadata in job repo"))
	mockRandomIdGenerator.On("Generate").Return(testJobID, nil)
	mockScheduler.On("AddToPendingList", testJobID, testExecutionData).Return(nil)

	jobId, err := testExec.ExecuteJob(context.Background(), testExecutionData)
	assert.Equal(t, "failed to check jobMetadata in job repo", err.Error())
	assert.Equal(t, uint64(0), jobId)
	jobRepoMock.AssertExpectations(t)
	mockScheduler.AssertNotCalled(t, "AddToPendingList", testJobID, testExecutionData)
	mockRandomIdGenerator.AssertNotCalled(t, "Generate")
}

func TestExecuteJobForSchedulerFailure(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)

	testExecutionData := &clientCPproto.RequestForExecute{
		JobName: "testJobName1",
		JobData: map[string]string{
			"env1": "envValue1",
		},
	}
	testJobID := uint64(12345)

	jobRepoMock.On("CheckJobIsAvailable", testExecutionData.JobName).Return(true, nil)
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

func TestGetMetadata(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)
	testClientInfo := &clientCPproto.ClientInfo{
		ClientEmail: "test@gmail.com",
		AccessToken: "random",
	}
	testRequestForDescribe := &clientCPproto.RequestForDescribe{
		JobName:    "testJobName",
		ClientInfo: testClientInfo,
	}
	var testMetadata = &clientCPproto.Metadata{
		Name:        "testJobName",
		Description: "This is a test image",
		ImageName:   "images/test-image",
	}
	metadataRepoMock.On("GetValue", testRequestForDescribe.JobName).Return(testMetadata, nil)
	resultMetadata, getMetadataErr := testExec.GetMetadata(context.Background(), testRequestForDescribe)
	assert.Equal(t, testMetadata, resultMetadata)
	assert.Nil(t, getMetadataErr)
	metadataRepoMock.AssertExpectations(t)

}

func TestGetJobList(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)

	var jobList []string
	jobList = append(jobList, "demo-image-name")
	jobList = append(jobList, "demo-image-name-1")

	testResponse := &clientCPproto.JobList{
		Jobs: jobList,
	}

	metadataRepoMock.On("GetAvailableJobList").Return(testResponse, nil)

	res, err := testExec.GetJobList(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, testResponse, res)
	mockScheduler.AssertExpectations(t)
}

func TestGetJobListForGetAvailableJobListFunctionErr(t *testing.T) {
	jobRepoMock := new(job.JobMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewExec(metadataRepoMock, executorRepoMock, jobRepoMock, mockRandomIdGenerator, mockScheduler)

	metadataRepoMock.On("GetAvailableJobList").Return(&clientCPproto.JobList{}, errors.New("error in GetAvailableJobList function"))

	_, err := testExec.GetJobList(context.Background())
	assert.Equal(t, "error in GetAvailableJobList function", err.Error())
	mockScheduler.AssertExpectations(t)
}
