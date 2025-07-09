package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/config"

	"stakater-cmd/internal/info"
)

// DigConfigProvider provides a config.Provider based on the --env flag.
func DigConfigProvider() (config.Provider, error) {
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

// DigRootCmd provides the root Cobra command with injected config.Provider.
func DigRootCmd(cfg config.Provider) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "cloudstart",
		Short: "cloudstart is a CLI tool for cloud operations",
	}
	rootCmd.PersistentFlags().String("env", "", "Environment (e.g. local)")
	rootCmd.AddCommand(info.NewInfoCommand(&info.ProviderHolder{Provider: cfg}))
	return rootCmd
}
