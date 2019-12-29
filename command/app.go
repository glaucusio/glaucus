package command

import (
	"context"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

type App struct {
	context.Context
	Group   *errgroup.Group
	GitHead string
	Date    string
}

func NewApp(ctx context.Context) *App {
	eg, ctx := errgroup.WithContext(ctx)

	return &App{
		Context: ctx,
		Group:   eg,
		GitHead: buildGitHead,
		Date:    buildDate,
	}
}

func (a *App) RegisterRoot(root *cobra.Command) {
	// todo
}
