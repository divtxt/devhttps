package certbot

import (
	"testing"
)

func TestParseCertificates(t *testing.T) {
	// Sample certbot output with multiple certificates
	output := `Found the following certs:
  Certificate Name: dev.slackduty.com
    Domains: dev.slackduty.com
    Expiry Date: 2024-10-02 16:56:11+00:00 (INVALID: EXPIRED)
    Certificate Path: /Users/div/.devhttps/certbot/config/live/dev.slackduty.com/fullchain.pem
    Private Key Path: /Users/div/.devhttps/certbot/config/live/dev.slackduty.com/privkey.pem
  Certificate Name: dev.txtlabs.com
    Domains: dev.txtlabs.com
    Expiry Date: 2023-05-07 08:36:02+00:00 (INVALID: EXPIRED)
    Certificate Path: /Users/div/.devhttps/certbot/config/live/dev.txtlabs.com/fullchain.pem
    Private Key Path: /Users/div/.devhttps/certbot/config/live/dev.txtlabs.com/privkey.pem
  Certificate Name: dev.divtxt.com
    Domains: dev.divtxt.com
    Expiry Date: 2026-06-02 04:53:04+00:00 (VALID: 89 days)
    Certificate Path: /Users/div/.devhttps/certbot/config/live/dev.divtxt.com/fullchain.pem
    Private Key Path: /Users/div/.devhttps/certbot/config/live/dev.divtxt.com/privkey.pem`

	certs := parseCertificates(output)

	// Verify count
	if len(certs) != 3 {
		t.Fatalf("expected 3 certificates, got %d", len(certs))
	}

	tests := []struct {
		name     string
		idx      int
		domain   string
		valid    bool
		daysLeft int
	}{
		{
			name:     "dev.slackduty.com",
			idx:      0,
			domain:   "dev.slackduty.com",
			valid:    false,
			daysLeft: 0,
		},
		{
			name:     "dev.txtlabs.com",
			idx:      1,
			domain:   "dev.txtlabs.com",
			valid:    false,
			daysLeft: 0,
		},
		{
			name:     "dev.divtxt.com",
			idx:      2,
			domain:   "dev.divtxt.com",
			valid:    true,
			daysLeft: 89,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cert := certs[tt.idx]

			if cert.Domain != tt.domain {
				t.Errorf("domain: expected %q, got %q", tt.domain, cert.Domain)
			}
			if cert.Valid != tt.valid {
				t.Errorf("valid: expected %v, got %v", tt.valid, cert.Valid)
			}
			if cert.DaysLeft != tt.daysLeft {
				t.Errorf("daysLeft: expected %d, got %d", tt.daysLeft, cert.DaysLeft)
			}
		})
	}
}

func TestParseCertificatesEmpty(t *testing.T) {
	output := ""

	certs := parseCertificates(output)

	// Empty output should return empty slice, not nil
	if certs == nil {
		t.Fatal("expected empty slice, got nil")
	}
	if len(certs) != 0 {
		t.Fatalf("expected 0 certificates, got %d", len(certs))
	}
}
