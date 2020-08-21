package execution

import (
	"github.com/stretchr/testify/assert"
	"octavius/internal/cli/daemon"
	"octavius/internal/cli/fileUtil"
	"octavius/pkg/protobuf"
	"strings"

	"octavius/internal/cli/printer"
	"testing"
)

func TestExecuteCmdHelp(t *testing.T) {
mockOctaviusDClient := new(daemon.MockClient)
mockPrinter := new(printer.MockPrinter)
testCreateCmd := NewCmd(mockOctaviusDClient,mockPrinter)
assert.Equal(t, "Execute the existing job", testCreateCmd.Short)
assert.Equal(t, "This command helps to execute the job which is already created in server", testCreateCmd.Long)
assert.Equal(t, "octavius execute <job-name> arg1=argvalue1 arg2=argvalue2", testCreateCmd.Example)
}

func TestExecuteCmd(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	mockPrinter := new(printer.MockPrinter)
	testCreateCmd := NewCmd(mockOctaviusDClient, mockPrinter)
	status := &protobuf.response{
		status: "success",
	}

	mockOctaviusDClient.On("ExecuteJob", strings.NewReader("DemoJob"),jobData).Return(, nil).Once()
	mockPrinter.On("Println", "name")

	testCreateCmd.SetArgs([]string{"DemoJob","arg1=argvalue1"})
	testCreateCmd.Execute()

	mockOctaviusDClient.AssertExpectations(t)
	mockPrinter.AssertExpectations(t)
}
