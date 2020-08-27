package config

import (
	"fmt"
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
var config OctaviusConfig

type OctaviusConfig struct {
	viper                *viper.Viper
	LogLevel             string
	AppPort              string
	EtcdPort             string
	ExecutorPingDeadline time.Duration
}

func load() OctaviusConfig {
	fang := viper.New()

	fang.SetConfigType("json")
	fang.SetConfigName("controller_config")
	fang.AddConfigPath(".")

	//will be nil if file is read properly
	err := fang.ReadInConfig()
	if err != nil {
		fmt.Println("file not read", err)
	}
	octaviusConfig := OctaviusConfig{
		viper:                fang,
		LogLevel:             GetStringDefault(fang, "log_level", "info"),
		EtcdPort:             fang.GetString("etcd_port"),
		AppPort:              fang.GetString("app_port"),
		ExecutorPingDeadline: time.Duration(GetIntDefault(fang, "executor_ping_deadline", 30)) * time.Second,
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

func Config() OctaviusConfig {
	once.Do(func() {
		config = load()
	})

	if reset.Get() {
		config = load()
		reset.Set(false)
	}
	return config
}
