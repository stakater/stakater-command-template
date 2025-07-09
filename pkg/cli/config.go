package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config provides a Viper instance for configuration management.
func Config() (*viper.Viper, error) {
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
