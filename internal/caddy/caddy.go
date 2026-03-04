package caddy

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/divtxt/devhttps/internal/certbot"
	"github.com/divtxt/devhttps/internal/config"
)

// Dir returns the Caddy config directory: ~/.devhttps/caddy
func Dir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".devhttps", "caddy"), nil
}

// GenerateCaddyfile generates Caddyfile content from all config entries and returns it as a string
func GenerateCaddyfile(cfg *config.Config) (string, error) {
	// Get certbot config directory for cert paths
	certConfigDir, _, _, err := certbot.Dirs()
	if err != nil {
		return "", err
	}

	// Generate Caddyfile content
	var buf strings.Builder
	for _, entry := range cfg.Entries {
		certPath := filepath.Join(certConfigDir, "live", entry.Domain)
		fullchainPath := filepath.Join(certPath, "fullchain.pem")
		privkeyPath := filepath.Join(certPath, "privkey.pem")

		buf.WriteString(entry.Domain + " {\n")
		buf.WriteString(fmt.Sprintf("\treverse_proxy 127.0.0.1:%d\n", entry.Port))
		buf.WriteString(fmt.Sprintf("\ttls %s %s\n", fullchainPath, privkeyPath))
		buf.WriteString("}\n\n")
	}

	return buf.String(), nil
}

// WriteCaddyfile writes content to the Caddyfile at ~/.devhttps/caddy/Caddyfile
func WriteCaddyfile(content string) error {
	dir, err := Dir()
	if err != nil {
		return err
	}

	// Create directory if it doesn't exist
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	// Write Caddyfile
	caddyfilePath := filepath.Join(dir, "Caddyfile")
	if err := os.WriteFile(caddyfilePath, []byte(content), 0600); err != nil {
		return err
	}

	return nil
}

// ReadCaddyfile reads the Caddyfile from disk and returns its content as a string
func ReadCaddyfile() (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(filepath.Join(dir, "Caddyfile"))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Validate pipes content to `caddy validate --config -` via stdin and returns any error
func Validate(content string) error {
	dir, err := Dir()
	if err != nil {
		return err
	}
	cmd := exec.Command("caddy", "validate")
	cmd.Dir = dir
	cmd.Stdin = strings.NewReader(content)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %w", strings.TrimSpace(string(output)), err)
	}
	return nil
}
