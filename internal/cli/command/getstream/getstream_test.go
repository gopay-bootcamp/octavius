package getstream

import (
	"github.com/stretchr/testify/assert"
	"octavius/internal/cli/daemon"
	"octavius/internal/cli/printer"
	"octavius/pkg/protobuf"
	"testing"
)


func TestGetStreamCmdHelp(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	mockPrinter := new(printer.MockPrinter)
	testCreateCmd := NewCmd(mockOctaviusDClient,mockPrinter)
	assert.Equal(t, "Get job log data", testCreateCmd.Short)
	assert.Equal(t, "Get job log by giving arguments", testCreateCmd.Long)
}

func TestGetStreamCmd(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	mockPrinter := new(printer.MockPrinter)
	testCreateCmd := NewCmd(mockOctaviusDClient, mockPrinter)
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
	mockPrinter.On("Println", "Test log 1\n")
	mockPrinter.On("Println", "Test log 2\n")

	testCreateCmd.SetArgs([]string{"DemoJob"})
	testCreateCmd.Execute()

	mockOctaviusDClient.AssertExpectations(t)
	mockPrinter.AssertExpectations(t)
}

