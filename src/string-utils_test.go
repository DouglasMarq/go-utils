package goutils

import (
	"testing"
)

func TestCorrectString(t *testing.T) {
	s := "notEmptyOrNull"

	result := IsEmptyOrNullString(s)

	if result != false {
		t.Errorf("String is Empty or Null %v", result)
	}
}

func TestWhitespaceString(t *testing.T) {
	s := " "

	result := IsEmptyOrNullString(s)

	if result != true {
		t.Errorf("String is Empty or Null %v", result)
	}
}

func TestBlankString(t *testing.T) {
	s := ""

	result := IsEmptyOrNullString(s)

	if result != true {
		t.Errorf("String is Empty or Null %v", result)
	}
}
