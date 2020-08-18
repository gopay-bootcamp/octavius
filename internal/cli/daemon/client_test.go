package daemon

import (
	"octavius/internal/cli/client"
	"octavius/internal/cli/config"
	"testing"
)

func TestCreateMetadata(t *testing.T) {
	mockGrpcClient := new(client.MockGrpcClient)
	mockConfigLoader := new(config.MockLoader)
	mockGrpcClient.On("CreateJob")
}
