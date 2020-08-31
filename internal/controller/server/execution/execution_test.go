package execution

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"octavius/internal/controller/server/repository/executor"
	job "octavius/internal/controller/server/repository/job"
	"octavius/internal/controller/server/repository/metadata"
	"octavius/internal/controller/server/scheduler"
	"octavius/internal/pkg/idgen"
	"testing"
)

func init() {
	//log.Init("info", "")
}

/*func Test_execution_SaveMetadataToDb(t *testing.T) {
	metadataRepoMock := new(metadata.MetadataMock)

	metadataVal := &protobuf.Metadata{
		Author:      "littlestar642",
		ImageName:   "demo image",
		Name:        "test data",
		Description: "sample test metadata",
	}
	metadataResp := &protobuf.MetadataName{
		Name: "test data",
		Err: &protobuf.Error{
			ErrorCode:    0,
			ErrorMessage: "no error",
		},
	}
	metadataRepoMock.On("Save", "test data", metadataVal).Return(metadataResp, nil)
	type fields struct {
		metadata metadata.MetadataMock
	}
	type args struct {
		ctx      context.Context
		metadata *protobuf.Metadata
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *protobuf.MetadataName
		wantErr bool
	}{
		{
			fields: fields{
				metadata: *metadataRepoMock,
			},
			args: args{
				ctx:      context.Background(),
				metadata: metadataVal,
			},
			want: &protobuf.MetadataName{
				Name: "test data",
				Err: &protobuf.Error{
					ErrorCode:    0,
					ErrorMessage: "no error",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &execution{
				metadataRepo: tt.fields.metadata,
			}
			got, err := e.Save(tt.args.ctx, tt.args.metadata)
			if (err != nil) != tt.wantErr {
				t.Errorf("execution.SaveMetadataToDb() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("execution.SaveMetadataToDb() = %v, want %v", got, tt.want)
			}
		})
	}
}*/


func TestExecuteJob(t *testing.T) {
	jobExecutorRepoMock:=new(job.JobExecutorMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock :=	 new(executor.ExecutorMock)

	execution:=NewExec(metadataRepoMock,executorRepoMock,jobExecutorRepoMock,mockRandomIdGenerator,mockScheduler)

	testJobData:= map[string]string {
		"env1": "envValue1",
		"env2": "envValue2",
	}

	mockScheduler.On("AddToPendingList", uint64(11)).Return(nil)
	mockRandomIdGenerator.On("Generate").Return(uint64(11),nil)
	jobExecutorRepoMock.On("ExecuteJob","11","testJob",testJobData).Return(nil)
	jobExecutorRepoMock.On("CheckJobMetadataIsAvailable","testJob").Return(true,nil)

	jobId,err:= execution.ExecuteJob(context.Background(),"testJob",testJobData)
	assert.Nil(t,err)
	assert.Equal(t,uint64(11),jobId)
	jobExecutorRepoMock.AssertExpectations(t)
	mockScheduler.AssertExpectations(t)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestExecuteJobForJobExecutorRepoFailure(t *testing.T) {
	jobExecutorRepoMock:=new(job.JobExecutorMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock :=	 new(executor.ExecutorMock)

	execution:=NewExec(metadataRepoMock,executorRepoMock,jobExecutorRepoMock,mockRandomIdGenerator,mockScheduler)

	testJobData:= map[string]string {
		"env1": "envValue1",
		"env2": "envValue2",
	}
	mockScheduler.On("AddToPendingList", uint64(11)).Return(nil)
	mockRandomIdGenerator.On("Generate").Return(uint64(11),nil)
	jobExecutorRepoMock.On("ExecuteJob","11","testJob",testJobData).Return(errors.New("failed to execute job"))
	jobExecutorRepoMock.On("CheckJobMetadataIsAvailable","testJob").Return(true,nil)

	jobId,err:= execution.ExecuteJob(context.Background(),"testJob",testJobData)
	assert.Equal(t,"failed to execute job",err.Error())
	assert.Equal(t,uint64(11),jobId)
	jobExecutorRepoMock.AssertExpectations(t)
	mockScheduler.AssertExpectations(t)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestExecuteJobForSchedulerFailure(t *testing.T) {
	jobExecutorRepoMock:=new(job.JobExecutorMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock :=	 new(executor.ExecutorMock)

	execution:=NewExec(metadataRepoMock,executorRepoMock,jobExecutorRepoMock,mockRandomIdGenerator,mockScheduler)

	testJobData:= map[string]string {
		"env1": "envValue1",
		"env2": "envValue2",
	}
	mockScheduler.On("AddToPendingList", uint64(11)).Return(errors.New("failed to add job in pending list"))
	mockRandomIdGenerator.On("Generate").Return(uint64(11),nil)
	jobExecutorRepoMock.On("ExecuteJob","11","testJob",testJobData).Return(nil)
	jobExecutorRepoMock.On("CheckJobMetadataIsAvailable","testJob").Return(true,nil)

	jobId,err:= execution.ExecuteJob(context.Background(),"testJob",testJobData)
	assert.Equal(t,err.Error(),"failed to add job in pending list")
	assert.Equal(t,uint64(0),jobId)
	jobExecutorRepoMock.AssertNotCalled(t,"ExecuteJob","11","testJob",testJobData)
	mockScheduler.AssertExpectations(t)
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestExecuteJobForRandomIdGeneratorFailure(t *testing.T) {
	jobExecutorRepoMock:=new(job.JobExecutorMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock :=	 new(executor.ExecutorMock)

	execution:=NewExec(metadataRepoMock,executorRepoMock,jobExecutorRepoMock,mockRandomIdGenerator,mockScheduler)

	testJobData:= map[string]string {
		"env1": "envValue1",
		"env2": "envValue2",
	}
	mockScheduler.On("AddToPendingList", uint64(11)).Return(nil)
	mockRandomIdGenerator.On("Generate").Return(uint64(0),errors.New("failed to generate random id"))
	jobExecutorRepoMock.On("ExecuteJob","11","testJob",testJobData).Return(nil)
	jobExecutorRepoMock.On("CheckJobMetadataIsAvailable","testJob").Return(true,nil)

	jobId,err:= execution.ExecuteJob(context.Background(),"testJob",testJobData)
	assert.Equal(t,err.Error(),"failed to generate random id")
	assert.Equal(t,uint64(0),jobId)
	jobExecutorRepoMock.AssertNotCalled(t,"ExecuteJob","11","testJob",testJobData)
	mockScheduler.AssertNotCalled(t,"AddToPendingList", uint64(11))
	mockRandomIdGenerator.AssertExpectations(t)
}

func TestExecuteJobForJobNameNotAvailable(t *testing.T) {
	jobExecutorRepoMock:=new(job.JobExecutorMock)
	metadataRepoMock := new(metadata.MetadataMock)
	mockScheduler := new(scheduler.SchedulerMock)
	mockRandomIdGenerator := new(idgen.IdGeneratorMock)
	executorRepoMock :=	 new(executor.ExecutorMock)

	execution:=NewExec(metadataRepoMock,executorRepoMock,jobExecutorRepoMock,mockRandomIdGenerator,mockScheduler)

	testJobData:= map[string]string {
		"env1": "envValue1",
		"env2": "envValue2",
	}
	mockScheduler.On("AddToPendingList", uint64(11)).Return(nil)
	mockRandomIdGenerator.On("Generate").Return(uint64(11),nil)
	jobExecutorRepoMock.On("ExecuteJob","11","testJob",testJobData).Return(nil)
	jobExecutorRepoMock.On("CheckJobMetadataIsAvailable","testJob").Return(false,nil)

	jobId,err:= execution.ExecuteJob(context.Background(),"testJob",testJobData)
	assert.Equal(t,err.Error(),"job with given name not available")
	assert.Equal(t,uint64(0),jobId)
	jobExecutorRepoMock.AssertNotCalled(t,"ExecuteJob","11","testJob",testJobData)
	mockScheduler.AssertNotCalled(t,"AddToPendingList", uint64(11))
	mockRandomIdGenerator.AssertNotCalled(t,"Generate")
}

