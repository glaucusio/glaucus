package cmd

import (
	"github.com/glaucusio/glaucus/command"
	"github.com/glaucusio/glaucus/glaucus"

	"github.com/spf13/cobra"
)

func NewGlaucusCommand(app *command.App) *cobra.Command {
	m := glaucusCmd{
		Service: &command.Service{
			App: app,
		},
	}

	cmd := &cobra.Command{
		Use:   "gctl",
		Short: "Command-line interface to Glaucus service",
		Args:  cobra.NoArgs,
		RunE:  m.run,
	}

	cmd.AddCommand(
	// todo
	)

	return cmd
}

type glaucusCmd struct {
	*command.Service
}

func (m *glaucusCmd) run(*cobra.Command, []string) error {
	h := new(glaucus.Handler)

	return m.Serve(h)
}
