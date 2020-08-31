package config

import (
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/spf13/viper"
)

func GetStringDefault(viper *viper.Viper, key string, defaultValue string) string {
	viper.SetDefault(key, defaultValue)
	return viper.GetString(key)
}

func GetIntDefault(viper *viper.Viper, key string, defaultValue int) int {
	viper.SetDefault(key, defaultValue)
	return viper.GetInt(key)
}

var once sync.Once
var config OctaviusExecutorConfig

type OctaviusExecutorConfig struct {
	viper          *viper.Viper
	CPHost         string
	ID             string
	AccessToken    string
	ConnTimeOutSec time.Duration
	LogLevel       string
	PingInterval   time.Duration
	LogFilePath    string
}

func load() OctaviusExecutorConfig {
	fang := viper.New()

	fang.SetConfigType("json")
	fang.SetConfigName("executor_config")
	fang.AddConfigPath(".")

	value, available := os.LookupEnv("CONFIG_LOCATION")
	if available {
		fang.AddConfigPath(value)
	}
	//will be nil if file is read properly
	err := fang.ReadInConfig()
	if err != nil {
		return OctaviusExecutorConfig{}
	}
	octaviusConfig := OctaviusExecutorConfig{
		viper:          fang,
		LogLevel:       GetStringDefault(fang, "log_level", "info"),
		CPHost:         fang.GetString("cp_host"),
		ID:             fang.GetString("id"),
		AccessToken:    fang.GetString("access_token"),
		ConnTimeOutSec: time.Duration(GetIntDefault(fang, "conn_time_out", 10)) * time.Second,
		PingInterval:   time.Duration(GetIntDefault(fang, "ping_interval", 30)) * time.Second,
		LogFilePath:    GetStringDefault(fang, "log_file_path", "executor.log"),
	}
	return octaviusConfig
}

type AtomBool struct{ flag int32 }

func (b *AtomBool) Set(value bool) {
	var i int32 = 0
	if value {
		i = 1
	}
	atomic.StoreInt32(&(b.flag), int32(i))
}

func (b *AtomBool) Get() bool {
	return atomic.LoadInt32(&(b.flag)) != 0
}

var reset = new(AtomBool)

func init() {
	reset.Set(false)
}

func Reset() {
	reset.Set(true)
}

func Config() OctaviusExecutorConfig {
	once.Do(func() {
		config = load()
	})

	if reset.Get() {
		config = load()
		reset.Set(false)
	}
	return config
}
