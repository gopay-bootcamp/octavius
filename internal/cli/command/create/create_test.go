package create

import (
	"errors"
	daemon "octavius/internal/cli/daemon/metadata"
	"octavius/internal/pkg/file"
	"octavius/internal/pkg/log"
	protobuf "octavius/internal/pkg/protofiles"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	log.Init("info", "", false, 1)
}

func TestCreateCmdHelp(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	mockFileUtil := new(file.MockFileUtil)
	testCreateCmd := NewCmd(mockOctaviusDClient, mockFileUtil)
	assert.Equal(t, "Create new octavius job metadata", testCreateCmd.Short)
	assert.Equal(t, "This command helps create new job metadata to your CP host with proper metadata.json file", testCreateCmd.Long)
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
	mockOctaviusDClient.On("Post", strings.NewReader("test-metadata-handler-string")).Return(testMetadataName, nil).Once()

	testCreateCmd.SetArgs([]string{"--job-path", "testfile/test_metadata.json"})
	err := testCreateCmd.Execute()
	assert.Nil(t, err)

	mockFileUtil.AssertExpectations(t)
	mockOctaviusDClient.AssertExpectations(t)
}

func TestCreateCmdForIoError(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	mockFileUtil := new(file.MockFileUtil)
	testCreateCmd := NewCmd(mockOctaviusDClient, mockFileUtil)
	mockFileUtil.On("GetIoReader", "testfile/test_metadata.json").Return(strings.NewReader(""), errors.New("test io error"))

	testCreateCmd.SetArgs([]string{"--job-path", "testfile/test_metadata.json"})
	err := testCreateCmd.Execute()
	assert.Nil(t, err)
	mockFileUtil.AssertExpectations(t)
	mockOctaviusDClient.AssertNotCalled(t, "Post")
}
