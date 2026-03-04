package validate

import (
	"fmt"
	"strconv"
)

func Port(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid port %q (must be an integer)", s)
	}
	if n < 1 || n > 65535 {
		return 0, fmt.Errorf("invalid port %d (must be between 1 and 65535)", n)
	}
	return n, nil
}
