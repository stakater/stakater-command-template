// Package info provides the internal implementation of the 'info' CLI command.
package info

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewInfoCommand(cfg *viper.Viper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Print parsed config information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("App Name: %s\n", cfg.GetString("app.name"))
			fmt.Printf("App Version: %s\n", cfg.GetString("app.version"))
			fmt.Printf("Cloud Provider: %s\n", cfg.GetString("cloud.provider"))
			fmt.Printf("Cloud Region: %s\n", cfg.GetString("cloud.region"))
		},
	}

	cmd.Flags().String("name", "", "App name")
	cmd.Flags().String("version", "", "App version")
	cmd.Flags().String("provider", "", "Cloud provider")
	cmd.Flags().String("region", "", "Cloud region")

	// Bind flags to Viper
	cfg.BindPFlag("app.name", cmd.Flags().Lookup("name"))
	cfg.BindPFlag("app.version", cmd.Flags().Lookup("version"))
	cfg.BindPFlag("cloud.provider", cmd.Flags().Lookup("provider"))
	cfg.BindPFlag("cloud.region", cmd.Flags().Lookup("region"))

	return cmd
}
