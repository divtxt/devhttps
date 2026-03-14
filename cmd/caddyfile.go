package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/divtxt/devhttps/internal/caddy"
	"github.com/divtxt/devhttps/internal/certbot"
	"github.com/divtxt/devhttps/internal/validate"
	"github.com/urfave/cli/v3"
)

func newCaddyfileCommand() *cli.Command {
	return &cli.Command{
		Name:      "caddyfile",
		Usage:     "Print a Caddyfile for a domain and port to stdout",
		ArgsUsage: "<domain> <port>",
		Description: `Print a Caddyfile for <domain> and <port> to stdout

A valid certificate for <domain> must exist (use add command).

Example:
  devhttps caddyfile dev.example.com 3000 > ./Caddyfile`,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.Args().Len() != 2 {
				fmt.Fprintln(os.Stderr, cmd.Description)
				fmt.Fprintf(os.Stderr, "\nUsage: devhttps %s %s\n", cmd.Name, cmd.ArgsUsage)
				return cli.Exit("", 1)
			}

			domain := cmd.Args().Get(0)
			portStr := cmd.Args().Get(1)

			if err := validate.Domain(domain); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				return cli.Exit("", 1)
			}

			port, err := validate.Port(portStr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				return cli.Exit("", 1)
			}

			certs, err := certbot.Certificates()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error checking certificates: %s\n", err)
				return cli.Exit("", 1)
			}
			certValid := false
			for _, c := range certs {
				if c.Domain == domain && c.Valid {
					certValid = true
					break
				}
			}
			if !certValid {
				fmt.Fprintf(os.Stderr, "Error: no valid certificate for '%s' (use add first)\n", domain)
				return cli.Exit("", 1)
			}

			content, err := caddy.GenerateCaddyfile(domain, port)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error generating Caddyfile: %s\n", err)
				return cli.Exit("", 1)
			}

			fmt.Print(content)
			return nil
		},
	}
}
