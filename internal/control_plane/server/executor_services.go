package server

import (
	"context"
	"fmt"
	"octavius/internal/control_plane/id_generator"
	"octavius/internal/control_plane/logger"
	"octavius/internal/control_plane/server/execution"
	"octavius/internal/control_plane/util"
	executorCPproto "octavius/internal/pkg/protofiles/executor_CP"
)

type executorCPServicesServer struct {
	procExec execution.Execution
}

// NewExecutorServiceServer used to create a new execution context
func NewExecutorServiceServer(exec execution.Execution) executorCPproto.ExecutorCPServicesServer {
	return &executorCPServicesServer{
		procExec: exec,
	}
}

func (e *executorCPServicesServer) HealthCheck(ctx context.Context, ping *executorCPproto.Ping) (*executorCPproto.HealthResponse, error) {
	uuid, err := id_generator.NextID()
	if err != nil {
		logger.Error(err, "Error while assigning id to the request")
	}
	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	logger.Info(fmt.Sprintf("request id: %v, Recieve Health Check from executor with id %s", uuid, ping.ID))
	res, err := e.procExec.UpdateExecutorStatus(ctx, ping)
	if err != nil {
		logger.Error(err, "error in health check")
	}
	return res, err
}

func (e *executorCPServicesServer) Register(ctx context.Context, request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error) {
	return e.procExec.RegisterExecutor(ctx, request)
}
