package create

import (
	"octavius/internal/cli/daemon"
	"octavius/internal/cli/fileUtil"
	"octavius/internal/cli/printer"
	"octavius/pkg/protobuf"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCmdHelp(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	mockFileUtil := new(fileUtil.MockFileUtil)
	mockPrinter := new(printer.MockPrinter)
	testCreateCmd := NewCmd(mockOctaviusDClient, mockFileUtil, mockPrinter)
	assert.Equal(t, "Create new octavius job metadata", testCreateCmd.Short)
	assert.Equal(t, "This command helps create new jobmetadata to your CP host with proper metadata.json file", testCreateCmd.Long)
	assert.Equal(t, "octavius create --job-path <filepath>/metadata.json", testCreateCmd.Example)
}

func TestCreateCmd(t *testing.T) {
	mockOctaviusDClient := new(daemon.MockClient)
	mockFileUtil := new(fileUtil.MockFileUtil)
	mockPrinter := new(printer.MockPrinter)
	testCreateCmd := NewCmd(mockOctaviusDClient, mockFileUtil, mockPrinter)
	testResponse := &protobuf.Response{
		Status: "success",
	}
	mockFileUtil.On("GetIoReader", "testfile/test_metadata.json").Return(strings.NewReader("test-metadata-handler-string"), nil)
	mockOctaviusDClient.On("CreateMetadata", strings.NewReader("test-metadata-handler-string")).Return(testResponse, nil).Once()
	mockPrinter.On("Println", "success")

	testCreateCmd.SetArgs([]string{"--job-path", "testfile/test_metadata.json"})
	testCreateCmd.Execute()

	mockFileUtil.AssertExpectations(t)
	mockOctaviusDClient.AssertExpectations(t)
	mockPrinter.AssertExpectations(t)
}
