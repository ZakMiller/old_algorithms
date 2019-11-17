package closesthigher

import "testing"

func TestClosestHigher(t *testing.T) {
	tests := []struct {
		i, o int
	}{
		{0, 0},
		{1, 2},
		{2, 4},
		{3, 5},
		{7, 11},
		{5, 6},
	}

	for _, test := range tests {
		got := closestHigher(test.i)
		if got != test.o {
			t.Errorf("closestHigher(%v) error got %v expected %v", test.i, got, test.o)
		}
	}
}
