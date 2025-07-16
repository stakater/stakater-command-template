package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// DigRootCmd provides the root Cobra command with injected *viper.Viper and command register.
func DigRootCmd(cfgViper *viper.Viper, commandRegister []CommandProvider) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "cloudstart",
		Short: "cloudstart is a CLI tool for cloud operations",
	}
	rootCmd.PersistentFlags().String("env", "", "Environment (e.g. local)")

	for _, provider := range commandRegister {
		rootCmd.AddCommand(provider(cfgViper))
	}
	return rootCmd
}
