package oneaway

import "testing"

func TestOneAway(t *testing.T) {
	tests := []struct{
		first string
		second string
		expected bool
	}{
		{"one","oneb", true},
		{"one", "ne",  true},
		{"one", "onr", true},
		{"one", "onebr", false},
		{"one", "xyz", false},
		{"one", "o", false},
	}

	for _, test := range tests {
		actual := OneAway(test.first, test.second)
		if actual != test.expected {
			t.Errorf("OneAway(%v, %v) failure actual %v expected %v", test.first, test.second, actual, test.expected)
		}
	}
}
