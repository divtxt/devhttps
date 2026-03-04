package cmd

import (
	"context"
	"fmt"

	"github.com/divtxt/devhttps/internal/certbot"
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
			certs, certsErr := certbot.Certificates()
			printConfiguredDomains(cfg, certs, certsErr != nil)
			fmt.Println("(for more details, use check command)")
			return nil
		},
	}
}
