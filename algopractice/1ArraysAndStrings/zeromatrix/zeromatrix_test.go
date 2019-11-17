package zeromatrix

import "testing"

func isSame(one [][]int, two [][]int) bool {
	for i := range one {
		for j := range one[i] {
			if one[i][j] != two[i][j] {
				return false
			}
		}
	}
	return true
}

func TestZeroOut(t *testing.T) {
	tests := []struct {
		initial  [][]int
		expected [][]int
	}{
		{
			[][]int{{1, 0, 2, 3, 4}, {5, 6, 7, 8, 9}, {10, 11, 12, 13, 0}},
			[][]int{{0, 0, 0, 0, 0}, {5, 0, 7, 8, 0}, {0, 0, 0, 0, 0}},
		},
	}
	for _, test := range tests {
		ZeroOut(test.initial)
		if !isSame(test.initial, test.expected) {
			t.Errorf("ZeroOut failure actual %v expected %v", test.initial, test.expected)
		}
	}
}
