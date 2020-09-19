package registration

import (
	"context"
	"octavius/internal/pkg/protofiles"

	executorRepo "octavius/internal/controller/server/repository/executor"
)

// Execution interface for methods related to execution
type RegistrationExecution interface {
	RegisterExecutor(ctx context.Context, request *protofiles.RegisterRequest) (*protofiles.RegisterResponse, error)
}
type registrationExecution struct {
	executorRepo executorRepo.Repository
}

// NewRegistrationExec creates a new instance of executor respository
func NewRegistrationExec(executorRepo executorRepo.Repository) RegistrationExecution {
	return &registrationExecution{
		executorRepo: executorRepo,
	}
}

//RegisterExecutor saves executor information in DB
func (e *registrationExecution) RegisterExecutor(ctx context.Context, request *protofiles.RegisterRequest) (*protofiles.RegisterResponse, error) {
	key := request.ID
	value := request.ExecutorInfo
	return e.executorRepo.Save(ctx, key, value)
}
