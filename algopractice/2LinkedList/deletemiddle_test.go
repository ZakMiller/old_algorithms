package linkedlist

import "testing"

func TestDeleteMiddle(t *testing.T) {
	tests := []struct {
		ll                     *LinkedList
		kthToLastindexToDelete int
		expected               string
	}{
		{createLinkedList("one", "two", "three"), 1, "onethree"},
		{createLinkedList("one", "two", "three", "four", "five"), 1, "onetwothreefive"},
	}

	for _, test := range tests {
		before := test.ll.String()

		node := GetKthToLast(test.ll, test.kthToLastindexToDelete)
		originalNodeValue := node.value
		DeleteMiddle(node)
		actual := test.ll.String()
		if actual != test.expected {
			t.Errorf("DeleteMiddle(%v) for %v error actual %v expected %v", originalNodeValue, before, actual, test.expected)
		}
	}
}
