package cmd

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

var Version = "1.0.2"

var app = &cli.Command{
	Name:    "devhttps",
	Usage:   "Easy HTTPS reverse proxy for local development",
	Version: Version,
	Commands: []*cli.Command{
		newAddCommand(),
		newRunCommand(),
		newShowCommand(),
		newCheckCommand(),
		newHttpCommand(),
	},
}

func Execute() {
	app.Run(context.Background(), os.Args)
}
