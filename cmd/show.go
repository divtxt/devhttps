package cmd

import (
	"context"
	"fmt"

	"github.com/divtxt/devhttps/internal/config"
	"github.com/urfave/cli/v3"
)

func newShowCommand() *cli.Command {
	return &cli.Command{
		Name:  "show",
		Usage: "Show configured development domains",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			cfg, err := config.Load()
			if err != nil {
				path, _ := config.Path()
				fmt.Printf("Error loading config (%s): %s\n", path, err)
				return cli.Exit("", 1)
			}
			if len(cfg.Entries) == 0 {
				fmt.Println("Configured domains: (none)")
				return nil
			}
			for _, e := range cfg.Entries {
				fmt.Printf("https://%s/ → http://localhost:%d/\n", e.Domain, e.Port)
			}
			return nil
		},
	}
}
