package certbot

import (
	"os"
	"os/exec"
	"path/filepath"
)

func dirs() (configDir, logsDir, workDir string, err error) {
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
	configDir, logsDir, workDir, err := dirs()
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
