package caddy

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/divtxt/devhttps/internal/certbot"
	"github.com/divtxt/devhttps/internal/config"
)

func TestGenerateCaddyfile(t *testing.T) {
	cfg := &config.Config{
		Entries: []config.Entry{
			{Domain: "dev.sekret.dev", Port: 3000},
			{Domain: "dev.divtxt.com", Port: 4000},
		},
	}

	content, err := GenerateCaddyfile(cfg)
	if err != nil {
		t.Fatalf("GenerateCaddyfile failed: %v", err)
	}

	// Build expected string
	certConfigDir, _, _, _ := certbot.Dirs()
	expected := fmt.Sprintf(`dev.sekret.dev {
	reverse_proxy 127.0.0.1:3000
	tls %s %s
}

dev.divtxt.com {
	reverse_proxy 127.0.0.1:4000
	tls %s %s
}

`,
		filepath.Join(certConfigDir, "live", "dev.sekret.dev", "fullchain.pem"),
		filepath.Join(certConfigDir, "live", "dev.sekret.dev", "privkey.pem"),
		filepath.Join(certConfigDir, "live", "dev.divtxt.com", "fullchain.pem"),
		filepath.Join(certConfigDir, "live", "dev.divtxt.com", "privkey.pem"),
	)

	if content != expected {
		t.Errorf("Generated Caddyfile does not match expected.\nExpected:\n%s\nGot:\n%s", expected, content)
	}
}
