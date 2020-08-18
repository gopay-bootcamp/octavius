package config

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"octavius/internal/cli/config"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func NewCmd() *cobra.Command {
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
			configFile := filepath.Join(config.ConfigFileDir(), "octavius_client.yaml")
			if _, err := os.Stat(configFile); err == nil {
				fmt.Println("[Warning] This will overwrite current config:")
				existingOctaviusConfig, err := ioutil.ReadFile(configFile)
				if err != nil {
					fmt.Printf("Error reading config file: %s in Control Plane", configFile)
					return
				}
				fmt.Println(string(existingOctaviusConfig))
				fmt.Println("\nDo you want to continue (Y/n)?\t")

				in := bufio.NewReader(os.Stdin)
				userPermission, err := in.ReadString('\n')

				if err != nil {
					fmt.Println("Error getting user permission for overwriting config")
					return
				}

				if userPermission != "y\n" && userPermission != "Y\n" {
					fmt.Println("Skipped configuring octavius client")
					return
				}
			}

			CreateDirIfNotExist(config.ConfigFileDir())
			var configFileContent string
			configFileContent += fmt.Sprintf("%s: %s\n", config.OctaviusCPHost, cpHost)
			configFileContent += fmt.Sprintf("%s: %s\n", config.EmailId, emailId)
			configFileContent += fmt.Sprintf("%s: %s\n", config.AccessToken, accessToken)
			configFileContent += fmt.Sprintf("%s: %v\n", config.ConnectionTimeoutSecs, connectionTimeOutSecs)

			configFileContentBytes := []byte(configFileContent)
			f, err := os.Create(configFile)
			if err != nil {
				fmt.Printf("Error creating config file %s: %s", configFile, err.Error())
				return
			}
			_, err = f.Write(configFileContentBytes)
			if err != nil {
				fmt.Printf("Error writing content %v \n to config file %s: %s", configFileContentBytes, configFile, err.Error())
				return
			}
			defer f.Close()
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
