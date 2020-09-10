package list

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

func TestListCmdHelp(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	testListCmd := NewCmd(mockOctaviusDClient)
	assert.Equal(t, "Get job list", testListCmd.Short)
	assert.Equal(t, "Get job list will give available jobs in octavius", testListCmd.Long)
}

func TestListCmd(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	testListCmd := NewCmd(mockOctaviusDClient)
	var jobList []string
	jobList = append(jobList, "demo-image-name")
	jobList = append(jobList, "demo-image-name-1")

	response := &protobuf.JobList{
		Jobs: jobList,
	}
	mockOctaviusDClient.On("GetJobList").Return(response, nil)

	testListCmd.Execute()

	mockOctaviusDClient.AssertExpectations(t)
}
