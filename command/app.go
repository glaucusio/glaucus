package command

import (
	"context"

	"github.com/spf13/cobra"
)

type App struct {
	context.Context

	Version string
	GitHead string
	Date    string
}

func NewApp(ctx context.Context) *App {
	return &App{
		Context: ctx,
		GitHead: buildGitHead,
		Date:    buildDate,
	}
}

func (a *App) RegisterRoot(root *cobra.Command) {
	// todo
}
