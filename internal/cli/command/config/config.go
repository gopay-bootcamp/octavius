//Package config provides cli command to set client side configurations
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
		Args:    cobra.MaximumNArgs(0),

		Run: func(cmd *cobra.Command, args []string) {

			configFilePath := filepath.Join(config.ConfigFileDir(), "octavius_client.yaml")
			isConfigFileExist := fileUtil.IsFileExist(configFilePath)

			if isConfigFileExist == true {
				log.Warn(fmt.Sprintln("[Warning] This will overwrite current config:"))
				printer.Println("[Warning] This will overwrite the current config.", color.FgYellow)

				existingOctaviusConfig, err := fileUtil.ReadFile(configFilePath)
				if err != nil {
					log.Error(err, fmt.Sprintln(existingOctaviusConfig))
					printer.Println("error while reading the existing configurations", color.FgRed)
					return
				}

				printer.Println("Do you want to continue (Y/n)?\t", color.FgYellow)
				userPermission, err := fileUtil.GetUserInput()
				if err != nil {
					log.Error(err, "error getting user permission for overwriting config")
					printer.Println("error while getting the user permission", color.FgRed)
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
					printer.Println("error in creating config file directory", color.FgRed)
					return
				}
				err = fileUtil.CreateFile(configFilePath)
				if err != nil {
					log.Error(err, "error in creating config file")
					printer.Println("error in creating config file", color.FgRed)
					return
				}
			}
			printer.Println("Applying the configurations.", color.FgBlack)

			var configFileContent string
			configFileContent += fmt.Sprintf("%s: %s\n", config.OctaviusCPHost, cpHost)
			configFileContent += fmt.Sprintf("%s: %s\n", config.EmailID, emailID)
			configFileContent += fmt.Sprintf("%s: %s\n", config.AccessToken, accessToken)
			configFileContent += fmt.Sprintf("%s: %v\n", config.ConnectionTimeoutSecs, connectionTimeOutSecs)

			err := fileUtil.WriteFile(configFilePath, configFileContent)
			if err != nil {
				log.Error(err, fmt.Sprintf("error writing content %v to config file %s \n", configFileContent, configFilePath))
				printer.Println("error in writing the configurations", color.FgRed)
				return
			}

			log.Info("Octavius client configured successfully")
			printer.Println("Octavius client configured successfully.", color.FgGreen)
		},
	}
	configCmd.Flags().StringVarP(&cpHost, "cp-host", "", "", "CP_HOST port address(required)")
	configCmd.Flags().StringVarP(&emailID, "email-id", "", "", "Client Email-id")
	configCmd.Flags().StringVarP(&accessToken, "token", "", "", "Client Access Token")
	configCmd.Flags().IntVarP(&connectionTimeOutSecs, "time-out", "", 0, "Connection Time Out In Sec")
	err := configCmd.MarkFlagRequired("cp-host")
	if err != nil {
		log.Error(err, "error while setting the flag required")
		printer.Println("error while setting the flag required", color.FgRed)
		return nil
	}
	return configCmd
}
