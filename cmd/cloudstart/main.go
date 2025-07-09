package main

import (
	"log"

	"github.com/spf13/cobra"
	"go.uber.org/dig"

	"stakater-cmd/pkg/cli"
)

func main() {
	c := dig.New()
	// Provide dependencies
	c.Provide(cli.Config)
	c.Provide(cli.CommandRegister)
	c.Provide(cli.DigRootCmd)
	// Invoke CLI execution
	err := c.Invoke(func(rootCmd *cobra.Command) {
		if err := rootCmd.Execute(); err != nil {
			log.Fatal(err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}
