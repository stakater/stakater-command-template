// Package info provides the internal implementation of the 'info' CLI command.
package info

import (
	"fmt"
	"stakater-cmd/internal/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewInfoCommand(cfgViper *viper.Viper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Print parsed config information",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.GetConfig(cfgViper)
			if err != nil {
				fmt.Println("Could not load config")
				return
			}

			fmt.Printf("App Name: %s\n", cfg.App.Name)
			fmt.Printf("App Version: %s\n", cfg.App.Version)
			fmt.Printf("Cloud Provider: %s\n", cfg.Cloud.Provider)
			fmt.Printf("Cloud Region: %s\n", cfg.Cloud.Region)
		},
	}

	cmd.Flags().String("name", "", "App name")
	cmd.Flags().String("version", "", "App version")
	cmd.Flags().String("provider", "", "Cloud provider")
	cmd.Flags().String("region", "", "Cloud region")

	// Bind flags to Viper
	cfgViper.BindPFlag("app.name", cmd.Flags().Lookup("name"))
	cfgViper.BindPFlag("app.version", cmd.Flags().Lookup("version"))
	cfgViper.BindPFlag("cloud.provider", cmd.Flags().Lookup("provider"))
	cfgViper.BindPFlag("cloud.region", cmd.Flags().Lookup("region"))

	return cmd
}
