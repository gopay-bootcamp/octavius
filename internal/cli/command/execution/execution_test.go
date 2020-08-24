package execution

import (
	"github.com/stretchr/testify/assert"
	"octavius/internal/cli/daemon"
	"octavius/internal/cli/logger"
	protobuf "octavius/pkg/protofiles/client_CP"
	"testing"
)

func init() {
	logger.Setup()
}

func TestExecuteCmdHelp(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	testCreateCmd := NewCmd(mockOctaviusDClient)
	assert.Equal(t, "Execute the existing job", testCreateCmd.Short)
	assert.Equal(t, "This command helps to execute the job which is already created in server", testCreateCmd.Long)
	assert.Equal(t, "octavius execute <job-name> arg1=argvalue1 arg2=argvalue2", testCreateCmd.Example)
}

func TestExecuteCmd(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	testCreateCmd := NewCmd(mockOctaviusDClient)
	var jobData = map[string]string{
		"Namespace": "default",
	}
	executedResponse := &protobuf.Response{
		Status: "success",
	}

	mockOctaviusDClient.On("ExecuteJob", "DemoJob", jobData).Return(executedResponse, nil).Once()

	testCreateCmd.SetArgs([]string{"DemoJob", "Namespace=default"})
	testCreateCmd.Execute()

	mockOctaviusDClient.AssertExpectations(t)
}
