package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/spf13/viper"
)

func GetStringDefault(viper *viper.Viper, key string, defaultValue string) string {
	viper.SetDefault(key, defaultValue)
	return viper.GetString(key)
}

//will be utilized in further implementation, line 18-53
func GetArrayString(viper *viper.Viper, key string) []string {
	return strings.Split(viper.GetString(key), ",")
}

func GetArrayStringDefault(viper *viper.Viper, key string, defaultValue []string) []string {
	viper.SetDefault(key, strings.Join(defaultValue, ","))
	return strings.Split(viper.GetString(key), ",")
}

func GetBoolDefault(viper *viper.Viper, key string, defaultValue bool) bool {
	viper.SetDefault(key, defaultValue)
	return viper.GetBool(key)
}

func GetInt64Ref(viper *viper.Viper, key string) *int64 {
	value := viper.GetInt64(key)
	return &value
}

func GetInt32Ref(viper *viper.Viper, key string) *int32 {
	value := viper.GetInt32(key)
	return &value
}

func GetMapFromJson(viper *viper.Viper, key string) map[string]string {
	var jsonStr = []byte(viper.GetString(key))
	var annotations map[string]string

	err := json.Unmarshal(jsonStr, &annotations)
	if err != nil {
		_ = fmt.Errorf("invalid Value for key %s, errors %v", key, err.Error())
	}

	return annotations
}

var once sync.Once
var config ProctorConfig

type ProctorConfig struct {
	viper    *viper.Viper
	LogLevel string
	AppPort  string
	EtcdPort string
}

func load() ProctorConfig {
	fang := viper.New()

	fang.SetConfigType("json")
	fang.SetConfigName("config")
	fang.AddConfigPath(".")
	value, available := os.LookupEnv("CONFIG_LOCATION")
	if available {
		fang.AddConfigPath(value)
	}
	//will be nil if file is read properly
	err := fang.ReadInConfig()
	if err != nil {
		fmt.Println("file not read", err)
	}
	proctorConfig := ProctorConfig{
		viper:    fang,
		LogLevel: GetStringDefault(fang, "log_level", "info"),
		EtcdPort: fang.GetString("etcd_port"),
		AppPort:  fang.GetString("app_port"),
	}
	return proctorConfig
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
	if atomic.LoadInt32(&(b.flag)) != 0 {
		return true
	}
	return false
}

var reset = new(AtomBool)

func init() {
	reset.Set(false)
}

func Reset() {
	reset.Set(true)
}

func Config() ProctorConfig {
	once.Do(func() {
		config = load()
	})

	if reset.Get() {
		config = load()
		reset.Set(false)
	}
	return config
}
