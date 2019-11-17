package other

import "testing"

func TestSumExists(t *testing.T) {
	tests := []struct {
		nums   []int
		sum    int
		wanted bool
	}{
		{[]int{1, 2, 3, 9}, 8, false},
		{[]int{1, 2, 4, 4}, 8, true},
	}

	for _, test := range tests {
		got := sumExists(test.nums, test.sum)
		if got != test.wanted {
			t.Errorf("sumExists(%v, %v) failure got %v wanted %v", test.nums, test.sum, got, test.wanted)
		}
	}
}
