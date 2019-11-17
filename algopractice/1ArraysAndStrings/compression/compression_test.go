package compression

import "testing"

func TestCompression(t *testing.T) {
	tests := []struct{
		s string
		expected string
	}{
		{"abcde", "abcde"},
		{"aabb", "aabb"},
		{"aaabbb", "a3b3"},
		{"rrrreea", "r4e2a1"},
	}

	for _, test := range tests {
		actual := Compress(test.s)
		if actual != test.expected {
			t.Errorf("Compress(%v) failure actual %v expected %v", test.s, actual, test.expected)
		}
	}
}