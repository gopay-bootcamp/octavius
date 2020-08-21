package config

import (
	"fmt"
	"octavius/internal/cli/config"
	"octavius/internal/cli/fileUtil"
	"octavius/internal/cli/logger"
	"octavius/internal/cli/printer"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

//NewCmd Returns an instance of Config command for Octavius Client
func NewCmd(fileUtil fileUtil.FileUtil, printer printer.Printer) *cobra.Command {
	var (
		cpHost                string
		emailId               string
		accessToken           string
		connectionTimeOutSecs int
	)
	configCmd := &cobra.Command{
		Use:     "config",
		Short:   "Configure octavius client",
		Long:    "This command helps configure client with control plane host, email id and access token",
		Example: "octavius config [flags]",

		Run: func(cmd *cobra.Command, args []string) {
			configFilePath := filepath.Join(config.ConfigFileDir(), "octavius_client.yaml")
			isConfigFileExist := fileUtil.IsFileExist(configFilePath)

			if isConfigFileExist == true {
				logger.Warn(fmt.Sprintln("[Warning] This will overwrite current config:"), printer)
				existingOctaviusConfig, err := fileUtil.ReadFile(configFilePath)
				logger.Error(err, fmt.Sprintln(existingOctaviusConfig), printer)
				printer.Println(fmt.Sprintln("\nDo you want to continue (Y/n)?\t"), color.FgYellow)
				userPermission, err := fileUtil.GetUserInput()
				if err != nil {
					logger.Error(err, fmt.Sprintln("error getting user permission for overwriting config"), printer)
					return
				}

				if userPermission != "y\n" && userPermission != "Y\n" {
					logger.Info(fmt.Sprintln("Skipped configuring octavius client"), printer)
					return
				}
			} else {
				err := fileUtil.CreateDirIfNotExist(config.ConfigFileDir())
				if err != nil {
					logger.Error(err, "Error in creating config file directory", printer)
					return
				}
				err = fileUtil.CreateFile(configFilePath)
				if err != nil {
					logger.Error(err, "Error in creating config file", printer)
					return
				}
			}

			var configFileContent string
			configFileContent += fmt.Sprintf("%s: %s\n", config.OctaviusCPHost, cpHost)
			configFileContent += fmt.Sprintf("%s: %s\n", config.EmailId, emailId)
			configFileContent += fmt.Sprintf("%s: %s\n", config.AccessToken, accessToken)
			configFileContent += fmt.Sprintf("%s: %v\n", config.ConnectionTimeoutSecs, connectionTimeOutSecs)

			err := fileUtil.WriteFile(configFilePath, configFileContent)
			if err != nil {
				logger.Error(err, fmt.Sprintf("Error writing content %v to config file %s \n", configFileContent, configFilePath), printer)
				return
			}

			logger.Info("Octavius client configured successfully", printer)
		},
	}
	configCmd.Flags().StringVarP(&cpHost, "cp-host", "", "", "CP_HOST port address(required)")
	configCmd.Flags().StringVarP(&emailId, "email-id", "", "", "Client Email-id")
	configCmd.Flags().StringVarP(&accessToken, "token", "", "", "Client Access Token")
	configCmd.Flags().IntVarP(&connectionTimeOutSecs, "time-out", "", 0, "Connection Time Out In Sec")
	configCmd.MarkFlagRequired("cp-host")
	return configCmd
}
