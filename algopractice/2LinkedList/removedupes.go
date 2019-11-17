package linkedlist

type present struct{}

func (ll *LinkedList) RemoveDupesWithMap() {
	values := make(map[string]present)
	current := ll.node
	values[current.value] = present{}
	for current.next != nil {
		_, ok := values[current.next.value]
		if ok {
			current.next = current.next.next
		} else {
			values[current.next.value] = present{}
			current = current.next
		}
	}
}

func (ll *LinkedList) RemoveDupesInPlace() {
	current := ll.node
	lookAhead := ll.node

	for current != nil {
		for lookAhead.next != nil {
			if lookAhead.next.value == current.value {
				lookAhead.next = lookAhead.next.next
			} else {
				lookAhead = lookAhead.next
			}
		}
		current = current.next
		lookAhead = current
	}
}
