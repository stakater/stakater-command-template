package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"stakater-cmd/internal/info"
)

type CommandProvider func(cfg *viper.Viper) *cobra.Command

// CommandRegister returns all available command providers for DI injection.
func CommandRegister() []CommandProvider {
	return []CommandProvider{
		info.NewInfoCommand,
	}
}
