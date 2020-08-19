package config

import (
	"fmt"
	"octavius/internal/cli/config"
	"octavius/internal/cli/fileUtil"
	"octavius/internal/cli/printer"
	"path/filepath"

	"github.com/spf13/cobra"
)

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
				fmt.Println("[Warning] This will overwrite current config:")
				existingOctaviusConfig, err := fileUtil.ReadFile(configFilePath)
				if err != nil {
					fmt.Printf("Error reading config file: %v\n", err)
					return
				}
				fmt.Println(existingOctaviusConfig)

				fmt.Println("\nDo you want to continue (Y/n)?\t")

				userPermission, err := fileUtil.GetUserInput()
				if err != nil {
					fmt.Println("Error getting user permission for overwriting config")
					return
				}

				if userPermission != "y\n" && userPermission != "Y\n" {
					fmt.Println("Skipped configuring octavius client")
					return
				}
			} else {
				err := fileUtil.CreateDirIfNotExist(config.ConfigFileDir())
				if err != nil {
					fmt.Printf("Error in creating config file directory, %v\n", err)
					return
				}
				err = fileUtil.CreateFile(configFilePath)
				if err != nil {
					fmt.Printf("Error in creating config file, %v\n", err)
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
				fmt.Printf("Error writing content %v \n to config file %s: %s\n", configFileContent, configFilePath, err)
				return
			}

			fmt.Println("Octavius client configured successfully")
		},
	}
	configCmd.Flags().StringVarP(&cpHost, "cp-host", "", "", "CP_HOST port address(required)")
	configCmd.Flags().StringVarP(&emailId, "email-id", "", "", "Client Email-id")
	configCmd.Flags().StringVarP(&accessToken, "token", "", "", "Client Access Token")
	configCmd.Flags().IntVarP(&connectionTimeOutSecs, "time-out", "", 0, "Connection Time Out In Sec")
	configCmd.MarkFlagRequired("cp-host")
	return configCmd
}
