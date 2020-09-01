package create

import (
	"errors"
	"octavius/internal/cli/daemon"
	"octavius/internal/pkg/file"
	"octavius/internal/pkg/log"
	protobuf "octavius/internal/pkg/protofiles/client_cp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	log.Init("info", "", true)
}

func TestCreateCmdHelp(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	mockFileUtil := new(file.MockFileUtil)
	testCreateCmd := NewCmd(mockOctaviusDClient, mockFileUtil)
	assert.Equal(t, "Create new octavius job metadata", testCreateCmd.Short)
	assert.Equal(t, "This command helps create new jobmetadata to your CP host with proper metadata.json file", testCreateCmd.Long)
	assert.Equal(t, "octavius create --job-path <filepath>/metadata.json", testCreateCmd.Example)
}

func TestCreateCmd(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	mockFileUtil := new(file.MockFileUtil)
	testCreateCmd := NewCmd(mockOctaviusDClient, mockFileUtil)
	testMetadataName := &protobuf.MetadataName{
		Name: "name",
	}
	mockFileUtil.On("GetIoReader", "testfile/test_metadata.json").Return(strings.NewReader("test-metadata-handler-string"), nil)
	mockOctaviusDClient.On("CreateMetadata", strings.NewReader("test-metadata-handler-string")).Return(testMetadataName, nil).Once()

	testCreateCmd.SetArgs([]string{"--job-path", "testfile/test_metadata.json"})
	testCreateCmd.Execute()

	mockFileUtil.AssertExpectations(t)
	mockOctaviusDClient.AssertExpectations(t)
}

func TestCreateCmdForIoError(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	mockFileUtil := new(file.MockFileUtil)
	testCreateCmd := NewCmd(mockOctaviusDClient, mockFileUtil)
	mockFileUtil.On("GetIoReader", "testfile/test_metadata.json").Return(strings.NewReader(""), errors.New("test io error"))

	testCreateCmd.SetArgs([]string{"--job-path", "testfile/test_metadata.json"})
	testCreateCmd.Execute()

	mockFileUtil.AssertExpectations(t)
	mockOctaviusDClient.AssertNotCalled(t, "CreateMetadata")
}
