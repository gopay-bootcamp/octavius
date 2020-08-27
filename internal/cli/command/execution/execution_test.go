package execution

import (
	"octavius/internal/cli/daemon"
	"octavius/internal/pkg/log"
	protobuf "octavius/internal/pkg/protofiles/client_cp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	log.Init("info", "")
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
