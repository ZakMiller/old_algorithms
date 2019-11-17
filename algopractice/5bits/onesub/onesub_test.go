package onesub

import "testing"

func TestOneSub(t *testing.T) {
	tests := []struct {
		input, expected int
	}{
		{50, 3},
		{0, 1},
		{60, 5},
		{68, 2},
		{1, 2},
	}

	for _, test := range tests {
		got := oneSub(test.input)
		if got != test.expected {
			t.Errorf("oneSub(%v) error got %v wanted %v", test.input, got, test.expected)
		}
	}

}
