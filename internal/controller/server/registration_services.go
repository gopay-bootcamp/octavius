package server

import (
	"context"
	"fmt"
	"octavius/internal/controller/server/execution/registration"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/protofiles"
	"octavius/internal/pkg/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type registrationServicesServer struct {
	procExec registration.RegistrationExecution
	idgen    idgen.RandomIdGenerator
}

// NewExecutorServiceServer used to create a new execution context
func NewRegistrationServiceServer(exec registration.RegistrationExecution, idgen idgen.RandomIdGenerator) protofiles.RegistrationServiceServer {
	return &registrationServicesServer{
		procExec: exec,
		idgen:    idgen,
	}
}

func (e *registrationServicesServer) Register(ctx context.Context, request *protofiles.RegisterRequest) (*protofiles.RegisterResponse, error) {
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
