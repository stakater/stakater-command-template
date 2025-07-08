package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/config"

	"stakater-cmd/internal/info"
)

func newRootCmd() *cobra.Command {
	var env string
	providerHolder := &info.ProviderHolder{}

	rootCmd := &cobra.Command{
		Use:   "cloudstart",
		Short: "cloudstart is a CLI tool for cloud operations",
	}
	rootCmd.PersistentFlags().StringVar(&env, "env", "", "Environment (e.g. local)")

	infoCmd := info.NewInfoCommand(providerHolder)
	rootCmd.AddCommand(infoCmd)

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		var err error
		if env != "" {
			overridePath := fmt.Sprintf("configs/config.%s.yaml", env)
			providerHolder.Provider, err = config.NewYAML(
				config.File("configs/config.yaml"),
				config.File(overridePath),
			)
		} else {
			providerHolder.Provider, err = config.NewYAML(config.File("configs/config.yaml"))
		}
		return err
	}

	return rootCmd
}

func Execute() {
	rootCmd := newRootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
