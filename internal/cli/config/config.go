package config

import (
	"errors"
	"fmt"
	"octavius/internal/pkg/constant"
	"os"
	"time"

	"github.com/spf13/viper"
)

const (
	Environment           = "ENVIRONMENT"
	OctaviusCPHost        = "CP_HOST"
	EmailID               = "EMAIL_ID"
	AccessToken           = "ACCESS_TOKEN"
	ConnectionTimeoutSecs = "CONNECTION_TIMEOUT_SECS"
	LogFilePath           = "./cli.log"
)

// OctaviusConfig Struct containing Octavius Client details
type OctaviusConfig struct {
	Host                  string
	Email                 string
	AccessToken           string
	ConnectionTimeoutSecs time.Duration
}

// ConfigError Struct containing error and Message to solve the Client Config Error
type ConfigError struct {
	error
	Message string
}

// RootError Returns the error inside the ConfigError
func (c *ConfigError) RootError() error {
	return c.error
}

// Loader Octavius Client Loader interface with Method: Load()
type Loader interface {
	Load() (OctaviusConfig, ConfigError)
}

type loader struct{}

// NewLoader Returns an instance of Octavius Client Config Loader class
func NewLoader() Loader {
	return &loader{}
}

//Load Loads the config by reading from octavius_client.yaml and returning Config and Error
func (loader *loader) Load() (OctaviusConfig, ConfigError) {
	viper.SetDefault(ConnectionTimeoutSecs, 10)
	viper.AutomaticEnv()

	viper.AddConfigPath(ConfigFileDir())
	viper.SetConfigName("octavius_client")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		configFileUsed := viper.ConfigFileUsed()
		message := ""
		if _, err := os.Stat(configFileUsed); os.IsNotExist(err) {
			message = fmt.Sprintf("Config file not found in %s/octavius_client.yaml\n", ConfigFileDir())
			message += fmt.Sprintf("Setup config using `octavius config --cp-host some.host ...`\n\n")
		}
		return OctaviusConfig{}, ConfigError{error: err, Message: message}
	}
	octaviusCPHost := viper.GetString(OctaviusCPHost)
	if octaviusCPHost == "" {
		return OctaviusConfig{}, ConfigError{error: errors.New("Mandatory Config Missing"), Message: constant.ConfigOctaviusHostMissingError}
	}
	emailID := viper.GetString(EmailID)
	accessToken := viper.GetString(AccessToken)
	connectionTimeout := time.Duration(viper.GetInt(ConnectionTimeoutSecs)) * time.Second
	return OctaviusConfig{
		Host:                  octaviusCPHost,
		Email:                 emailID,
		AccessToken:           accessToken,
		ConnectionTimeoutSecs: connectionTimeout,
	}, ConfigError{}
}

// ConfigFileDir Returns Config file directory
// This allows to test on dev environment without conflicting with installed octavius config file
func ConfigFileDir() string {
	// localConfigDir, localConfigAvailable := os.LookupEnv("LOCAL_CONFIG_DIR")
	// if localConfigAvailable {
	// 	return localConfigDir
	// } else if os.Getenv(Environment) == "test" {
	// 	return "/tmp"
	// } else {
	// 	return fmt.Sprintf("%s/.octavius", os.Getenv("HOME"))
	// }
	return "./job_data_example/config"
}
