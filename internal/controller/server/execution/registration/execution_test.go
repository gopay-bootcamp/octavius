// Package registration implements executor registration related functions
package registration

import (
	"context"
	"github.com/stretchr/testify/assert"
	"octavius/internal/controller/server/repository/executor"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/protofiles"
	"testing"
)

func init() {
	log.Init("info", "", true, 1)
}

func TestRegistration(t *testing.T) {
	executorRepoMock := new(executor.ExecutorMock)
	testExec := NewRegistrationExec(executorRepoMock)

	testRegisterRequest := protofiles.RegisterRequest{
		ID: "Random-Id",
		ExecutorInfo: &protofiles.ExecutorInfo{
			Info: "Test Executor",
		},
	}

	testRegisterResponse := protofiles.RegisterResponse{
		Registered: true,
	}
	key := testRegisterRequest.ID
	value := testRegisterRequest.ExecutorInfo
	executorRepoMock.On("Save", key, value).Return(&testRegisterResponse, nil).Once()
	res, err := testExec.RegisterExecutor(context.Background(), &testRegisterRequest)
	executorRepoMock.AssertExpectations(t)
	assert.Nil(t, err)
	assert.True(t, res.Registered)
}
