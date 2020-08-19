package config

import (
	"errors"
	"fmt"
	"octavius/pkg/constant"
	"os"
	"time"

	"github.com/spf13/viper"
)

const (
	Environment           = "ENVIRONMENT"
	OctaviusCPHost        = "CP_HOST"
	EmailId               = "EMAIL_ID"
	AccessToken           = "ACCESS_TOKEN"
	ConnectionTimeoutSecs = "CONNECTION_TIMEOUT_SECS"
)

type OctaviusConfig struct {
	Host                  string
	Email                 string
	AccessToken           string
	ConnectionTimeoutSecs time.Duration
}

type ConfigError struct {
	error
	Message string
}

func (c *ConfigError) RootError() error {
	return c.error
}

type Loader interface {
	Load() (OctaviusConfig, ConfigError)
}

type loader struct{}

func NewLoader() Loader {
	return &loader{}
}

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
			message += fmt.Sprintf("Setup config using `octavius config CP_HOST=some.host ...`\n\n")
		}
		return OctaviusConfig{}, ConfigError{error: err, Message: message}
	}
	octaviusCPHost := viper.GetString(OctaviusCPHost)
	if octaviusCPHost == "" {
		return OctaviusConfig{}, ConfigError{error: errors.New("Mandatory Config Missing"), Message: constant.ConfigOctaviusHostMissingError}
	}
	emailId := viper.GetString(EmailId)
	accessToken := viper.GetString(AccessToken)
	connectionTimeout := time.Duration(viper.GetInt(ConnectionTimeoutSecs)) * time.Second
	return OctaviusConfig{
		Host:                  octaviusCPHost,
		Email:                 emailId,
		AccessToken:           accessToken,
		ConnectionTimeoutSecs: connectionTimeout,
	}, ConfigError{}
}

// Returns Config file directory
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
