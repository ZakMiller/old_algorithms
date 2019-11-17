package lower

import "testing"

func TestToLower(t *testing.T) {
	tests := []struct {
		s, expected string
	}{
		{"ABC", "abc"},
		{"abc", "abc"},
		{"Abc", "abc"},
		{"013A", "013a"},
	}

	for _, test := range tests {
		got := toLower(test.s)
		if got != test.expected {
			t.Errorf("toLower(%v) failure got %v expected %v", test.s, got, test.expected)
		}
	}
}
