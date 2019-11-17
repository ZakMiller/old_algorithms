package matrixrotation

import "testing"

func TestRotate(t *testing.T) {
	tests := []struct {
		initial [][]uint32
		rotated [][]uint32
	}{
		{
			[][]uint32{{1, 2, 3}, {3, 4, 5}, {4, 5, 6}},
			[][]uint32{{4, 3, 1}, {5, 4, 2}, {6, 5, 3}},
		},
		{
			[][]uint32{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}},
			[][]uint32{{21, 16, 11, 6, 1}, {22, 17, 12, 7, 2}, {23, 18, 13, 8, 3}, {24, 19, 14, 9, 4}, {25, 20, 15, 10, 5}},
		},
	}

	for _, test := range tests {
		actual := Rotate(test.initial)
		if isSame(actual, test.rotated) == false {
			t.Errorf("Rotated(%v) failure actual %v expected %v", test.initial, actual, test.rotated)
		}
	}
}

func TestRotateInPlace(t *testing.T) {
	tests := []struct {
		initial [][]uint32
		rotated [][]uint32
	}{
		{
			[][]uint32{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]uint32{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			[][]uint32{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}},
			[][]uint32{{21, 16, 11, 6, 1}, {22, 17, 12, 7, 2}, {23, 18, 13, 8, 3}, {24, 19, 14, 9, 4}, {25, 20, 15, 10, 5}},
		},
	}

	for _, test := range tests {
		RotateInPlace(test.initial)
		if isSame(test.initial, test.rotated) == false {
			t.Errorf("Rotated() failure actual %v expected %v", test.initial, test.rotated)
		}
	}
}

func isSame(one [][]uint32, two [][]uint32) bool {
	for i := range one {
		for j := range one[i] {
			if one[i][j] != two[i][j] {
				return false
			}
		}
	}
	return true
}
