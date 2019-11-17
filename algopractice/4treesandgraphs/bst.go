package treesandgraphs

func buildBalancedBST(values []int) *node {
	middleIndex := len(values) / 2

	node := &node{nil, nil, values[middleIndex]}
	if len(values) > 1 {
		left := values[0:middleIndex]
		node.left = buildBalancedBST(left)
	}
	if len(values) > 2 {
		right := values[middleIndex+1:]
		node.right = buildBalancedBST(right)
	}
	return node
}

func (n *node) isBST() bool {
	leftValid, rightValid := true, true
	if n.left != nil {
		if n.value < n.left.value {
			return false
		}
		leftValid = n.left.isBST()
	}
	if n.right != nil {
		if n.value > n.right.value {
			return false
		}
		rightValid = n.right.isBST()
	}
	return leftValid && rightValid
}

func (n *node) getLowestCommonAncestor(val1, val2 int) int {
	for {
		if n.value > val1 && n.value > val2 {
			n = n.left
		} else if n.value < val1 && n.value < val2 {
			n = n.right
		} else {
			return n.value
		}
	}
}
