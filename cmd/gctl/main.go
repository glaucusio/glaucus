package main

import (
	"context"
	"fmt"
	"os"

	"github.com/glaucusio/glaucus/command"
	"github.com/glaucusio/glaucus/gctl/cmd"
)

func main() {
	ctx, cancel := command.WithSignals(context.Background())
	defer cancel()

	app := command.NewApp(ctx)
	root := cmd.NewGctlCommand(app)

	app.RegisterRoot(root)

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(command.ExitCode(err))
	}
}
