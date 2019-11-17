package linkedlist

func Partition(ll *LinkedList, value string) {
	node := ll.node
	var less *Node
	var more *Node
	var middle *Node
	for node != nil {
		if node.value < value {
			if less == nil {
				less = node
				ll.node = less
			} else {
				less.next = node
				less = less.next
			}
		} else {
			if more == nil {
				more = node
				middle = more
			} else {
				more.next = node
				more = more.next
			}
		}
		node = node.next
	}

	less.next = middle
	more.next = nil
}
