package twoDee

import "testing"

func TestCanCut(t *testing.T) {
	tests := []struct {
		forest [][]int
		wanted bool
	}{
		{[][]int{[]int{1, 2, 3}, []int{0, 0, 4}, []int{7, 6, 5}}, true},
	}
	_ = tests
}

func TestTreeCount(t *testing.T) {
	tests := []struct {
		forest [][]int
		wanted int
	}{
		{[][]int{[]int{1, 2, 3}, []int{0, 0, 4}, []int{7, 6, 5}}, 6},
		{[][]int{[]int{1, 2, 3}, []int{0, 0, 0}, []int{7, 6, 5}}, 5},
	}
	for _, test := range tests {
		got := treeCount(test.forest)
		if got != test.wanted {
			t.Errorf("treeCount(%v) error got %v wanted %v", test.forest, got, test.wanted)
		}
	}
}

func TestReachableTreeCount(t *testing.T) {
	tests := []struct {
		forest [][]int
		wanted int
	}{
		{[][]int{[]int{1, 2, 3}, []int{0, 0, 4}, []int{7, 6, 5}}, 6},
		{[][]int{[]int{1, 2, 3}, []int{0, 0, 0}, []int{7, 6, 5}}, 2},
		{[][]int{[]int{2, 3, 4}, []int{0, 0, 5}, []int{8, 6, 7}}, 7},
	}
	for _, test := range tests {
		got := reachableTreeCount(test.forest)
		if got != test.wanted {
			t.Errorf("treeCount(%v) error got %v wanted %v", test.forest, got, test.wanted)
		}
	}
}

func TestCanReachAllTrees(t *testing.T) {
	tests := []struct {
		forest [][]int
		wanted bool
	}{
		{[][]int{[]int{1, 2, 3}, []int{0, 0, 4}, []int{7, 6, 5}}, true},
		{[][]int{[]int{1, 2, 3}, []int{0, 0, 0}, []int{7, 6, 5}}, false},
		{[][]int{[]int{2, 3, 4}, []int{0, 0, 5}, []int{8, 6, 7}}, true},
	}
	for _, test := range tests {
		got := canReachAllTrees(test.forest)
		if got != test.wanted {
			t.Errorf("treeCount(%v) error got %v wanted %v", test.forest, got, test.wanted)
		}
	}
}
