package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name    string `mapstructure:"name"`
		Version string `mapstructure:"version"`
	} `mapstructure:"app"`

	Cloud struct {
		Provider string `mapstructure:"provider"`
		Region   string `mapstructure:"region"`
	} `mapstructure:"cloud"`
}

// Config provides a Viper instance for configuration management.
func InitConfig() (*viper.Viper, error) {
	var env string

	// Parse --env argument
	for i, arg := range os.Args {
		if arg == "--env" && i+1 < len(os.Args) {
			env = os.Args[i+1]
			break
		}
	}

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	if env != "" {
		overridePath := fmt.Sprintf("configs/config.%s.yaml", env)
		if _, err := os.Stat(overridePath); err == nil {
			v.SetConfigFile(overridePath)
			if err := v.MergeInConfig(); err != nil {
				return nil, fmt.Errorf("failed to merge override config: %w", err)
			}
		}
	}

	return v, nil
}

func GetConfig(cfgViper *viper.Viper) (*Config, error) {
	var cfgObj Config
	if err := cfgViper.Unmarshal(&cfgObj); err != nil {
		return nil, err
	}

	return &cfgObj, nil
}
