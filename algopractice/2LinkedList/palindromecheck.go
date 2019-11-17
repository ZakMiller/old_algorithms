package linkedlist

type stack struct {
	data  []string
	index int
}

func (st *stack) push(s string) {
	st.data = append(st.data, s)
	st.index++
}

func (st *stack) pop() string {
	value := st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]
	return value
}

func lookAheadTraverse(ll *LinkedList) (*stack, *Node) {
	middle := ll.node
	lookAhead := ll.node

	s := &stack{}
	for lookAhead.next != nil {
		s.push(middle.value)
		lookAhead = lookAhead.next
		if lookAhead.next != nil {
			lookAhead = lookAhead.next
			middle = middle.next
		}
	}
	return s, middle
}

func IsPalindrome(ll *LinkedList) bool {
	if ll.node.next == nil {
		return true
	} else if ll.node.next.next == nil {
		return ll.node.value == ll.node.next.value
	}

	stack, middle := lookAheadTraverse(ll)
	laterHalf := middle.next
	for laterHalf != nil {
		expected := stack.pop()
		if expected != laterHalf.value {
			return false
		}
		laterHalf = laterHalf.next
	}
	return true
}
