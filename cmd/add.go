package cmd

import (
	"context"
	"fmt"

	"github.com/divtxt/devhttps/internal/hostcheck"
	"github.com/divtxt/devhttps/internal/validate"
	"github.com/urfave/cli/v3"
)

func newAddCommand() *cli.Command {
	return &cli.Command{
		Name:      "add",
		Usage:     "Add a development domain proxied to a local port",
		ArgsUsage: "<dev-domain> <port>",
		Description: `Configure DevHttps to proxy HTTPS traffic for <dev-domain> to localhost:<port>.

You must have the ability to make DNS entries for this domain.

<dev-domain> must be a fully-qualified domain name (e.g. dev.myapp.com) that
resolves to 127.0.0.1 or ::1, via DNS.

<port> is the local port your dev server is listening on (1-65535).

Example:
  devhttps add dev.myapp.com 3000`,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.Args().Len() != 2 {
				fmt.Println(cmd.Description)
				fmt.Printf("\nUsage: devhttps %s %s\n", cmd.Name, cmd.ArgsUsage)
				return nil
			}

			domain := cmd.Args().Get(0)
			portStr := cmd.Args().Get(1)

			if err := validate.Domain(domain); err != nil {
				fmt.Printf("Error: %s\n", err)
				return nil
			}

			if _, err := validate.Port(portStr); err != nil {
				fmt.Printf("Error: %s\n", err)
				return nil
			}

			result, err := hostcheck.CheckResolvesToLocalhost(domain)
			if err != nil {
				fmt.Printf("Error checking host resolution: %s\n", err)
				return nil
			}

			if !result.FoundInHostsFile && !result.FoundViaDNS {
				fmt.Printf("'%s' does not resolve to localhost.\n", domain)
				fmt.Printf(`
To fix this, use ONE of the following options:

  Option A — Add a DNS A record (works across your network):
    In your DNS provider, add an A record:
      %s  →  127.0.0.1

  Option B — Edit /etc/hosts (quick, local-only):
    Add this line to /etc/hosts:
      127.0.0.1   %s


Once done, re-run:
  devhttps add %s %s
`, domain, domain, domain, portStr)
				return nil
			}

			fmt.Println("Added!")
			return nil
		},
	}
}
