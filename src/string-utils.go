package goutils

import (
	"strings"
)

func IsEmptyOrNullString(s string) bool {
	if strings.TrimSpace(s) == "" {
		return true
	}

	return false
}
