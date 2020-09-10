package job

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/db/etcd"
	"octavius/internal/pkg/log"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.Init("info", "", false, 1)
}

func TestSave(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	var testExecutionData = &clientCPproto.RequestForExecute{
		JobName: "testJobName",
		JobData: map[string]string{
			"env1": "envValue1",
		},
	}

	testExecutionDataAsByte, err := proto.Marshal(testExecutionData)

	mockClient.On("PutValue", "jobs/pending/12345678", string(testExecutionDataAsByte)).Return(nil)

	err = jobRepository.Save(context.Background(), uint64(12345678), testExecutionData)
	assert.Nil(t, err)
	mockClient.AssertExpectations(t)
}

func TestSaveForEtcdClientFailure(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	var testExecutionData = &clientCPproto.RequestForExecute{
		JobName: "testJobName",
		JobData: map[string]string{
			"env1": "envValue1",
		},
	}

	testExecutionDataAsByte, err := proto.Marshal(testExecutionData)

	mockClient.On("PutValue", "jobs/pending/12345678", string(testExecutionDataAsByte)).Return(errors.New("failed to put value in database"))

	err = jobRepository.Save(context.Background(), uint64(12345678), testExecutionData)
	assert.Equal(t, "failed to put value in database", err.Error())
	mockClient.AssertExpectations(t)
}

func TestCheckJobIsAvailableForJobAvailable(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	mockClient.On("GetValue", "metadata/testJobName").Return("metadata of testJobName", nil)

	jobAvailabilityStatus, err := jobRepository.CheckJobIsAvailable(context.Background(), "testJobName")
	assert.Nil(t, err)
	assert.True(t, jobAvailabilityStatus)
	mockClient.AssertExpectations(t)
}

func TestCheckJobIsAvailableForJobNotAvailable(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	mockClient.On("GetValue", "metadata/testJobName").Return("", errors.New(constant.NoValueFound))

	jobAvailabilityStatus, err := jobRepository.CheckJobIsAvailable(context.Background(), "testJobName")
	assert.Equal(t, status.Error(codes.NotFound, constant.Etcd+"job with testJobName name not found").Error(), err.Error())
	assert.False(t, jobAvailabilityStatus)
	mockClient.AssertExpectations(t)
}

func TestCheckJobIsAvailableForEtcdClientFailure(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	mockClient.On("GetValue", "metadata/testJobName").Return("", errors.New("failed to get value from database"))

	jobAvailabilityStatus, err := jobRepository.CheckJobIsAvailable(context.Background(), "testJobName")
	assert.Equal(t, status.Error(codes.Internal, "failed to get value from database").Error(), err.Error())
	assert.False(t, jobAvailabilityStatus)
	mockClient.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	mockClient.On("DeleteKey").Return(true, nil)
	err := jobRepository.Delete(context.Background(), "12345")
	assert.Nil(t, err)
	mockClient.AssertExpectations(t)

}

func TestDeleteForEtcdClientFailure(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	mockClient.On("DeleteKey").Return(false, errors.New("failed to delete key from database"))
	err := jobRepository.Delete(context.Background(), "12345")
	assert.Equal(t, status.Error(codes.Internal, "failed to delete key from database").Error(), err.Error())
	mockClient.AssertExpectations(t)

}

func TestFetchNextJob(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	var keys []string
	keys = append(keys, "jobs/pending/123")
	keys = append(keys, "jobs/pending/234")

	var values []string

	executionData1 := &clientCPproto.RequestForExecute{
		JobName: "testJobName1",
		JobData: map[string]string{
			"env1": "envValue1",
		},
	}

	executionData2 := &clientCPproto.RequestForExecute{
		JobName: "testJobName2",
		JobData: map[string]string{
			"env1": "envValue1",
		},
	}

	executionData1AsByte, err := proto.Marshal(executionData1)
	executionData2AsByte, err := proto.Marshal(executionData2)

	values = append(values, string(executionData1AsByte))
	values = append(values, string(executionData2AsByte))
	var nextExecutionData *clientCPproto.RequestForExecute
	nextExecutionData = &clientCPproto.RequestForExecute{}
	err = proto.Unmarshal([]byte(values[0]), nextExecutionData)
	mockClient.On("GetAllKeyAndValues", "jobs/pending/").Return(keys, values, nil)
	nextJobID, nextExecutionData, err := jobRepository.FetchNextJob(context.Background())
	assert.Equal(t, "123", nextJobID)
	assert.Equal(t, executionData1.JobName, nextExecutionData.JobName)
	assert.Equal(t, executionData1.JobData, nextExecutionData.JobData)
	assert.Nil(t, err)
	mockClient.AssertExpectations(t)

}

func TestFetchNextJobForEtcdClientFailure(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	var keys []string

	var values []string

	mockClient.On("GetAllKeyAndValues", "jobs/pending/").Return(keys, values, errors.New("failed to get keys and values from database"))
	nextJobID, nextExecutionData, err := jobRepository.FetchNextJob(context.Background())

	assert.Nil(t, nextExecutionData)
	assert.Equal(t, "", nextJobID)
	assert.Equal(t, err.Error(), status.Error(codes.Internal, "failed to get keys and values from database").Error())
	mockClient.AssertExpectations(t)

}

func TestFetchNextJobForJobNotAvailable(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	var keys []string

	var values []string

	mockClient.On("GetAllKeyAndValues", "jobs/pending/").Return(keys, values, nil)
	nextJobID, nextExecutionData, err := jobRepository.FetchNextJob(context.Background())

	assert.Nil(t, nextExecutionData)
	assert.Equal(t, "", nextJobID)
	assert.Equal(t, err.Error(), status.Error(codes.NotFound, constant.Controller+"no pending job in pending job list").Error())
	mockClient.AssertExpectations(t)

}

func TestValidateJobForSuccess(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	var testExecutionData = &clientCPproto.RequestForExecute{
		JobName: "testJobName",
		JobData: map[string]string{
			"env1": "envValue1",
		},
	}
	jobName := testExecutionData.JobName
	key := "metadata/" + jobName
	var testArgsArray []*clientCPproto.Arg
	var testArg = &clientCPproto.Arg{
		Name:        "env1",
		Description: "test env",
		Required:    true,
	}
	testArgsArray = append(testArgsArray, testArg)
	var testEnvVars = &clientCPproto.EnvVars{
		Args: testArgsArray,
	}
	var testMetadata = &clientCPproto.Metadata{
		Name:        "testJobName",
		Description: "This is a test image",
		ImageName:   "images/test-image",
		EnvVars:     testEnvVars,
	}
	str, err := proto.Marshal(testMetadata)
	if err != nil {
		log.Error(err, "error in test data marshalling")
	}
	mockClient.On("GetValue", key).Return(string(str), nil)
	flag, err := jobRepository.ValidateJob(context.Background(), testExecutionData)
	assert.Equal(t, flag, true)
	assert.Nil(t, err)
	mockClient.AssertExpectations(t)
}

func TestValidateJobForOptionalArgSuccess(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	var testExecutionData = &clientCPproto.RequestForExecute{
		JobName: "testJobName",
		JobData: map[string]string{
			"env1": "envValue1",
			"env2": "envValue2",
		},
	}
	jobName := testExecutionData.JobName
	key := "metadata/" + jobName
	var testArgsArray []*clientCPproto.Arg
	var testArg1 = &clientCPproto.Arg{
		Name:        "env1",
		Description: "test env1",
		Required:    true,
	}
	testArgsArray = append(testArgsArray, testArg1)
	var testArg2 = &clientCPproto.Arg{
		Name:        "env2",
		Description: "test env2",
		Required:    false,
	}
	testArgsArray = append(testArgsArray, testArg2)
	var testArg3 = &clientCPproto.Arg{
		Name:        "env3",
		Description: "test env3",
		Required:    false,
	}
	testArgsArray = append(testArgsArray, testArg3)

	var testEnvVars = &clientCPproto.EnvVars{
		Args: testArgsArray,
	}
	var testMetadata = &clientCPproto.Metadata{
		Name:        "testJobName",
		Description: "This is a test image",
		ImageName:   "images/test-image",
		EnvVars:     testEnvVars,
	}
	str, err := proto.Marshal(testMetadata)
	if err != nil {
		log.Error(err, "error in test data marshalling")
	}
	mockClient.On("GetValue", key).Return(string(str), nil)
	flag, err := jobRepository.ValidateJob(context.Background(), testExecutionData)
	assert.Equal(t, flag, true)
	assert.Nil(t, err)
	mockClient.AssertExpectations(t)
}

func TestValidateJobForArgMissingFailure(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	var testExecutionData = &clientCPproto.RequestForExecute{
		JobName: "testJobName",
		JobData: map[string]string{
			"env1": "envValue1",
		},
	}
	jobName := testExecutionData.JobName
	key := "metadata/" + jobName

	var testArgsArray []*clientCPproto.Arg
	var testArg1 = &clientCPproto.Arg{
		Name:        "env1",
		Description: "test env1",
		Required:    true,
	}
	testArgsArray = append(testArgsArray, testArg1)
	var testArg2 = &clientCPproto.Arg{
		Name:        "env2",
		Description: "test env2",
		Required:    true,
	}
	testArgsArray = append(testArgsArray, testArg2)
	var testEnvVars = &clientCPproto.EnvVars{
		Args: testArgsArray,
	}
	var testMetadata = &clientCPproto.Metadata{
		Name:        "testJobName",
		Description: "This is a test image",
		ImageName:   "images/test-image",
		EnvVars:     testEnvVars,
	}

	str, err := proto.Marshal(testMetadata)
	if err != nil {
		log.Error(err, "error in test data marshalling")
	}
	mockClient.On("GetValue", key).Return(string(str), nil)
	flag, err := jobRepository.ValidateJob(context.Background(), testExecutionData)
	assert.Equal(t, flag, false)
	assert.Nil(t, err)
	mockClient.AssertExpectations(t)
}

func TestValidateJobForExtraArgFailure(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	jobRepository := NewJobRepository(mockClient)

	var testExecutionData = &clientCPproto.RequestForExecute{
		JobName: "testJobName",
		JobData: map[string]string{
			"env1": "envValue1",
			"env2": "envValue2",
		},
	}
	jobName := testExecutionData.JobName
	key := "metadata/" + jobName

	var testArgsArray []*clientCPproto.Arg
	var testArg1 = &clientCPproto.Arg{
		Name:        "env1",
		Description: "test env1",
		Required:    true,
	}
	testArgsArray = append(testArgsArray, testArg1)
	var testEnvVars = &clientCPproto.EnvVars{
		Args: testArgsArray,
	}
	var testMetadata = &clientCPproto.Metadata{
		Name:        "testJobName",
		Description: "This is a test image",
		ImageName:   "images/test-image",
		EnvVars:     testEnvVars,
	}

	str, err := proto.Marshal(testMetadata)
	if err != nil {
		log.Error(err, "error in test data marshalling")
	}
	mockClient.On("GetValue", key).Return(string(str), nil)
	flag, err := jobRepository.ValidateJob(context.Background(), testExecutionData)
	assert.Equal(t, flag, false)
	assert.Nil(t, err)
	mockClient.AssertExpectations(t)
}
