package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/divtxt/devhttps/internal/caddy"
	"github.com/divtxt/devhttps/internal/certbot"
	"github.com/divtxt/devhttps/internal/config"
	"github.com/urfave/cli/v3"
)

func newRunCommand() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "Run Caddy with the configured domains",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// Load config and certs
			cfg, cfgErr := config.Load()
			certs, certsErr := certbot.Certificates()

			if cfgErr != nil {
				fmt.Printf("Error loading config: %v\n", cfgErr)
				return cli.Exit("", 1)
			}

			// Show configured domains
			printConfiguredDomains(cfg, certs, certsErr != nil)

			// Generate Caddyfile
			fmt.Printf("Generating Caddyfile...\n")
			content, err := caddy.GenerateCaddyfile(cfg)
			if err != nil {
				fmt.Printf("Error generating Caddyfile: %s\n", err)
				return cli.Exit("", 1)
			}

			// Write Caddyfile to disk
			if err := caddy.WriteCaddyfile(content); err != nil {
				fmt.Printf("Error writing Caddyfile: %s\n", err)
				return cli.Exit("", 1)
			}

			fmt.Printf("\nStarting Caddy...\n\n")

			caddyDir, err := caddy.Dir()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return cli.Exit("", 1)
			}

			c := exec.Command("caddy", "run")
			c.Dir = caddyDir
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			return c.Run()
		},
	}
}
