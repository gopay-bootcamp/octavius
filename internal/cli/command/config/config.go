package config

import (
	"fmt"
	"octavius/internal/cli/config"
	"octavius/internal/pkg/file"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/printer"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

//NewCmd Returns an instance of Config command for Octavius Client
func NewCmd(fileUtil file.File) *cobra.Command {
	var (
		cpHost                string
		emailID               string
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
				log.Warn(fmt.Sprintln("[Warning] This will overwrite current config:"))
				existingOctaviusConfig, err := fileUtil.ReadFile(configFilePath)
				if err != nil {
					log.Error(err, fmt.Sprintln(existingOctaviusConfig))
					return
				}

				printer.Println(fmt.Sprintln("\nDo you want to continue (Y/n)?\t"), color.FgYellow)
				userPermission, err := fileUtil.GetUserInput()
				if err != nil {
					log.Error(err, fmt.Sprintln("error getting user permission for overwriting config"))
					return
				}

				if userPermission != "y\n" && userPermission != "Y\n" {
					log.Info(fmt.Sprintln("Skipped configuring octavius client"))
					return
				}
			} else {
				err := fileUtil.CreateDirIfNotExist(config.ConfigFileDir())
				if err != nil {
					log.Error(err, "error in creating config file directory")
					return
				}
				err = fileUtil.CreateFile(configFilePath)
				if err != nil {
					log.Error(err, "error in creating config file")
					return
				}
			}

			var configFileContent string
			configFileContent += fmt.Sprintf("%s: %s\n", config.OctaviusCPHost, cpHost)
			//TODO: config.EmailId should be config.EmailID
			configFileContent += fmt.Sprintf("%s: %s\n", config.EmailId, emailID)
			configFileContent += fmt.Sprintf("%s: %s\n", config.AccessToken, accessToken)
			configFileContent += fmt.Sprintf("%s: %v\n", config.ConnectionTimeoutSecs, connectionTimeOutSecs)

			err := fileUtil.WriteFile(configFilePath, configFileContent)
			if err != nil {
				log.Error(err, fmt.Sprintf("error writing content %v to config file %s \n", configFileContent, configFilePath))
				return
			}

			log.Info("Octavius client configured successfully")
		},
	}
	configCmd.Flags().StringVarP(&cpHost, "cp-host", "", "", "CP_HOST port address(required)")
	configCmd.Flags().StringVarP(&emailID, "email-id", "", "", "Client Email-id")
	configCmd.Flags().StringVarP(&accessToken, "token", "", "", "Client Access Token")
	configCmd.Flags().IntVarP(&connectionTimeOutSecs, "time-out", "", 0, "Connection Time Out In Sec")
	configCmd.MarkFlagRequired("cp-host")
	return configCmd
}
