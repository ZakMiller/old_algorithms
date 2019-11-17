package replace

import "testing"

func TestReplace(t *testing.T) {
	tests := []struct {
		n, m, expected int
		i, j           uint
	}{
		{1 << 10, 19, 1100, 2, 6},
	}

	for _, test := range tests {
		got := replace(test.n, test.m, test.i, test.j)
		if got != test.expected {
			t.Errorf("replace(%v,%v,%v,%v) error got %v expected %v", test.n, test.m, test.i, test.j, got, test.expected)
		}
	}
}
