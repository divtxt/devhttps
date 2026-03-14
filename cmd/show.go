package cmd

import (
	"context"
	"fmt"

	"github.com/divtxt/devhttps/internal/certbot"
	"github.com/urfave/cli/v3"
)

func newShowCommand() *cli.Command {
	return &cli.Command{
		Name:  "show",
		Usage: "Show available certificates",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			certs, err := certbot.Certificates()
			if err != nil {
				fmt.Printf("Error listing certificates: %s\n", err)
				return cli.Exit("", 1)
			}
			if len(certs) == 0 {
				fmt.Println("No certificates found (use add to add one)")
				return nil
			}
			for _, c := range certs {
				if c.Valid {
					fmt.Printf("  ✓ %s (%d days left)\n", c.Domain, c.DaysLeft)
				} else {
					fmt.Printf("  ✗ %s (INVALID)\n", c.Domain)
				}
			}
			return nil
		},
	}
}
