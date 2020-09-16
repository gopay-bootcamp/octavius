package execution

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

func TestExecuteCmdHelp(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	testExecutionCmd := NewCmd(mockOctaviusDClient)
	assert.Equal(t, "Execute the existing job", testExecutionCmd.Short)
	assert.Equal(t, "This command helps to execute the job which is already created in server", testExecutionCmd.Long)
	assert.Equal(t, "octavius execute --job-name <job-name> --args arg1=value1,arg2=value2", testExecutionCmd.Example)
}

func TestExecuteCmd(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	testExecutionCmd := NewCmd(mockOctaviusDClient)
	var jobData = map[string]string{
		"Namespace": "default",
	}
	executedResponse := &protobuf.Response{
		Status: "success",
	}

	mockOctaviusDClient.On("ExecuteJob", "DemoJob", jobData).Return(executedResponse, nil).Once()

	testExecutionCmd.SetArgs([]string{"--job-name", "DemoJob", "--args", "Namespace=default"})

	testExecutionCmd.Execute()

	mockOctaviusDClient.AssertExpectations(t)
}
