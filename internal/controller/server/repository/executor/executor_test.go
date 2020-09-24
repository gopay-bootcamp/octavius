// Package executor implements executor repository related functions
package executor

import (
	"context"
	"octavius/internal/pkg/db/etcd"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/protofiles"

	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

func init() {
	log.Init("info", "", false, 1)
}

func Test_ExecutorRepo_Save(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	executorInfo := &protofiles.ExecutorInfo{
		Info: "values here about info",
	}

	testExecutorRepo := NewExecutorRepository(mockClient)

	val, err := proto.Marshal(executorInfo)
	if err != nil {
		t.Error("error in marshalling metadata")
	}

	mockClient.On("PutValue", "executor/register/random ID", string(val)).Return(nil)
	ctx := context.Background()

	res, err := testExecutorRepo.Save(ctx, "random ID", executorInfo)

	if err != nil {
		t.Errorf("error in saving executor info, %v", err)
	}
	if res.Registered != true {
		t.Errorf("error in registration : expected %v got %v", true, false)
	}

	mockClient.AssertExpectations(t)
}

func Test_ExecutorRepo_Save_PutValueError(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	executorInfo := &protofiles.ExecutorInfo{
		Info: "values here about info",
	}

	testExecutorRepo := NewExecutorRepository(mockClient)

	val, err := proto.Marshal(executorInfo)
	if err != nil {
		t.Error("error in marshalling metadata")
	}

	mockClient.On("PutValue", "executor/register/random ID", string(val)).Return(errors.New("some error"))
	ctx := context.Background()

	_, err = testExecutorRepo.Save(ctx, "random ID", executorInfo)

	if err != nil {
		t.Skip()
	}
}

func Test_ExecutorRepo_Get(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	executorInfo := &protofiles.ExecutorInfo{
		Info: "values here about info",
	}

	val, err := proto.Marshal(executorInfo)
	if err != nil {
		t.Error("error in marshalling metadata")
	}

	testExecutorRepo := NewExecutorRepository(mockClient)

	mockClient.On("GetValue", "executor/register/random ID").Return(string(val), nil)
	ctx := context.Background()

	_, err = testExecutorRepo.Get(ctx, "random ID")

	if err != nil {
		t.Errorf("error in getting executor info, %v", err)
	}

	mockClient.AssertExpectations(t)
}

func Test_ExecutorRepo_Get_GetValueError(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	testExecutorRepo := NewExecutorRepository(mockClient)

	mockClient.On("GetValue", "executor/register/random ID").Return("", errors.New("some error"))
	ctx := context.Background()

	_, err := testExecutorRepo.Get(ctx, "random ID")

	if err != nil {
		t.Skip()
	}
}

func Test_ExecutorRepo_UpdateStatus(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	testExecutorRepo := NewExecutorRepository(mockClient)

	status := "expired"
	mockClient.On("PutValue", "executor/status/random ID", status).Return(nil)
	ctx := context.Background()

	err := testExecutorRepo.UpdateStatus(ctx, "random ID", status)

	if err != nil {
		t.Errorf("error in saving executor info, %v", err)
	}

	mockClient.AssertExpectations(t)
}
