package urlify

import "testing"

func TestUrlifySpaces(t *testing.T) {
	if Urlify("a b c") != "a%20b%20c" {
		t.Error("spaces not replaced correctly")
	}
}

func TestUrlifyNoSpaces(t *testing.T) {
	if Urlify("abc") != "abc" {
		t.Error("something went wrong with string that didn't have spaces")
	}
}

func TestUrlifyNothing(t *testing.T) {
	if Urlify("") != "" {
		t.Error("something went wrong with empty string")
	}
}
