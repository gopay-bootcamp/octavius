package server

import (
	"context"
	"fmt"
	"octavius/internal/controller/config"
	"octavius/internal/controller/server/execution"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"octavius/internal/pkg/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type executorCPServicesServer struct {
	procExec execution.Execution
	idgen    idgen.RandomIdGenerator
}

// NewExecutorServiceServer used to create a new execution context
func NewExecutorServiceServer(exec execution.Execution, idgen idgen.RandomIdGenerator) executorCPproto.ExecutorCPServicesServer {
	return &executorCPServicesServer{
		procExec: exec,
		idgen:    idgen,
	}
}

func (e *executorCPServicesServer) HealthCheck(ctx context.Context, ping *executorCPproto.Ping) (*executorCPproto.HealthResponse, error) {
	controllerConfig, err := config.Loader()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pingTimeOut := controllerConfig.ExecutorPingDeadline
	res, err := e.procExec.UpdateExecutorStatus(ctx, ping, pingTimeOut)
	if err != nil {
		log.Error(err, fmt.Sprintf("executor id: %s, error in running health check", ping.ID))
		return nil, err
	}
	return res, err
}

func (e *executorCPServicesServer) Register(ctx context.Context, request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error) {
	uuid, err := e.idgen.Generate()
	if err != nil {
		log.Error(err, fmt.Sprintf("executor id: %s, error while assigning id to the request", request.ID))
		return nil, status.Error(codes.Internal, err.Error())
	}

	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	log.Info(fmt.Sprintf("request id: %v, recieve register request from executor with id %s", uuid, request.ID))

	res, err := e.procExec.RegisterExecutor(ctx, request)
	if err != nil {
		log.Error(err, fmt.Sprintf("request id: %v, error in registering executor with id %s", uuid, request.ID))
		return nil, err
	}
	return res, err
}

func (e *executorCPServicesServer) FetchJob(ctx context.Context, executorData *executorCPproto.ExecutorID) (*executorCPproto.Job, error) {
	uuid, err := e.idgen.Generate()
	if err != nil {
		log.Error(err, fmt.Sprintf("executor id: %s, error while assigning id to the request", executorData.Id))
		return nil, status.Error(codes.Internal, err.Error())
	}
	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)

	res, err := e.procExec.GetJob(ctx, executorData)
	if err != nil {
		if err.Error() == status.Error(codes.NotFound, constant.Controller+"no pending job").Error() {
			return &executorCPproto.Job{HasJob: "no"}, nil
		}
		log.Error(err, fmt.Sprintf("request id: %v, executor id: %s, error while assigning job to executor", uuid, executorData.Id))
		return nil, err
	}
	return res, err
}

func (e *executorCPServicesServer) SendExecutionContext(ctx context.Context, executionData *executorCPproto.ExecutionContext) (*executorCPproto.Acknowledgement, error) {
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
		return &executorCPproto.Acknowledgement{Recieved: true}, err
	}
	return &executorCPproto.Acknowledgement{Recieved: true}, nil
}
