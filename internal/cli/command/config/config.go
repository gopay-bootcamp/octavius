package config

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"octavius/internal/cli/config"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
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
	return &cobra.Command{
		Use:     "config",
		Short:   "Configure octavius client",
		Long:    "This command helps configure client with control plane host, email id and access token",
		Example: fmt.Sprintf("octavius config %s=example.octavius.com %s=example@octavius.com %s=XXXXX", config.OctaviusCPHost, config.EmailId, config.AccessToken),
		Args:    cobra.MinimumNArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			configFile := filepath.Join(config.ConfigFileDir(), "octavius_client.yaml")
			if _, err := os.Stat(configFile); err == nil {
				fmt.Println("[Warning] This will overwrite current config:")
				existingOctaviusConfig, err := ioutil.ReadFile(configFile)
				if err != nil {
					fmt.Println(fmt.Sprintf("Error reading config file: %s", configFile), color.FgRed)
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
			for _, v := range args {
				arg := strings.Split(v, "=")

				if len(arg) != 2 {
					fmt.Println(fmt.Sprintf("\nIncorrect config key-value pair format: %s. Correct format: CONFIG_KEY=VALUE\n", v))
					return
				}

				switch arg[0] {
				case config.OctaviusCPHost:
					configFileContent += fmt.Sprintf("%s: %s\n", config.OctaviusCPHost, arg[1])
				case config.EmailId:
					configFileContent += fmt.Sprintf("%s: %s\n", config.EmailId, arg[1])
				case config.AccessToken:
					configFileContent += fmt.Sprintf("%s: %s\n", config.AccessToken, arg[1])
				case config.ConnectionTimeoutSecs:
					configFileContent += fmt.Sprintf("%s: %s\n", config.ConnectionTimeoutSecs, arg[1])
				default:
					fmt.Println(fmt.Sprintf("Octavius doesn't support config key: %s", arg[0]))
				}
			}

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
}
