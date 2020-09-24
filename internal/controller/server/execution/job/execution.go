// Package job implements job related functions
package job

import (
	"context"
	"errors"
	"fmt"

	jobRepo "octavius/internal/controller/server/repository/job"
	"octavius/internal/controller/server/scheduler"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/protofiles"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// JobExecution interface for methods related to execution
type JobExecution interface {
	GetJob(ctx context.Context) (*protofiles.Job, error)
	ExecuteJob(ctx context.Context, request *protofiles.RequestToExecute) (uint64, error)
	GetJobLogs(ctx context.Context, jobk8sName string) (string, error)
	SaveJobExecutionData(ctx context.Context, executionData *protofiles.ExecutionContext) error
	PostExecutorStatus(ctx context.Context, ID string, status *protofiles.Status) error
	checkJobIsAvailable(ctx context.Context, jobName string) (bool, error)
	validateJob(ctx context.Context, executionData *protofiles.RequestToExecute) (bool, error)
}
type jobExecution struct {
	jobRepo     jobRepo.Repository
	idGenerator idgen.RandomIdGenerator
	scheduler   scheduler.Scheduler
}

// NewJobExec creates a new instance of job respository
func NewJobExec(jobRepo jobRepo.Repository, idGenerator idgen.RandomIdGenerator, scheduler scheduler.Scheduler) JobExecution {
	return &jobExecution{
		jobRepo:     jobRepo,
		idGenerator: idGenerator,
		scheduler:   scheduler,
	}
}

// checkJobIsAvailable returns true if given job is available otherwise returns false
func (e *jobExecution) checkJobIsAvailable(ctx context.Context, jobName string) (bool, error) {
	_, err := e.jobRepo.GetMetadata(ctx, jobName)
	if err != nil {
		if err.Error() == status.Error(codes.NotFound, constant.NoValueFound).Error() {
			return false, status.Error(codes.NotFound, constant.Etcd+fmt.Sprintf("job with %v name not found", jobName))
		}
		return false, err
	}
	return true, nil
}
func (e *jobExecution) validateJob(ctx context.Context, executionData *protofiles.RequestToExecute) (bool, error) {
	jobName := executionData.JobName
	jobData := executionData.JobData
	metadata, err := e.jobRepo.GetMetadata(ctx, jobName)
	if err != nil {
		return false, status.Error(codes.Internal, err.Error())
	}
	args := metadata.EnvVars.Args

	for _, arg := range args {
		if arg.Required {
			if _, ok := jobData[arg.Name]; !ok {
				return false, nil
			}
		}
	}
	for jobKey := range jobData {
		if !isPresentInArgs(jobKey, args) {
			return false, nil
		}
	}
	return true, nil
}

func isPresentInArgs(jobKey string, args []*protofiles.Arg) bool {
	for _, arg := range args {
		if arg.Name == jobKey {
			return true
		}
	}
	return false
}

// ExecuteJob function will call job repository and get jobId
func (e *jobExecution) ExecuteJob(ctx context.Context, executionData *protofiles.RequestToExecute) (uint64, error) {
	isAvailable, err := e.checkJobIsAvailable(ctx, executionData.JobName)
	if err != nil {
		return uint64(0), err
	}
	if !isAvailable {
		return uint64(0), status.Errorf(codes.Internal, "job with name %s not available", executionData.JobName)
	}
	valid, err := e.validateJob(ctx, executionData)
	if err != nil {
		return 0, err
	}
	if !valid {
		return 0, errors.New("job data not as per metadata")
	}
	jobID, err := e.idGenerator.Generate()
	if err != nil {
		return uint64(0), err
	}
	err = e.scheduler.AddToPendingList(ctx, jobID, executionData)
	if err != nil {
		return uint64(0), err
	}
	return jobID, err
}

func (e *jobExecution) GetJobLogs(ctx context.Context, jobk8sName string) (string, error) {
	return e.jobRepo.GetLogs(ctx, jobk8sName)
}

func (e *jobExecution) SaveJobExecutionData(ctx context.Context, executionData *protofiles.ExecutionContext) error {
	return e.jobRepo.SaveJobExecutionData(ctx, executionData.JobK8SName, executionData)
}

func (e *jobExecution) GetJob(ctx context.Context) (*protofiles.Job, error) {
	jobID, clientJob, err := e.scheduler.FetchJob(ctx)
	if err != nil {
		return nil, err
	}

	metadataName := clientJob.JobName
	metadata, err := e.jobRepo.GetMetadata(ctx, metadataName)
	if err != nil {
		return nil, err
	}
	imageName := metadata.ImageName

	job := &protofiles.Job{
		HasJob:    true,
		JobID:     jobID,
		ImageName: imageName,
		JobData:   clientJob.JobData,
	}

	return job, nil
}

func (e *jobExecution) PostExecutorStatus(ctx context.Context, ID string, status *protofiles.Status) error {
	return e.jobRepo.UpdateStatus(ctx, ID, status.Status)
}
