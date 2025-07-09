package cli

import (
	"github.com/spf13/cobra"
	"go.uber.org/config"
)

// DigRootCmd provides the root Cobra command with injected config.Provider and command register.
func DigRootCmd(cfg config.Provider, commandRegister []CommandProvider) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "cloudstart",
		Short: "cloudstart is a CLI tool for cloud operations",
	}
	rootCmd.PersistentFlags().String("env", "", "Environment (e.g. local)")

	for _, provider := range commandRegister {
		rootCmd.AddCommand(provider(cfg))
	}
	return rootCmd
}
