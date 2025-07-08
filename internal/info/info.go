// Package info provides the internal implementation of the 'info' CLI command.
package info

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/config"
)

type ProviderHolder struct {
	Provider config.Provider
}

func NewInfoCommand(holder *ProviderHolder) *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Print parsed config information",
		Run: func(cmd *cobra.Command, args []string) {
			if holder.Provider == nil {
				fmt.Println("Config not loaded")
				os.Exit(1)
			}
			fmt.Printf("App Name: %s\n", holder.Provider.Get("app.name").String())
			fmt.Printf("App Version: %s\n", holder.Provider.Get("app.version").String())
			fmt.Printf("Cloud Provider: %s\n", holder.Provider.Get("cloud.provider").String())
			fmt.Printf("Cloud Region: %s\n", holder.Provider.Get("cloud.region").String())
		},
	}
}
