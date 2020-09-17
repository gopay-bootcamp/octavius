package describe

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

func TestDescribeCmdHelp(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	testDescribeCmd := NewCmd(mockOctaviusDClient)
	assert.Equal(t, "Describe the existing job", testDescribeCmd.Short)
	assert.Equal(t, "This command helps to describe the job which is already created in server", testDescribeCmd.Long)
	assert.Equal(t, "octavius describe --job-name <job-name>", testDescribeCmd.Example)
}

func TestDescribeCmd(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	testDescribeCmd := NewCmd(mockOctaviusDClient)

	arg := protobuf.Arg{
		Name:        "test",
		Description: "test description",
		Required:    false,
	}
	var args []*protobuf.Arg
	args = append(args, &arg)
	var testEnvVars = &protobuf.EnvVars{
		Args: args,
	}
	describeResponse := &protobuf.Metadata{
		Name:    "Demo Job",
		EnvVars: testEnvVars,
	}

	mockOctaviusDClient.On("DescribeJob", "DemoJob").Return(describeResponse, nil).Once()
	testDescribeCmd.SetArgs([]string{"--job-name", "DemoJob"})
	err := testDescribeCmd.Execute()
	assert.Nil(t,err)
	mockOctaviusDClient.AssertExpectations(t)
}
