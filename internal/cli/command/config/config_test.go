package config

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"octavius/internal/cli/config"
	"octavius/internal/cli/fileUtil"
	"octavius/internal/cli/printer"
	"testing"
)

func TestConfigCmdHelp(t *testing.T) {
	mockFileUtil := new(fileUtil.MockFileUtil)
	mockPrinter := new(printer.MockPrinter)
	testConfigCmd := NewCmd(mockFileUtil, mockPrinter)

	assert.Equal(t, "Configure octavius client", testConfigCmd.Short)
	assert.Equal(t, "This command helps configure client with control plane host, email id and access token", testConfigCmd.Long)
	assert.Equal(t, "octavius config [flags]", testConfigCmd.Example)
}

func TestConfigCmdForConfigFileNotExist(t *testing.T) {
	mockFileUtil := new(fileUtil.MockFileUtil)
	mockPrinter := new(printer.MockPrinter)
	testConfigCmd := NewCmd(mockFileUtil, mockPrinter)

	mockFileUtil.On("IsFileExist", "job_data_example/config/octavius_client.yaml").Return(false).Once()
	mockFileUtil.On("CreateDirIfNotExist", "./job_data_example/config").Return(nil).Once()
	mockFileUtil.On("CreateFile", "job_data_example/config/octavius_client.yaml").Return(nil).Once()
	var configFileContent string
	configFileContent += fmt.Sprintf("%s: %s\n", config.OctaviusCPHost, "localhost:5050")
	configFileContent += fmt.Sprintf("%s: %s\n", config.EmailId, "jaimin.rathod@go-jek.com")
	configFileContent += fmt.Sprintf("%s: %s\n", config.AccessToken, "AllowMe")
	configFileContent += fmt.Sprintf("%s: %v\n", config.ConnectionTimeoutSecs, 10)
	mockFileUtil.On("WriteFile", "job_data_example/config/octavius_client.yaml", configFileContent).Return(nil).Once()
	mockPrinter.On("Println",fmt.Sprintln("Octavius client configured successfully")).Once()

	testConfigCmd.SetArgs([]string{"--cp-host", "localhost:5050", "--email-id", "jaimin.rathod@go-jek.com", "--time-out", "10", "--token", "AllowMe"})
	testConfigCmd.Execute()

	mockFileUtil.AssertExpectations(t)
}

func TestConfigCmdForConfigFileNotExistForDirectoryCreationFailure(t *testing.T) {
	mockFileUtil := new(fileUtil.MockFileUtil)
	mockPrinter := new(printer.MockPrinter)
	testConfigCmd := NewCmd(mockFileUtil, mockPrinter)

	mockFileUtil.On("IsFileExist", "job_data_example/config/octavius_client.yaml").Return(false).Once()
	mockFileUtil.On("CreateDirIfNotExist", "./job_data_example/config").Return(errors.New("failed to create directory")).Once()
	mockPrinter.On("Println",fmt.Sprintf("Error in creating config file directory, %v\n","failed to create directory" )).Once()

	testConfigCmd.SetArgs([]string{"--cp-host", "localhost:5050", "--email-id", "jaimin.rathod@go-jek.com", "--time-out", "10", "--token", "AllowMe"})
	testConfigCmd.Execute()

	mockFileUtil.AssertExpectations(t)
}

func TestConfigCmdForConfigFileNotExistForConfigFileCreationFailure(t *testing.T) {
	mockFileUtil := new(fileUtil.MockFileUtil)
	mockPrinter := new(printer.MockPrinter)
	testConfigCmd := NewCmd(mockFileUtil, mockPrinter)

	mockFileUtil.On("IsFileExist", "job_data_example/config/octavius_client.yaml").Return(false).Once()
	mockFileUtil.On("CreateDirIfNotExist", "./job_data_example/config").Return(nil).Once()
	mockFileUtil.On("CreateFile", "job_data_example/config/octavius_client.yaml").Return(errors.New("failed to create file")).Once()
	mockPrinter.On("Println",fmt.Sprintf("Error in creating config file, %v\n", "failed to create file")).Once()

	testConfigCmd.SetArgs([]string{"--cp-host", "localhost:5050", "--email-id", "jaimin.rathod@go-jek.com", "--time-out", "10", "--token", "AllowMe"})
	testConfigCmd.Execute()

	mockFileUtil.AssertExpectations(t)
}

func TestConfigCmdForConfigFileNotExistForConfigFileWritingFailure(t *testing.T) {
	mockFileUtil := new(fileUtil.MockFileUtil)
	mockPrinter := new(printer.MockPrinter)
	testConfigCmd := NewCmd(mockFileUtil, mockPrinter)

	mockFileUtil.On("IsFileExist", "job_data_example/config/octavius_client.yaml").Return(false).Once()
	mockFileUtil.On("CreateDirIfNotExist", "./job_data_example/config").Return(nil).Once()
	mockFileUtil.On("CreateFile", "job_data_example/config/octavius_client.yaml").Return(nil).Once()
	var configFileContent string
	configFileContent += fmt.Sprintf("%s: %s\n", config.OctaviusCPHost, "localhost:5050")
	configFileContent += fmt.Sprintf("%s: %s\n", config.EmailId, "jaimin.rathod@go-jek.com")
	configFileContent += fmt.Sprintf("%s: %s\n", config.AccessToken, "AllowMe")
	configFileContent += fmt.Sprintf("%s: %v\n", config.ConnectionTimeoutSecs, 10)
	mockFileUtil.On("WriteFile", "job_data_example/config/octavius_client.yaml", configFileContent).Return(errors.New("failed to write into file")).Once()
	mockPrinter.On("Println",fmt.Sprintf("Error writing content %v \n to config file %s: %s\n", configFileContent, "job_data_example/config/octavius_client.yaml", "failed to write into file")).Once()

	testConfigCmd.SetArgs([]string{"--cp-host", "localhost:5050", "--email-id", "jaimin.rathod@go-jek.com", "--time-out", "10", "--token", "AllowMe"})
	testConfigCmd.Execute()

	mockFileUtil.AssertExpectations(t)
}

func TestConfigCmdForConfigFileExist(t *testing.T) {
	mockFileUtil := new(fileUtil.MockFileUtil)
	mockPrinter := new(printer.MockPrinter)
	testConfigCmd := NewCmd(mockFileUtil, mockPrinter)

	mockFileUtil.On("IsFileExist", "job_data_example/config/octavius_client.yaml").Return(true).Once()
	mockFileUtil.On("ReadFile", "job_data_example/config/octavius_client.yaml").Return("old content", nil).Once()
	mockFileUtil.On("GetUserInput").Return("Y\n", nil).Once()
	var configFileContent string
	configFileContent += fmt.Sprintf("%s: %s\n", config.OctaviusCPHost, "localhost:5050")
	configFileContent += fmt.Sprintf("%s: %s\n", config.EmailId, "jaimin.rathod@go-jek.com")
	configFileContent += fmt.Sprintf("%s: %s\n", config.AccessToken, "AllowMe")
	configFileContent += fmt.Sprintf("%s: %v\n", config.ConnectionTimeoutSecs, 10)
	mockFileUtil.On("WriteFile", "job_data_example/config/octavius_client.yaml", configFileContent).Return(nil).Once()
	mockPrinter.On("Println",fmt.Sprintln("[Warning] This will overwrite current config:")).Once()
	mockPrinter.On("Println",fmt.Sprintln("old content")).Once()
	mockPrinter.On("Println",fmt.Sprintln("\nDo you want to continue (Y/n)?\t")).Once()
	mockPrinter.On("Println",fmt.Sprintln("Octavius client configured successfully")).Once()

	testConfigCmd.SetArgs([]string{"--cp-host", "localhost:5050", "--email-id", "jaimin.rathod@go-jek.com", "--time-out", "10", "--token", "AllowMe"})
	testConfigCmd.Execute()

	mockFileUtil.AssertExpectations(t)
}

func TestConfigCmdForConfigFileExistForOldFileReadingFailure(t *testing.T) {
	mockFileUtil := new(fileUtil.MockFileUtil)
	mockPrinter := new(printer.MockPrinter)
	testConfigCmd := NewCmd(mockFileUtil, mockPrinter)

	mockFileUtil.On("IsFileExist", "job_data_example/config/octavius_client.yaml").Return(true).Once()
	mockFileUtil.On("ReadFile", "job_data_example/config/octavius_client.yaml").Return("", errors.New("failed to read file")).Once()
	mockPrinter.On("Println",fmt.Sprintln("[Warning] This will overwrite current config:")).Once()
	mockPrinter.On("Println",fmt.Sprintf("Error reading config file: %v\n","failed to read file")).Once()

	testConfigCmd.SetArgs([]string{"--cp-host", "localhost:5050", "--email-id", "jaimin.rathod@go-jek.com", "--time-out", "10", "--token", "AllowMe"})
	testConfigCmd.Execute()

	mockFileUtil.AssertExpectations(t)
}

func TestConfigCmdForConfigFileExistForUserPermissionReadingFailure(t *testing.T) {
	mockFileUtil := new(fileUtil.MockFileUtil)
	mockPrinter := new(printer.MockPrinter)
	testConfigCmd := NewCmd(mockFileUtil, mockPrinter)

	mockFileUtil.On("IsFileExist", "job_data_example/config/octavius_client.yaml").Return(true).Once()
	mockFileUtil.On("ReadFile", "job_data_example/config/octavius_client.yaml").Return("old content", nil).Once()
	mockFileUtil.On("GetUserInput").Return("", errors.New("failed to get user input")).Once()
	mockPrinter.On("Println",fmt.Sprintln("[Warning] This will overwrite current config:")).Once()
	mockPrinter.On("Println",fmt.Sprintln("old content")).Once()
	mockPrinter.On("Println",fmt.Sprintln("\nDo you want to continue (Y/n)?\t")).Once()
	mockPrinter.On("Println",fmt.Sprintln("Error getting user permission for overwriting config")).Once()

	testConfigCmd.SetArgs([]string{"--cp-host", "localhost:5050", "--email-id", "jaimin.rathod@go-jek.com", "--time-out", "10", "--token", "AllowMe"})
	testConfigCmd.Execute()

	mockFileUtil.AssertExpectations(t)
}

func TestConfigCmdForConfigFileExistForNegativeUserInput(t *testing.T) {
	mockFileUtil := new(fileUtil.MockFileUtil)
	mockPrinter := new(printer.MockPrinter)
	testConfigCmd := NewCmd(mockFileUtil, mockPrinter)

	mockFileUtil.On("IsFileExist", "job_data_example/config/octavius_client.yaml").Return(true).Once()
	mockFileUtil.On("ReadFile", "job_data_example/config/octavius_client.yaml").Return("old content", nil).Once()
	mockFileUtil.On("GetUserInput").Return("n\n", nil).Once()
	mockPrinter.On("Println",fmt.Sprintln("[Warning] This will overwrite current config:")).Once()
	mockPrinter.On("Println",fmt.Sprintln("old content")).Once()
	mockPrinter.On("Println",fmt.Sprintln("\nDo you want to continue (Y/n)?\t")).Once()
	mockPrinter.On("Println",fmt.Sprintln("Skipped configuring octavius client")).Once()

	testConfigCmd.SetArgs([]string{"--cp-host", "localhost:5050", "--email-id", "jaimin.rathod@go-jek.com", "--time-out", "10", "--token", "AllowMe"})
	testConfigCmd.Execute()

	mockFileUtil.AssertExpectations(t)
}

func TestConfigCmdForConfigFileExistForConfigFileWritingFailure(t *testing.T) {
	mockFileUtil := new(fileUtil.MockFileUtil)
	mockPrinter := new(printer.MockPrinter)
	testConfigCmd := NewCmd(mockFileUtil, mockPrinter)

	mockFileUtil.On("IsFileExist", "job_data_example/config/octavius_client.yaml").Return(true).Once()
	mockFileUtil.On("ReadFile", "job_data_example/config/octavius_client.yaml").Return("old content", nil).Once()
	mockFileUtil.On("GetUserInput").Return("Y\n", nil).Once()
	var configFileContent string
	configFileContent += fmt.Sprintf("%s: %s\n", config.OctaviusCPHost, "localhost:5050")
	configFileContent += fmt.Sprintf("%s: %s\n", config.EmailId, "jaimin.rathod@go-jek.com")
	configFileContent += fmt.Sprintf("%s: %s\n", config.AccessToken, "AllowMe")
	configFileContent += fmt.Sprintf("%s: %v\n", config.ConnectionTimeoutSecs, 10)
	mockFileUtil.On("WriteFile", "job_data_example/config/octavius_client.yaml", configFileContent).Return(errors.New("failed to write file")).Once()
	mockPrinter.On("Println",fmt.Sprintln("[Warning] This will overwrite current config:")).Once()
	mockPrinter.On("Println",fmt.Sprintln("old content")).Once()
	mockPrinter.On("Println",fmt.Sprintln("\nDo you want to continue (Y/n)?\t")).Once()
	mockPrinter.On("Println",fmt.Sprintf("Error writing content %v \n to config file %s: %s\n", configFileContent, "job_data_example/config/octavius_client.yaml", "failed to write file")).Once()


	testConfigCmd.SetArgs([]string{"--cp-host", "localhost:5050", "--email-id", "jaimin.rathod@go-jek.com", "--time-out", "10", "--token", "AllowMe"})
	testConfigCmd.Execute()

	mockFileUtil.AssertExpectations(t)
}
