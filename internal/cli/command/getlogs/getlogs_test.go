package getlogs

import (
	"octavius/internal/cli/daemon"
	"octavius/internal/pkg/log"
	protobuf "octavius/internal/pkg/protofiles/client_cp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	log.Init("info", "", false, 1)
}

func TestGetLogsCmdHelp(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	testGetLogsCmd := NewCmd(mockOctaviusDClient)
	assert.Equal(t, "Get job log data", testGetLogsCmd.Short)
	assert.Equal(t, "Get job log by giving arguments", testGetLogsCmd.Long)
}

func TestGetLogsCmd(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	testGetLogsCmd := NewCmd(mockOctaviusDClient)
	logResponse := protobuf.Log{
		Log: "sample log 1",
	}

	mockOctaviusDClient.On("GetLogs", "DemoJob").Return(&logResponse, nil).Once()

	testGetLogsCmd.SetArgs([]string{"--job-id", "DemoJob"})

	testGetLogsCmd.Execute()

	mockOctaviusDClient.AssertExpectations(t)
}
