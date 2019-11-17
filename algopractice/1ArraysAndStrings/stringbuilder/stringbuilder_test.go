package stringbuilder

import (
	"fmt"
	"testing"
)

func TestTypicalCase(t *testing.T) {
	sb := &StringBuilder{}
	sb.AddString("hello ")
	sb.AddString("there my friend!")
	fmt.Println()
	if sb.ToString() != "hello there my friend!" {
		t.Errorf("Two strings not added correctly.")
	}
}

func TestNothingAdded(t *testing.T) {
	sb := &StringBuilder{}
	fmt.Println()
	if sb.ToString() != "" {
		t.Errorf("Zeroeth case incorrect.")
	}
}
