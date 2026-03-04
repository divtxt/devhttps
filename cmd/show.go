package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func newShowCommand() *cli.Command {
	return &cli.Command{
		Name:  "show",
		Usage: "Show configured development domains",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("Configured domains: (none yet — storage not implemented)")
			return nil
		},
	}
}
