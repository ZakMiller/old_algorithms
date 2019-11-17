package stringrotation

import "testing"

func TestIsRotatedString(t *testing.T) {
	tests := []struct {
		one      string
		two      string
		expected bool
	}{
		{"waterbottle", "erbottlewat", true},
		{"one", "two", false},
		{"abcdefg", "gabcdef", true},
		{"abadafag", "adabagaf", false},
	}

	for _, test := range tests {
		actual := IsRotatedString(test.one, test.two)
		if actual != test.expected {
			t.Errorf("IsRotatedString(%v, %v) error actual %v expected %v", test.one, test.two, actual, test.expected)
		}
	}
}
