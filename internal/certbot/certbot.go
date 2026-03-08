package certbot

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func Dirs() (configDir, logsDir, workDir string, err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}
	base := filepath.Join(home, ".devhttps", "certbot")
	configDir = filepath.Join(base, "config")
	logsDir = filepath.Join(base, "logs")
	workDir = filepath.Join(base, "work")
	return
}

func Run(domain string) error {
	configDir, logsDir, workDir, err := Dirs()
	if err != nil {
		return err
	}
	for _, dir := range []string{configDir, logsDir, workDir} {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return err
		}
	}
	cmd := exec.Command("certbot",
		"certonly",
		"--config-dir", configDir,
		"--logs-dir", logsDir,
		"--work-dir", workDir,
		"--manual",
		"--preferred-challenges", "dns",
		"-d", domain,
	)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

type CertInfo struct {
	Domain   string
	Valid    bool
	DaysLeft int // only meaningful when Valid=true
}

func Certificates() ([]CertInfo, error) {
	configDir, logsDir, workDir, err := Dirs()
	if err != nil {
		return nil, err
	}
	out, err := exec.Command("certbot", "certificates",
		"--config-dir", configDir,
		"--logs-dir", logsDir,
		"--work-dir", workDir,
	).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", strings.TrimSpace(string(out)), err)
	}
	return parseCertificates(string(out)), nil
}

func CertificatesVerbose() ([]CertInfo, error) {
	configDir, logsDir, workDir, err := Dirs()
	if err != nil {
		return nil, err
	}
	fmt.Printf("\n------------ (RUNNING CERTBOT CERTIFICATES) ------------\n\n")
	var buf bytes.Buffer
	cmd := exec.Command("certbot", "certificates",
		"--config-dir", configDir,
		"--logs-dir", logsDir,
		"--work-dir", workDir,
	)
	cmd.Stdout = io.MultiWriter(os.Stdout, &buf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &buf)
	err = cmd.Run()
	fmt.Printf("\n------------ (END CERTBOT CERTIFICATES) ------------\n")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", strings.TrimSpace(buf.String()), err)
	}
	return parseCertificates(buf.String()), nil
}

func parseCertificates(output string) []CertInfo {
	certRegex := regexp.MustCompile(`Certificate Name:\s*(\S+)`)
	statusRegex := regexp.MustCompile(`Expiry Date:[^(]*\((VALID|INVALID)(?::\s*(\d+)\s*days)?`)

	lines := strings.Split(output, "\n")
	certs := make([]CertInfo, 0)
	var currentDomain string

	for _, line := range lines {
		certMatch := certRegex.FindStringSubmatch(line)
		if certMatch != nil {
			currentDomain = certMatch[1]
			continue
		}

		statusMatch := statusRegex.FindStringSubmatch(line)
		if statusMatch != nil && currentDomain != "" {
			validStr := statusMatch[1]
			valid := validStr == "VALID"
			daysLeft := 0

			if valid && statusMatch[2] != "" {
				daysLeft, _ = strconv.Atoi(statusMatch[2])
			}

			certs = append(certs, CertInfo{
				Domain:   currentDomain,
				Valid:    valid,
				DaysLeft: daysLeft,
			})

			currentDomain = ""
		}
	}

	return certs
}
