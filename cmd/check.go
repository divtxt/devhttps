package cmd

import (
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/divtxt/devhttps/internal/certbot"
	"github.com/urfave/cli/v3"
)

var versionRegex = regexp.MustCompile(`v?(\d+)\.\d+`)

func newCheckCommand() *cli.Command {
	return &cli.Command{
		Name:  "check",
		Usage: "Check that required tools are installed and meet minimum versions",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			ok := true
			ok = checkTool("certbot", []string{"--version"}, 5) && ok
			ok = checkTool("caddy", []string{"version"}, 2) && ok
			if !ok {
				return cli.Exit("", 1)
			}
			fmt.Println()
			checkCertbotCerts()
			return nil
		},
	}
}

func checkTool(name string, versionArgs []string, minMajor int) bool {
	toolPath, err := exec.LookPath(name)
	if err != nil {
		fmt.Printf("✗ %s: not found in PATH\n", name)
		return false
	}

	out, err := exec.Command(toolPath, versionArgs...).CombinedOutput()
	if err != nil {
		fmt.Printf("✗ %s: version check failed: %s\n", name, strings.TrimSpace(string(out)))
		return false
	}

	m := versionRegex.FindSubmatch(out)
	if m == nil {
		fmt.Printf("✗ %s: could not parse version from: %s\n", name, strings.TrimSpace(string(out)))
		return false
	}
	major, _ := strconv.Atoi(string(m[1]))
	if major < minMajor {
		fmt.Printf("✗ %s: version too old (v%d.x found, v%d+ required)\n", name, major, minMajor)
		return false
	}
	fmt.Printf("✓ %s: %s (%s)\n", name, strings.TrimSpace(string(out)), toolPath)
	return true
}

func checkCertbotCerts() {
	fmt.Println("Certbot certificates:")
	certs, err := certbot.Certificates()
	if err != nil {
		fmt.Printf("  ✗ Failed to retrieve certificates: %v\n", err)
		return
	}
	if len(certs) == 0 {
		fmt.Println("  (no certificates found)")
		return
	}
	for _, c := range certs {
		if c.Valid {
			fmt.Printf("  ✓ %s  (VALID: %d days)\n", c.Domain, c.DaysLeft)
		} else {
			fmt.Printf("  ✗ %s  (INVALID: EXPIRED)\n", c.Domain)
		}
	}
}
