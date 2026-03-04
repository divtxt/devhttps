package cmd

import (
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/divtxt/devhttps/internal/certbot"
	"github.com/divtxt/devhttps/internal/config"
	"github.com/urfave/cli/v3"
)

var versionRegex = regexp.MustCompile(`v?(\d+)\.\d+`)

func newCheckCommand() *cli.Command {
	return &cli.Command{
		Name:  "check",
		Usage: "Run various checks (required tools, config file, certificates etc)",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("Tools:")
			ok := true
			ok = checkTool("certbot", []string{"--version"}, 5) && ok
			ok = checkTool("caddy", []string{"version"}, 2) && ok
			if !ok {
				return cli.Exit("", 1)
			}

			cfg, cfgErr := config.Load()
			certs, certsErr := certbot.Certificates()

			fmt.Println()
			fmt.Println("Config:")
			path, _ := config.Path()
			if cfgErr != nil {
				fmt.Printf("  ✗ %s\n", path)
				return cli.Exit("", 1)
			}
			fmt.Printf("  ✓ %s\n", path)

			fmt.Println()
			printConfiguredDomains(cfg, certs, certsErr != nil)
			printUnusedCerts(cfg, certs, certsErr)

			fmt.Println("To edit port or renew certificates, use: devhttps add")
			fmt.Println()
			return nil
		},
	}
}

func checkTool(name string, versionArgs []string, minMajor int) bool {
	toolPath, err := exec.LookPath(name)
	if err != nil {
		fmt.Printf("  ✗ %s: not found in PATH\n", name)
		return false
	}

	out, err := exec.Command(toolPath, versionArgs...).CombinedOutput()
	if err != nil {
		fmt.Printf("  ✗ %s: version check failed: %s\n", name, strings.TrimSpace(string(out)))
		return false
	}

	m := versionRegex.FindSubmatch(out)
	if m == nil {
		fmt.Printf("  ✗ %s: could not parse version from: %s\n", name, strings.TrimSpace(string(out)))
		return false
	}
	major, _ := strconv.Atoi(string(m[1]))
	if major < minMajor {
		fmt.Printf("  ✗ %s: version too old (v%d.x found, v%d+ required)\n", name, major, minMajor)
		return false
	}
	fmt.Printf("  ✓ %s: %s (%s)\n", name, strings.TrimSpace(string(out)), toolPath)
	return true
}

func printConfiguredDomains(cfg *config.Config, certs []certbot.CertInfo, certbotFailed bool) {
	fmt.Println("Configured domains:")
	if len(cfg.Entries) == 0 {
		fmt.Println("  (none)")
		fmt.Println()
		return
	}
	certMap := make(map[string]certbot.CertInfo)
	for _, c := range certs {
		certMap[c.Domain] = c
	}
	for _, e := range cfg.Entries {
		if certbotFailed {
			fmt.Printf("  ? %s → :%d  (cert: unknown)\n", e.Domain, e.Port)
		} else if cert, found := certMap[e.Domain]; found {
			if cert.Valid {
				fmt.Printf("  ✓ %s → :%d  (cert: VALID, %d days left)\n", e.Domain, e.Port, cert.DaysLeft)
			} else {
				fmt.Printf("  ✗ %s → :%d  (cert: INVALID)\n", e.Domain, e.Port)
			}
		} else {
			fmt.Printf("  ✗ %s → :%d  (cert: MISSING)\n", e.Domain, e.Port)
		}
	}
	fmt.Println()
}

func printUnusedCerts(cfg *config.Config, certs []certbot.CertInfo, certsErr error) {
	if certsErr != nil {
		fmt.Println()
		fmt.Println("Unused Certbot certificates:")
		fmt.Printf("  ✗ Failed to retrieve certificates: %v\n", certsErr)
		return
	}
	configDomains := make(map[string]bool)
	for _, e := range cfg.Entries {
		configDomains[e.Domain] = true
	}
	unused := []certbot.CertInfo{}
	for _, c := range certs {
		if !configDomains[c.Domain] {
			unused = append(unused, c)
		}
	}
	if len(unused) == 0 {
		return
	}
	fmt.Println()
	fmt.Println("Unused Certbot certificates:")
	for _, c := range unused {
		if c.Valid {
			fmt.Printf("  ✓ %s\n", c.Domain)
		} else {
			fmt.Printf("  ✗ %s\n", c.Domain)
		}
	}
}
