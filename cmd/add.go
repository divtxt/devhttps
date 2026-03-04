package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/divtxt/devhttps/internal/certbot"
	"github.com/divtxt/devhttps/internal/config"
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
				return cli.Exit("", 1)
			}

			domain := cmd.Args().Get(0)
			portStr := cmd.Args().Get(1)

			// Detect if stdin is a terminal (interactive vs non-interactive)
			isInteractive := false
			if fi, err := os.Stdin.Stat(); err == nil {
				isInteractive = (fi.Mode() & os.ModeCharDevice) != 0
			}
			if !isInteractive {
				fmt.Println("Error: 'add' must be run in an interactive terminal")
				return cli.Exit("", 1)
			}

			if err := validate.Domain(domain); err != nil {
				fmt.Printf("Error: %s\n", err)
				return cli.Exit("", 1)
			}
			fmt.Printf("✓ Domain: %s\n", domain)

			port, err := validate.Port(portStr)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return cli.Exit("", 1)
			}
			fmt.Printf("✓ Port: %d\n\n", port)

			cfg, _ := config.Load()
			cfg.Add(domain, port)
			if err := config.Save(cfg); err != nil {
				fmt.Printf("Error saving config: %s\n", err)
				return cli.Exit("", 1)
			}
			fmt.Printf("✓ Saved config entry: %s → port %d\n\n", domain, port)

			reader := bufio.NewReader(os.Stdin)

			// Check DNS resolution in a loop until it succeeds or user cancels
			for {
				fmt.Printf("→ Checking DNS resolution for '%s'...\n", domain)
				result, err := hostcheck.CheckResolvesToLocalhost(domain)
				if err != nil {
					fmt.Printf("x Error checking host resolution: %s\n", err)
					return cli.Exit("", 1)
				}

				if result.FoundInHostsFile || result.FoundViaDNS {
					// DNS is set up correctly
					fmt.Printf("✓ '%s' resolves to 127.0.0.1 (or ::1)\n\n", domain)
					break
				}

				// DNS not set up yet
				fmt.Printf("x '%s' does not resolve to localhost yet.\n\n", domain)
				fmt.Printf("→ Do ONE of the following:\n\n")
				fmt.Printf("  (A) Add this A record in your DNS provider:\n")
				fmt.Printf("      %s  →  127.0.0.1\n\n", domain)
				fmt.Printf("  (B) Add this line to your /etc/hosts file:\n")
				fmt.Printf("\n127.0.0.1   %s\n\n", domain)

				fmt.Printf("Press Enter when you've made the DNS entry (or Ctrl-C to quit): ")

				_, err = reader.ReadString('\n')
				if err != nil {
					fmt.Printf("\nError reading input\n")
					return cli.Exit("", 1)
				}
				fmt.Println()
			}

			// Check if a valid cert already exists
			certOK := false
			var daysLeft int
			certs, _ := certbot.Certificates()
			for _, c := range certs {
				if c.Domain == domain && c.Valid {
					certOK = true
					daysLeft = c.DaysLeft
					break
				}
			}

			if certOK {
				fmt.Printf("✓ %s cert already exists (%d days left)\n", domain, daysLeft)
			} else {
				fmt.Printf("→ Will run certbot to generate certificate for '%s'...\n", domain)
				fmt.Printf("\n    You will be prompted to do the following:\n")
				fmt.Printf("    - Enter your email address\n")
				fmt.Printf("    - Agree to the terms of service\n")
				fmt.Printf("    - Create a DNS TXT record in your DNS provider\n")
				fmt.Printf("\nPress Enter to run certbot (or Ctrl-C to quit): ")
				_, err = reader.ReadString('\n')
				if err != nil {
					fmt.Printf("\nError reading input\n")
					return cli.Exit("", 1)
				}
				fmt.Printf("\n------------ (RUNNING CERTBOT) ------------\n\n")
				if err := certbot.Run(domain); err != nil {
					fmt.Printf("Error running certbot: %s\n", err)
					return cli.Exit("", 1)
				}
				fmt.Printf("\n------------ (END CERTBOT) ------------\n")
				certs2, err2 := certbot.Certificates()
				if err2 != nil {
					fmt.Printf("x Error checking cert after certbot: %s\n", err2)
					return cli.Exit("", 1)
				}
				certNowOK := false
				var certDaysLeft int
				for _, c := range certs2 {
					if c.Domain == domain && c.Valid {
						certNowOK = true
						certDaysLeft = c.DaysLeft
						break
					}
				}
				if certNowOK {
					fmt.Printf("✓ Certificate for '%s' is valid (%d days left)\n", domain, certDaysLeft)
				} else {
					fmt.Printf("x Certificate for '%s' not found or invalid after certbot run\n", domain)
					return cli.Exit("", 1)
				}
			}

			fmt.Printf("\nYour service is now available at:\n")
			fmt.Printf("\n  https://%s\n\n", domain)
			return nil
		},
	}
}
