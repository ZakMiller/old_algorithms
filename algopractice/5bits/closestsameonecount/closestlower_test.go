package closesthigher

import "testing"

func TestClosestLower(t *testing.T) {
	tests := []struct {
		i, o int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 3},
		{6, 5},
		{5, 3},
	}

	for _, test := range tests {
		got := closestLower(test.i)
		if got != test.o {
			t.Errorf("closestLower(%v) error got %v expected %v", test.i, got, test.o)
		}
	}
}
