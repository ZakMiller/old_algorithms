package permutationpalindrome

import "testing"

func TestIsPermutationPalindrome(t *testing.T) {
	tests := []struct{
		s string
		expected bool
	}{
		{"abba",  true},
		{"abab",  true},
		{"abcba",  true},
		{"cabab",  true},
		{"abcd",  false},
		{"abcde",  false},
		{"abcdeabcdf", false},
	}

	for _, test := range tests {
		actual := IsPermutationPalindrome(test.s)
		if actual != test.expected {
			t.Errorf("IsPermutationPalindrome(%v) actual %v want %v", test.s, actual, test.expected)
		}
	}
}
