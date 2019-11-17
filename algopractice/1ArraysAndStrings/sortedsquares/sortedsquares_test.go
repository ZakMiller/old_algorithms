package sortedsquares

import (
	"testing"
)

func compare(one, two []int) bool {
	if len(one) != len(two) {
		return false
	}
	for i := 0; i < len(one); i++ {
		if one[i] != two[i] {
			return false
		}
	}
	return true
}
func TestSortedSquares(t *testing.T) {
	tests := []struct {
		input, expected []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}

	for _, test := range tests {
		got := sortedSquares(test.input)
		if !compare(got, test.expected) {
			t.Errorf("sortedSquares(%v) error got %v expected %v", test.input, got, test.expected)
		}
	}
}
