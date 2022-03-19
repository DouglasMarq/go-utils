package goutils

import (
	"strings"
)

func IsEmptyOrNullString(s string) bool {
	sPointer := &s

	if sPointer == nil || strings.TrimSpace(s) == "" {
		return true
	}

	return false
}
