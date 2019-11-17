package linkedlist

import "testing"

func TestPartition(t *testing.T) {
	tests := []struct {
		ll       *LinkedList
		value    string
		expected *LinkedList
	}{
		{createLinkedList("a", "b", "c", "a", "d", "e", "z", "a"), "c", createLinkedList("a", "b", "a", "a", "c", "d", "e", "z")},
	}

	for _, test := range tests {
		before := test.ll
		Partition(test.ll, test.value)
		if test.expected.String() != test.ll.String() {
			t.Errorf("Partition(%v, %v) error expected %v actual %v", before, test.value, test.expected, test.ll)
		}

	}
}
