package config

import (
	"encoding/json"
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

func GetMapFromJson(viper *viper.Viper, key string) (map[string]string, error) {
	var jsonStr = []byte(viper.GetString(key))
	var annotations map[string]string

	err := json.Unmarshal(jsonStr, &annotations)
	if err != nil {
		return nil, err
	}

	return annotations, nil
}

var once sync.Once
var config OctaviusExecutorConfig
var err error

type OctaviusExecutorConfig struct {
	viper                        *viper.Viper
	CPHost                       string
	ID                           string
	AccessToken                  string
	ConnTimeOutSec               time.Duration
	LogLevel                     string
	PingInterval                 time.Duration
	LogFilePath                  string
	LogFileSize                  int
	KubeConfig                   string
	KubeContext                  string
	DefaultNamespace             string
	KubeServiceAccountName       string
	JobPodAnnotations            map[string]string
	KubeJobActiveDeadlineSeconds int
	KubeJobRetries               int
	KubeWaitForResourcePollCount int
}

func load() (OctaviusExecutorConfig, error) {
	fang := viper.New()

	fang.SetConfigType("json")
	fang.SetConfigName("executor_config")
	fang.AddConfigPath(".")
	err := fang.ReadInConfig()
	if err != nil {
		return OctaviusExecutorConfig{}, err
	}

	JobPodAnnotation, err := GetMapFromJson(fang, "job_pod_annotations")
	if err != nil {
		return OctaviusExecutorConfig{}, err
	}
	octaviusConfig := OctaviusExecutorConfig{
		viper:                        fang,
		LogLevel:                     GetStringDefault(fang, "log_level", "info"),
		CPHost:                       fang.GetString("cp_host"),
		ID:                           fang.GetString("id"),
		AccessToken:                  fang.GetString("access_token"),
		ConnTimeOutSec:               time.Duration(GetIntDefault(fang, "conn_time_out", 10)) * time.Second,
		PingInterval:                 time.Duration(GetIntDefault(fang, "ping_interval", 30)) * time.Second,
		LogFilePath:                  GetStringDefault(fang, "log_file_path", "executor.log"),
		LogFileSize:                  fang.GetInt("log_file_max_size_in_mb"),
		KubeConfig:                   fang.GetString("kube_config"),
		KubeContext:                  fang.GetString("kube_context"),
		DefaultNamespace:             fang.GetString("default_namespace"),
		KubeServiceAccountName:       fang.GetString("service_account_name"),
		JobPodAnnotations:            JobPodAnnotation,
		KubeJobActiveDeadlineSeconds: fang.GetInt("job_active_deadline_seconds"),
		KubeJobRetries:               fang.GetInt("job_retries"),
		KubeWaitForResourcePollCount: fang.GetInt("wait_for_resource_poll_count"),
	}
	return octaviusConfig, nil
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

func Loader() (OctaviusExecutorConfig, error) {
	once.Do(func() {
		config, err = load()
	})

	if reset.Get() {
		config, err = load()
		reset.Set(false)
	}
	if err != nil {
		return OctaviusExecutorConfig{}, err
	}
	return config, nil
}
