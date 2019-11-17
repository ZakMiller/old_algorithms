package linkedlist

import (
	"fmt"
	"strings"
)

type Node struct {
	value string
	next  *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("Node{value: %v,next: %v}", n.value, n.next)
}

type LinkedList struct {
	node *Node
}

func (ll *LinkedList) String() string {
	b := strings.Builder{}
	node := ll.node
	for node != nil {
		b.WriteString(node.value)
		node = node.next
	}
	return b.String()
}

func createLinkedList(values ...string) *LinkedList {
	ll := &LinkedList{}
	ll.node = &Node{
		value: values[0],
	}
	node := ll.node
	for i := 1; i < len(values); i++ {
		next := &Node{
			value: values[i],
		}
		node.next = next
		node = next
	}
	return ll
}
