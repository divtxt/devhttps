package caddy

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Dir returns the Caddy config directory: ~/.devhttps/caddy
func Dir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".devhttps", "caddy"), nil
}

// GenerateCaddyfile generates Caddyfile content for a single domain+port and returns it as a string
func GenerateCaddyfile(domain string, port int) (string, error) {
	fullchainPath := fmt.Sprintf("{$HOME}/.devhttps/certbot/config/live/%s/fullchain.pem", domain)
	privkeyPath := fmt.Sprintf("{$HOME}/.devhttps/certbot/config/live/%s/privkey.pem", domain)

	var buf strings.Builder
	buf.WriteString(domain + "\n\n")
	buf.WriteString(fmt.Sprintf("reverse_proxy 127.0.0.1:%d\n\n", port))
	buf.WriteString(fmt.Sprintf("tls \\\n %s \\\n %s\n", fullchainPath, privkeyPath))

	return buf.String(), nil
}

// WriteCaddyfile writes content to ~/.devhttps/caddy/Caddyfile.<domain>.<port> and returns the path
func WriteCaddyfile(domain string, port int, content string) (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", err
	}

	caddyfilePath := filepath.Join(dir, fmt.Sprintf("Caddyfile.%s.%d", domain, port))
	if err := os.WriteFile(caddyfilePath, []byte(content), 0600); err != nil {
		return "", err
	}

	return caddyfilePath, nil
}
