package caddy

import (
	"testing"
)

func TestGenerateCaddyfile(t *testing.T) {
	content, err := GenerateCaddyfile("dev.sekret.dev", 3000)
	if err != nil {
		t.Fatalf("GenerateCaddyfile failed: %v", err)
	}

	expected := `dev.sekret.dev

reverse_proxy 127.0.0.1:3000

tls \
 {$HOME}/.devhttps/certbot/config/live/dev.sekret.dev/fullchain.pem \
 {$HOME}/.devhttps/certbot/config/live/dev.sekret.dev/privkey.pem
`

	if content != expected {
		t.Errorf("Generated Caddyfile does not match expected.\nExpected:\n%s\nGot:\n%s", expected, content)
	}
}
