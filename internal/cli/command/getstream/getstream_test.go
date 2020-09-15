package getstream

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

func TestGetStreamCmdHelp(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	testGetStreamCmd := NewCmd(mockOctaviusDClient)
	assert.Equal(t, "Get job log data", testGetStreamCmd.Short)
	assert.Equal(t, "Get job log by giving arguments", testGetStreamCmd.Long)
}

func TestGetStreamCmd(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	testGetStreamCmd := NewCmd(mockOctaviusDClient)
	var logResponse []protobuf.Log
	log1 := &protobuf.Log{
		Log: "Test log 1",
	}
	log2 := &protobuf.Log{
		Log: "Test log 2",
	}
	logResponse = append(logResponse, *log1)
	logResponse = append(logResponse, *log2)

	mockOctaviusDClient.On("GetStreamLog", "DemoJob").Return(&logResponse, nil).Once()

	testGetStreamCmd.SetArgs([]string{"--job-id", "DemoJob"})

	testGetStreamCmd.Execute()

	mockOctaviusDClient.AssertExpectations(t)
}
