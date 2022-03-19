package goutils

import (
	"strings"
)

func CheckForEmptyOrNullString(s string) bool {
	if strings.TrimSpace(s) == "" {
		return true
	}

	return false
}
