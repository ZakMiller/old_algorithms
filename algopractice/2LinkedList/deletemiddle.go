package linkedlist

func DeleteMiddle(node *Node) {
	*node = *node.next
}
