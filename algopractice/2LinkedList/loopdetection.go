package linkedlist

func FindStartOfLoop(ll *LinkedList) *Node {
	tortoise, hare := ll.node, ll.node

	for {
		tortoise, hare = tortoise.next, hare.next.next
		if tortoise == hare {
			break
		}
	}

	floyd := ll.node

	for {
		if tortoise == floyd {
			return tortoise
		}
		tortoise, floyd = tortoise.next, floyd.next
	}
}
