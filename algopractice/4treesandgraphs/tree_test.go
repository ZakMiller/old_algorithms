package treesandgraphs

import (
	"container/list"
	"fmt"
	"strings"
	"testing"
)

func TestTree(t *testing.T) {
	tree := createTree(1, 2, 3, 4, 5, 6, 7, 8, 9)
	expected := `1
2 3
4 5 6 7
8 9
`
	if tree.String() != expected {
		t.Errorf("Tree creation issue got %v expected %v", tree.String(), expected)
	}
}

func TestTreeCreateNodes(t *testing.T) {
	values := []int{5, 2, 1, 9, 3, 10}
	nodes := createNodes(values...)
	if len(values) != len(nodes) {
		t.Errorf("createNodes(%v) failure got length %v expected length %v", values, len(nodes), len(values))
	}
	for i, node := range nodes {
		if values[i] != node.value {
			t.Errorf("createNodes(%v) failure on index %v got %v expected %v", values, i, node.value, values[i])
		}
	}
}

func TestTreeBuildTree(t *testing.T) {
	nodes := []*node{&node{nil, nil, 4}, &node{nil, nil, 8}, &node{nil, nil, 100}, &node{nil, nil, 101}}
	root := buildTree(nodes)
	if root.value != 4 {
		t.Errorf("buildTree error for root node expected %v got %v", 4, root.value)
	}
	if root.left.value != 8 {
		t.Errorf("buildTree error for left node expected %v got %v", 8, root.left.value)
	}
	if root.right.value != 100 {
		t.Errorf("buildTree error for right node expected %v got %v", 100, root.right.value)
	}
	if root.left.left.value != 101 {
		t.Errorf("buildTree error for left.left node expected %v got %v", 101, root.left.left.value)
	}
}

func listString(l *list.List) string {
	b := strings.Builder{}
	e := l.Front()
	for e != nil {
		b.WriteString(fmt.Sprintf("%v ", e.Value))
		e = e.Next()
	}
	return b.String()
}

func TestListString(t *testing.T) {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	if listString(l) != "1 2 3 " {
		t.Errorf("listString([1, 2, 3]) failed got %v expected %v", listString(l), "1 2 3 ")
	}
}
func TestGetListForEachLevelOfTree(t *testing.T) {
	bst := buildBalancedBST([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	rows := getListsForRows(bst)
	if len(rows) != 4 {
		t.Errorf("getListsForRows(bst0-9) failed got len %v wanted %v", len(rows), 4)
	}
	expectedRows := [][]int{[]int{5}, []int{2, 8}, []int{1, 4, 7, 9}, []int{0, 3, 6}}
	for i, row := range expectedRows {
		got := rows[i]
		e := got.Front()
		for j, v := range row {
			if v != e.Value {
				t.Errorf("getListsForRows error for row %v index %v got %v wanted %v", i, j, e.Value, v)
			}
			e = e.Next()
		}
		if e != nil {
			t.Errorf("getListsForRows error expected fewer nodes in line %v got at least %v", i, len(row)+1)
		}
	}
}

func TestTreeIsBalanced(t *testing.T) {
	unbalanced := newUnbalancedTree()
	_, unbalancedIsBalanced := unbalanced.isBalanced()
	if unbalancedIsBalanced {
		t.Errorf("Unbalanced tree reported balanced")
	}
	balanced := buildBalancedBST([]int{1, 2, 3, 4, 5, 6, 7})
	_, balancedIsBalanced := balanced.isBalanced()
	if balancedIsBalanced == false {
		t.Errorf("Balanced tree reported not balanced")
	}
}

func TestIsBST(t *testing.T) {
	nonBST := createTree(1, 2, 3, 4, 5, 6, 7, 8, 9)
	if nonBST.isBST() {
		t.Errorf("nonBST shows as BST")
	}
	BST := buildBalancedBST([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	if !BST.isBST() {
		t.Errorf("BST shows as nonBST")
	}
}

func TestPrintTree(t *testing.T) {
	tree := buildBalancedBST([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	pre := tree.printTreePreOrder()
	expectedPre := "532148769"
	if pre != expectedPre {
		t.Errorf("Pre order traversal failure got %v expected %v", pre, expectedPre)
	}
	post := tree.printTreePostOrder()
	expectedPost := "124367985"
	if post != expectedPost {
		t.Errorf("Post order traversal failure got %v expected %v", post, expectedPost)
	}
	inOrder := tree.printTreeInOrder()
	expected := "123456789"
	if inOrder != expected {
		t.Errorf("In order traversal failure got %v expected %v", inOrder, expected)
	}
}

func TestFindLowestCommonAncestorOfBinaryTree(t *testing.T) {
	tree := createTree(-2, -1, 0, 3, 4, 8)
	res := tree.findLowestCommonAncestor(8, 4)
	t.Errorf("findLowestCommonAncestor(%v, %v) returned %v", 8, 4, res)
}
