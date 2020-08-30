package execution

import (
	"context"
	"errors"
	executorRepo "octavius/internal/controller/server/repository/executor"
	metadataRepo "octavius/internal/controller/server/repository/metadata"
	"octavius/internal/pkg/constant"
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
	testMetadataRepo := new(metadataRepo.MetadataMock)
	testExecRepo := new(executorRepo.ExecutorMock)
	testExec := NewExec(testMetadataRepo, testExecRepo)
	ctx := context.Background()
	request := executorCPproto.Ping{
		ID:    "exec 1",
		State: "healthy",
	}
	testExecRepo.On("Get", "exec 1").Return(&executorCPproto.ExecutorInfo{}, errors.New(constant.NoValueFound))
	res, err := testExec.UpdateExecutorStatus(ctx, &request)
	testExecRepo.AssertExpectations(t)
	assert.Nil(t, res)
	assert.Equal(t, err, status.Error(codes.PermissionDenied, "executor not registered"))
}

func TestUpdateExecutorStatus(t *testing.T) {
	testMetadataRepo := new(metadataRepo.MetadataMock)
	testExecRepo := new(executorRepo.ExecutorMock)
	testExec := NewExec(testMetadataRepo, testExecRepo)
	ctx := context.Background()
	request := executorCPproto.Ping{
		ID:    "exec 1",
		State: "free",
	}
	testExecRepo.On("Get", "exec 1").Return(&executorCPproto.ExecutorInfo{}, nil)
	testExecRepo.On("UpdateStatus", "exec 1", "free").Return(nil)
	res, err := testExec.UpdateExecutorStatus(ctx, &request)
	_, ok := getActiveExecutorMap(testExec.(*execution)).Get("exec 1")
	assert.Equal(t, res.Recieved, true)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)

	res, err = testExec.UpdateExecutorStatus(ctx, &request)
	assert.Equal(t, res.Recieved, true)
	assert.Nil(t, err)
}
