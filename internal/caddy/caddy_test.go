package caddy

import (
	"testing"

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

	expected := `dev.sekret.dev {
	reverse_proxy 127.0.0.1:3000
	tls \
	 {$HOME}/.devhttps/certbot/config/live/dev.sekret.dev/fullchain.pem \
	 {$HOME}/.devhttps/certbot/config/live/dev.sekret.dev/privkey.pem
}

dev.divtxt.com {
	reverse_proxy 127.0.0.1:4000
	tls \
	 {$HOME}/.devhttps/certbot/config/live/dev.divtxt.com/fullchain.pem \
	 {$HOME}/.devhttps/certbot/config/live/dev.divtxt.com/privkey.pem
}

`

	if content != expected {
		t.Errorf("Generated Caddyfile does not match expected.\nExpected:\n%s\nGot:\n%s", expected, content)
	}
}
