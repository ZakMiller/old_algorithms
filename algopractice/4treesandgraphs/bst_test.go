package treesandgraphs

import "testing"

func TestBinarySearchTree(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	tree := buildBalancedBST(values)
	expected := `5
3 8
2 4 7 9
1 6
`
	if tree.String() != expected {
		t.Errorf("buildBalancedBST(%v) failed got %v expected %v", values, tree.String(), expected)
	}
}

func TestBSTLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		input                   []int
		first, second, expected int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, 4, 3},
		{[]int{-2, -1, 0, 2, 3, 4, 8}, 8, 4, 4},
	}

	for _, test := range tests {
		tree := buildBalancedBST(test.input)
		got := tree.getLowestCommonAncestor(test.first, test.second)
		if test.expected != got {
			t.Errorf("getLowestCommonAncestor failed for %v got %v expected %v", test.input, got, test.expected)
		}
	}
}
