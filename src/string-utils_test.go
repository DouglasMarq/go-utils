package goutils

import (
	"testing"
)

func TestCorrectString(t *testing.T) {
	s := "notEmptyOrNull"

	result := IsEmptyOrNullString(s)

	if result != false {
		t.Errorf("Expected string to not be empty or null, but got %v", result)
	}
}

func TestCorrectStringWithWhiteSpace(t *testing.T) {
	s := " not Empty Or Null"

	result := IsEmptyOrNullString(s)

	if result != false {
		t.Errorf("Expected string to not be empty or null, but got %v", result)
	}
}

func TestWhiteSpaceString(t *testing.T) {
	s := " "

	result := IsEmptyOrNullString(s)

	if result != true {
		t.Errorf("Expected string to be empty or null, but got %v", result)
	}
}

func TestBlankString(t *testing.T) {
	s := ""

	result := IsEmptyOrNullString(s)

	if result != true {
		t.Errorf("Expected string to be empty or null, but got %v", result)
	}
}
