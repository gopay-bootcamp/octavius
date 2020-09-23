// Package health implements functions related to executor health
package health

import (
	"context"
	"octavius/internal/controller/server/repository/executor"
	executorRepo "octavius/internal/controller/server/repository/executor"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/protofiles"
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
	pingChan := make(chan string)
	sessionID := uint64(1234)
	newActiveExecutor := activeExecutor{
		pingChan:  pingChan,
		sessionID: sessionID,
		timer:     time.NewTimer(1 * time.Second),
	}

	testExecutorMap := &activeExecutorMap{
		execMap: new(sync.Map),
	}
	testExecutorMap.Put("exec 1", &newActiveExecutor)

	testExecRepo := new(executorRepo.ExecutorMock)

	testExecution := &healthExecution{
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

	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewHealthExec(executorRepoMock)

	ctx := context.Background()
	request := protofiles.Ping{
		ID:    "exec 1",
		State: "healthy",
	}
	executorRepoMock.On("Get", "exec 1").Return(&protofiles.ExecutorInfo{}, status.Error(codes.NotFound, constant.Etcd+constant.NoValueFound))
	pingTimeOut := 20 * time.Second
	res, err := testExec.UpdateExecutorStatus(ctx, &request, pingTimeOut)
	executorRepoMock.AssertExpectations(t)
	assert.Nil(t, res)
	assert.Equal(t, err.Error(), status.Error(codes.PermissionDenied, "executor not registered").Error())
}

func TestUpdateExecutorStatus(t *testing.T) {

	executorRepoMock := new(executor.ExecutorMock)

	testExec := NewHealthExec(executorRepoMock)

	ctx := context.Background()
	request := protofiles.Ping{
		ID:    "exec 1",
		State: "idle",
	}
	executorRepoMock.On("Get", "exec 1").Return(&protofiles.ExecutorInfo{}, nil)
	executorRepoMock.On("UpdateStatus", "exec 1", "idle").Return(nil)
	res, err := testExec.UpdateExecutorStatus(ctx, &request, 20*time.Second)
	_, ok := getActiveExecutorMap(testExec.(*healthExecution)).Get("exec 1")
	assert.Equal(t, res.Recieved, true)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	pingTimeOut := 20 * time.Second
	res, err = testExec.UpdateExecutorStatus(ctx, &request, pingTimeOut)
	assert.Equal(t, res.Recieved, true)
	assert.Nil(t, err)
}
