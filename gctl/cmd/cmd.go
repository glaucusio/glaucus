package cmd

import (
	"github.com/glaucusio/glaucus/command"

	"github.com/spf13/cobra"
)

func NewGctlCommand(app *command.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gctl",
		Short: "Command-line interface to Glaucus service",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
	// todo
	)

	return cmd
}
