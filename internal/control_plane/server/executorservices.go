package server

import (
	"context"
	"octavius/internal/control_plane/server/execution"
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
	ack, err := e.procExec.UpdateExecutorStatus(ctx, ping)
	return ack, err
}

func (e *executorCPServicesServer) Register(ctx context.Context, request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error) {
	ack, err := e.procExec.RegisterExecutor(ctx, request)
	return ack, err
}
