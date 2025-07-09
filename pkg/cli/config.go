package cli

import (
	"fmt"
	"os"

	"go.uber.org/config"
)

// Config provides a config.Provider based on the --env flag.
func Config() (config.Provider, error) {
	var env string
	for i, arg := range os.Args {
		if arg == "--env" && i+1 < len(os.Args) {
			env = os.Args[i+1]
			break
		}
	}
	if env != "" {
		overridePath := fmt.Sprintf("configs/config.%s.yaml", env)
		return config.NewYAML(
			config.File("configs/config.yaml"),
			config.File(overridePath),
		)
	}
	return config.NewYAML(config.File("configs/config.yaml"))
}
