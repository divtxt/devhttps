package hostcheck

import (
	"bufio"
	"net"
	"os"
	"strings"
)

type Result struct {
	FoundInHostsFile bool
	FoundViaDNS      bool
}

func CheckResolvesToLocalhost(domain string) (Result, error) {
	// Check /etc/hosts first
	f, err := os.Open("/etc/hosts")
	if err != nil {
		return Result{}, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// Strip comments
		if idx := strings.Index(line, "#"); idx >= 0 {
			line = line[:idx]
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		ip := fields[0]
		if ip != "127.0.0.1" && ip != "::1" {
			continue
		}
		for _, hostname := range fields[1:] {
			if strings.EqualFold(hostname, domain) {
				return Result{FoundInHostsFile: true}, nil
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return Result{}, err
	}

	// DNS fallback
	addrs, err := net.LookupHost(domain)
	if err != nil {
		// NXDOMAIN or any lookup error — not an app error
		return Result{}, nil
	}
	for _, addr := range addrs {
		if addr == "127.0.0.1" || addr == "::1" {
			return Result{FoundViaDNS: true}, nil
		}
	}

	return Result{}, nil
}
