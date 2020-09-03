package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"sync"
	"sync/atomic"
)

var once sync.Once
var config KubernetesConfig

func GetMapFromJson(viper *viper.Viper, key string) map[string]string {
	var jsonStr = []byte(viper.GetString(key))
	var annotations map[string]string

	err := json.Unmarshal(jsonStr, &annotations)
	if err != nil {
		_ = fmt.Errorf("invalid Value for key %s, errors %v", key, err.Error())
	}

	return annotations
}

type KubernetesConfig struct {
	viper                        *viper.Viper
	KubeConfig                   string
	KubeContext                  string
	DefaultNamespace             string
	KubeServiceAccountName       string
	JobPodAnnotations            map[string]string
	KubeJobActiveDeadlineSeconds int
	KubeJobRetries               int
	KubeWaitForResourcePollCount int
}

func load() KubernetesConfig {
	fang := viper.New()

	fang.SetConfigType("json")
	fang.SetConfigName("kubernetes_config")
	fang.AddConfigPath(".")

	value, available := os.LookupEnv("CONFIG_LOCATION")
	if available {
		fang.AddConfigPath(value)
	}

	_ = fang.ReadInConfig()

	kubeConfig := KubernetesConfig{
		viper:                        fang,
		KubeConfig:                   fang.GetString("config"),
		KubeContext:                  fang.GetString("context"),
		DefaultNamespace:             fang.GetString("default_namespace"),
		KubeServiceAccountName:       fang.GetString("service_account_name"),
		JobPodAnnotations:            GetMapFromJson(fang, "job_pod_annotations"),
		KubeJobActiveDeadlineSeconds: fang.GetInt("job_active_deadline_seconds"),
		KubeJobRetries:               fang.GetInt("job_retries"),
		KubeWaitForResourcePollCount: fang.GetInt("wait_for_resource_poll_count"),
	}

	return kubeConfig
}

/*
 *	Instead of using AtomBool, can we use channel and have a separate
 * 	goroutine to reset config?
 */

type AtomBool struct{ flag int32 }

func (b *AtomBool) Set(value bool) {
	var i int32 = 0
	if value {
		i = 1
	}
	atomic.StoreInt32(&(b.flag), i)
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

func Config() KubernetesConfig {
	once.Do(func() {
		config = load()
	})

	if reset.Get() {
		config = load()
		reset.Set(false)
	}

	return config
}
