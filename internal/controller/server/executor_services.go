package server

import (
	"context"
	"fmt"
	"octavius/internal/controller/server/execution"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"
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
	uuid, err := idgen.NextID()
	if err != nil {
		log.Error(err, "Error while assigning id to the request")
	}
	ctx = context.WithValue(ctx, uidKey, uuid)
	log.Info(fmt.Sprintf("request id: %v, Recieve Health Check from executor with id %s", uuid, ping.ID))
	res, err := e.procExec.UpdateExecutorStatus(ctx, ping)
	if err != nil {
		log.Error(err, "error in health check")
	}
	return res, err
}

func (e *executorCPServicesServer) Register(ctx context.Context, request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error) {
	uuid, err := idgen.NextID()
	if err != nil {
		log.Error(err, "Error while assigning id to the request")
	}
	ctx = context.WithValue(ctx, uidKey, uuid)
	log.Info(fmt.Sprintf("request id: %v, recieve register request from executor with id %s", uuid, request.ID))
	res, err := e.procExec.RegisterExecutor(ctx, request)
	if err != nil {
		log.Error(err, fmt.Sprintf("request id: %v, error in registering executor with id %s", uuid, request.ID))
	}
	return res, err
}
