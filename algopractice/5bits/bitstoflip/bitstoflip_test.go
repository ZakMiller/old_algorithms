package bitstoflip

import "testing"

func TestBitsToFlip(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{3, 7, 1},
		{0, 0, 0},
		{10, 10, 0},
		{7, 5, 1},
		{15, 0, 4},
	}

	for _, test := range tests {
		got := bitsToFlip(test.a, test.b)
		if got != test.expected {
			t.Errorf("bitsToFlip(%v, %v) error got %v expected %v", test.a, test.b, got, test.expected)
		}
	}
}
