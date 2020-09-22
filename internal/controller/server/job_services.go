package server

import (
	"context"
	"fmt"
	"octavius/internal/controller/server/execution/job"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/protofiles"
	"octavius/internal/pkg/util"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type jobServicesServer struct {
	procExec job.JobExecution
	idgen    idgen.RandomIdGenerator
}

// JobServiceServer used to create a new execution context
func NewJobServiceServer(exec job.JobExecution, idgen idgen.RandomIdGenerator) protofiles.JobServiceServer {
	return &jobServicesServer{
		procExec: exec,
		idgen:    idgen,
	}
}

func (e *jobServicesServer) Get(ctx context.Context, executorData *protofiles.ExecutorID) (*protofiles.Job, error) {
	uuid, err := e.idgen.Generate()
	if err != nil {
		log.Error(err, fmt.Sprintf("executor id: %s, error while assigning id to the request", executorData.ID))
		return nil, status.Error(codes.Internal, err.Error())
	}
	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)

	res, err := e.procExec.GetJob(ctx)
	if err != nil {
		if err.Error() == status.Error(codes.NotFound, constant.Controller+"no pending job").Error() {
			return &protofiles.Job{HasJob: false}, nil
		}
		log.Error(err, fmt.Sprintf("request id: %v, executor id: %s, error while assigning job to executor", uuid, executorData.ID))
		return nil, err
	}
	return res, err
}

func (s *jobServicesServer) Logs(ctx context.Context, request *protofiles.RequestToGetLogs) (*protofiles.Log, error) {
	uuid, err := s.idgen.Generate()
	if err != nil {
		log.Error(err, "error while assigning id to the request")
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Info(fmt.Sprintf("request id: %v, getlogs request received", uuid))
	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	jobLogs, err := s.procExec.GetJobLogs(ctx, request.JobName)
	if err != nil {
		log.Error(fmt.Errorf("request id: %v, error in fetching logs, error details: %v", uuid, err), "")
		return nil, err
	}
	logString := &protofiles.Log{Log: jobLogs}
	return logString, nil
}

// ExecuteJob will call Execute function of execution and get jobId
func (s *jobServicesServer) Execute(ctx context.Context, executionData *protofiles.RequestToExecute) (*protofiles.Response, error) {
	uuid, err := s.idgen.Generate()
	if err != nil {
		log.Error(err, "error while assigning id to the request")
		return nil, status.Error(codes.Internal, err.Error())
	}

	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	log.Info(fmt.Sprintf("request id: %v, ExecuteJob request received with executionData %+v", uuid, executionData))

	jobID, err := s.procExec.ExecuteJob(ctx, executionData)
	if err != nil {
		log.Error(err, fmt.Sprintf("request id: %v, error in job execution", uuid))
		return &protofiles.Response{Status: "failure"}, err
	}

	jobIDString := strconv.FormatUint(jobID, 10)
	return &protofiles.Response{Status: jobIDString}, err
}

func (e *jobServicesServer) PostExecutionData(ctx context.Context, executionData *protofiles.ExecutionContext) (*protofiles.Acknowledgement, error) {
	uuid, err := e.idgen.Generate()
	if err != nil {
		log.Error(err, fmt.Sprintf("executor id: %s, error while assigning id to the request", executionData.ExecutorID))
		return nil, status.Error(codes.Internal, err.Error())
	}

	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	log.Info(fmt.Sprintf("request id: %v, recieved execution data: %+v", uuid, executionData))
	err = e.procExec.SaveJobExecutionData(ctx, executionData)
	if err != nil {
		log.Error(err, fmt.Sprintf("request id: %v, executor id: %s, error while saving job execution logs", uuid, executionData.ExecutorID))
		return &protofiles.Acknowledgement{Recieved: true}, err
	}
	return &protofiles.Acknowledgement{Recieved: true}, nil
}

func (e *jobServicesServer) PostExecutorStatus(ctx context.Context, stat *protofiles.Status) (*protofiles.Acknowledgement, error) {
	uuid, err := e.idgen.Generate()
	if err != nil {
		log.Error(err, fmt.Sprintf("executor id: %s, error while assigning id to the request", stat.ID))
		return nil, status.Error(codes.Internal, err.Error())
	}
	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	err = e.procExec.PostExecutorStatus(ctx, stat.ID, stat)
	if err != nil {
		log.Error(err, fmt.Sprintf("request id: %v, executor id: %s, error while saving executor status", uuid, stat.ID))
		return &protofiles.Acknowledgement{Recieved: true}, err
	}
	return &protofiles.Acknowledgement{Recieved: true}, nil
}
