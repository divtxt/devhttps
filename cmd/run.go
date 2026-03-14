package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/divtxt/devhttps/internal/caddy"
	"github.com/divtxt/devhttps/internal/certbot"
	"github.com/divtxt/devhttps/internal/validate"
	"github.com/urfave/cli/v3"
)

func newRunCommand() *cli.Command {
	return &cli.Command{
		Name:      "run",
		Usage:     "Run https server for a domain, and proxy to a http port",
		ArgsUsage: "<domain> <port>",
		Description: `Run Caddy as https proxy for <domain> and forward to http://127.0.0.1:<port>

A valid certificate for <domain> must exist (use add command).

Example:
  devhttps run dev.example.com 3000`,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.Args().Len() != 2 {
				fmt.Println(cmd.Description)
				fmt.Printf("\nUsage: devhttps %s %s\n", cmd.Name, cmd.ArgsUsage)
				return cli.Exit("", 1)
			}

			domain := cmd.Args().Get(0)
			portStr := cmd.Args().Get(1)

			if err := validate.Domain(domain); err != nil {
				fmt.Printf("Error: %s\n", err)
				return cli.Exit("", 1)
			}

			port, err := validate.Port(portStr)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return cli.Exit("", 1)
			}

			// Validate cert
			fmt.Printf("Checking certificate for '%s'...\n", domain)
			certs, err := certbot.Certificates()
			if err != nil {
				fmt.Printf("Error checking certificates: %s\n", err)
				return cli.Exit("", 1)
			}
			certValid := false
			for _, c := range certs {
				if c.Domain == domain && c.Valid {
					certValid = true
					fmt.Printf("✓ Certificate valid (%d days left)\n", c.DaysLeft)
					break
				}
			}
			if !certValid {
				fmt.Printf("Error: no valid certificate for '%s' (use add first)\n", domain)
				return cli.Exit("", 1)
			}

			// Generate Caddyfile
			content, err := caddy.GenerateCaddyfile(domain, port)
			if err != nil {
				fmt.Printf("Error generating Caddyfile: %s\n", err)
				return cli.Exit("", 1)
			}

			// Write Caddyfile
			caddyfilePath, err := caddy.WriteCaddyfile(domain, port, content)
			if err != nil {
				fmt.Printf("Error writing Caddyfile: %s\n", err)
				return cli.Exit("", 1)
			}
			fmt.Printf("✓ Caddyfile written: %s\n", caddyfilePath)


			fmt.Printf("\nStarting Caddy for https://%s → 127.0.0.1:%d\n\n", domain, port)

			c := exec.Command("caddy", "run", "--config", caddyfilePath)
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			return c.Run()
		},
	}
}
