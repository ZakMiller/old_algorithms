package hamming

import "testing"

func TestHamming(t *testing.T) {
	tests := []struct {
		one, two, expected int
	}{
		{1, 4, 2},
		{7, 1, 2},
		{8, 0, 1},
		{11, 3, 1},
	}

	for _, test := range tests {
		got := hamming(test.one, test.two)
		if got != test.expected {
			t.Errorf("hamming(%v, %v) error got %v expected %v", test.one, test.two, got, test.expected)
		}
	}
}
