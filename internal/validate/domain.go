package validate

import (
	"fmt"
	"regexp"
)

var domainRegex = regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)

func Domain(s string) error {
	if len(s) > 253 {
		return fmt.Errorf("domain name too long (max 253 characters)")
	}
	if !domainRegex.MatchString(s) {
		return fmt.Errorf("invalid domain name %q (must be a fully-qualified domain like dev.example.com)", s)
	}
	return nil
}
