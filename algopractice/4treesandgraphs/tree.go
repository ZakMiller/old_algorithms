package treesandgraphs

import (
	"container/list"
	"math"
	"strconv"
	"strings"
)

func createNodes(values ...int) []*node {
	nodes := make([]*node, 0, len(values))
	for _, v := range values {
		nodes = append(nodes, &node{nil, nil, v})
	}
	return nodes
}

func buildTree(nodes []*node) *node {
	rootIndex := 0
	for i := 1; i < len(nodes); i++ {
		for nodes[rootIndex].left != nil && nodes[rootIndex].right != nil {
			rootIndex++
		}
		if nodes[rootIndex].left == nil {
			nodes[rootIndex].left = nodes[i]
		} else {
			nodes[rootIndex].right = nodes[i]
		}
	}
	return nodes[0]
}

func createTree(values ...int) *node {
	nodes := createNodes(values...)
	root := buildTree(nodes)
	return root
}

func (n *node) String() string {
	w := &writer{}
	q := &queue{}
	q.enqueue(n)

	for !q.isEmpty() {
		next := q.dequeue()
		w.store(next.value)
		if next.left != nil {
			q.enqueue(next.left)
		}
		if next.right != nil {
			q.enqueue(next.right)
		}
	}

	return w.String()
}

func getListsForRows(n *node) []*list.List {
	lists := make([]*list.List, 1)
	lists[0] = list.New()
	l := list.New()
	l.PushBack(n)
	for l.Len() > 0 {
		listNode := l.Front()
		n = listNode.Value.(*node)
		l.Remove(listNode)
		if n.left != nil {
			l.PushBack(n.left)
		}
		if n.right != nil {
			l.PushBack(n.right)
		}
		currentRowIndex := len(lists)
		currentRow := lists[len(lists)-1]
		currentRow.PushBack(n.value)
		expectedNodesInRow := int(math.Pow(2, float64(currentRowIndex-1)))
		if currentRow.Len() == expectedNodesInRow {
			lists = append(lists, list.New())
		}
	}
	return lists
}

func (n *node) isBalanced() (int, bool) {
	var leftNodes, rightNodes int
	leftBalanced, rightBalanced := true, true
	if n.left != nil {
		leftNodes, leftBalanced = n.left.isBalanced()
	}
	if n.right != nil {
		rightNodes, rightBalanced = n.right.isBalanced()
	}
	if !leftBalanced || !rightBalanced || math.Abs(float64(leftNodes-rightNodes)) > 1 {
		return 0, false
	}
	return leftNodes + rightNodes + 1, true
}

func newUnbalancedTree() *node {
	nodeCount := 4
	nodes := make([]*node, 0, nodeCount)
	nodes = append(nodes, &node{nil, nil, 0})
	for i := 1; i < nodeCount; i++ {
		nodes = append(nodes, &node{nodes[i-1], nil, i})
	}
	return nodes[len(nodes)-1]
}

func printTreeInOrderR(b *strings.Builder, n *node) {
	if n.left != nil {
		printTreeInOrderR(b, n.left)
	}
	b.WriteString(strconv.Itoa(n.value))
	if n.right != nil {
		printTreeInOrderR(b, n.right)
	}
}
func (n *node) printTreeInOrder() string {
	b := &strings.Builder{}
	printTreeInOrderR(b, n)
	return b.String()
}

func printTreePreOrderR(b *strings.Builder, n *node) {
	b.WriteString(strconv.Itoa(n.value))

	if n.left != nil {
		printTreePreOrderR(b, n.left)
	}
	if n.right != nil {
		printTreePreOrderR(b, n.right)
	}
}
func (n *node) printTreePreOrder() string {
	b := &strings.Builder{}
	printTreePreOrderR(b, n)
	return b.String()
}

func printTreePostOrderR(b *strings.Builder, n *node) {
	if n.left != nil {
		printTreePostOrderR(b, n.left)
	}
	if n.right != nil {
		printTreePostOrderR(b, n.right)
	}
	b.WriteString(strconv.Itoa(n.value))
}

func (n *node) printTreePostOrder() string {
	b := &strings.Builder{}
	printTreePostOrderR(b, n)
	return b.String()
}

func findLCAOfNode(root *node, valOne int, valTwo int) *node {
	var left, right *node
	if root.left != nil {
		left = findLCAOfNode(root.left, valOne, valTwo)
	}
	if root.right != nil {
		right = findLCAOfNode(root.right, valOne, valTwo)
	}

	if root.value == valOne || root.value == valTwo {
		return root
	}
	if left != nil && right != nil {
		if (left.value == valOne || left.value == valTwo) && (right.value == valOne || right.value == valTwo) {
			return root
		}
	}

	if left != nil && (left.value == valOne || left.value == valTwo) {
		return left
	}

	if right != nil && (right.value == valOne || right.value == valTwo) {
		return right
	}
	return nil
}

func (n *node) findLowestCommonAncestor(valOne int, valTwo int) int {
	return findLCAOfNode(n, valOne, valTwo).value
}
