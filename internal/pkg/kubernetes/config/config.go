package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

var once sync.Once
var config KubernetesConfig

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

type KubernetesConfig struct {
	viper 							*viper.Viper
	KubeConfig 						string
	KubeContext                     string
	DefaultNamespace                string
	KubeServiceAccountName          string
	JobPodAnnotations				map[string]string
	KubeJobActiveDeadlineSeconds	*int64
	KubeJobRetries					*int32
	KubeWaitForResourcePollCount 	int
}

func load() KubernetesConfig {
	fang := viper.New()
	fang.SetEnvPrefix("OCTAVIUS")
	fang.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fang.AutomaticEnv()

	fang.SetConfigType("json")
	fang.SetConfigName("kubernetes_config")
	fang.AddConfigPath(".")

	value, available := os.LookupEnv("CONFIG_LOCATION")
	if available {
		fang.AddConfigPath(value)
	}

	_ = fang.ReadInConfig()

	kubeConfig := KubernetesConfig{
		viper:	fang,
		KubeConfig: fang.GetString("config"),
		KubeContext: fang.GetString("context"),
		DefaultNamespace: fang.GetString("default_namespace"),
		KubeServiceAccountName: fang.GetString("service_account_name"),
		JobPodAnnotations: GetMapFromJson(fang, "job_pod_annotations"),
		KubeJobActiveDeadlineSeconds: GetInt64Ref(fang, "job_active_deadline_seconds"),
		KubeJobRetries: GetInt32Ref(fang, "job_retries"),
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