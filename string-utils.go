package goutils

import (
	"reflect"
	"strings"
)

type stringTypes struct {
	s string
	s []string
}

func CheckForEmptyOrNullString(s string) {
	if strings.TrimSpace(s) == "" {
		return true
	}

	return false
}
