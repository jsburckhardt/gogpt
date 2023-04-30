//Will look in a config.env file or the actual system env.

package gogptconfig

import (
	"fmt"
	"gogpt/internal/directory"
	"gogpt/internal/logger"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

// Config is a central config object for the app
type Config struct {
}

var once sync.Once

var (
	instance *Config
)

// IsDevMode will return true if dev is set in the APP_MODE config
func (s *Config) IsDevMode() bool {
	return s.GetString(APP_MODE) == "dev"
}

// InCluster will return true if the code is running in a kube cluster
func (s *Config) InCluster() bool {
	return os.Getenv("KUBERNETES_SERVICE_HOST") != ""
}

// GetString returns the key value as a string
func (s *Config) GetString(key ConfigName) string {
	return viper.GetString(string(key))
}

// GetStringDefault returns the key value as a string, or the default value if not found
func (s *Config) GetStringDefault(key ConfigName, defaultValue string) string {
	val := viper.GetString(string(key))
	if val != "" {
		return val
	}
	return defaultValue
}

// GetBool returns the key value as a bool
func (s *Config) GetBool(key ConfigName) bool {
	return viper.GetBool(string(key))
}

// GetInt returns the key value as an int
func (s *Config) GetInt(key ConfigName) int64 {
	return viper.GetInt64(string(key))
}

// GetConfig returns the current gogpt config object. This object is a shared singleton. This will panic if there are problems
func GetConfig() (config *Config) {
	once.Do(func() { // <-- atomic, does not allow repeating

		log := logger.GetInstance()

		ex, err := os.Executable()

		if err != nil {
			panic(fmt.Errorf("Config failure: %w", err))
		}

		exPath := filepath.Dir(ex)
		exPath, err = directory.FileWalk(exPath, "gogpt.conf")

		if err == nil {
			viper.SetConfigName("gogpt.conf")
		}

		viper.AutomaticEnv()

		if exPath != "" {
			log.Debugf("gogpt config path: %s", exPath)
			viper.AddConfigPath(exPath)
		}

		viper.SetConfigType("env")

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Debugf("gogpt.conf config file not found")
			} else {
				panic(fmt.Errorf("Fatal error config file %w ", err))
			}
		}

		instance = &Config{}
	})

	if instance == nil {
		panic(fmt.Errorf("Config failure"))
	}

	return instance
}
