package cmd

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

var app = &cli.Command{
	Name:  "devhttps",
	Usage: "Easy HTTPS reverse proxy for local development",
	Commands: []*cli.Command{
		newAddCommand(),
		newShowCommand(),
	},
}

func Execute() {
	app.Run(context.Background(), os.Args)
}
