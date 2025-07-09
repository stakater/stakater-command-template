// Package info provides the internal implementation of the 'info' CLI command.
package info

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/config"
)

func NewInfoCommand(cfg config.Provider) *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Print parsed config information",
		Run: func(cmd *cobra.Command, args []string) {
			if cfg == nil {
				fmt.Println("Config not loaded")
				os.Exit(1)
			}
			fmt.Printf("App Name: %s\n", cfg.Get("app.name").String())
			fmt.Printf("App Version: %s\n", cfg.Get("app.version").String())
			fmt.Printf("Cloud Provider: %s\n", cfg.Get("cloud.provider").String())
			fmt.Printf("Cloud Region: %s\n", cfg.Get("cloud.region").String())
		},
	}
}
