package linkedlist

func GetKthToLast(ll *LinkedList, k int) *Node {
	lookAhead := ll.node
	finder := ll.node

	for i := 0; i < k; i++ {
		lookAhead = lookAhead.next
	}

	for lookAhead.next != nil {
		lookAhead = lookAhead.next
		finder = finder.next
	}

	return finder
}
