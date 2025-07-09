package cli

import (
	"github.com/spf13/cobra"
	"go.uber.org/config"

	"stakater-cmd/internal/info"
)

type CommandProvider func(cfg config.Provider) *cobra.Command

// CommandRegister returns all available command providers for DI injection.
func CommandRegister() []CommandProvider {
	return []CommandProvider{
		info.NewInfoCommand,
	}
}
