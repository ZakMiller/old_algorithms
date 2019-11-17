package linkedlist

import (
	"testing"
)

func areDuplicates(ll *LinkedList) bool {
	m := make(map[string]present)
	node := ll.node
	for node != nil {
		_, ok := m[node.value]
		if ok {
			return true
		}
		m[node.value] = present{}
		node = node.next
	}
	return false
}

func TestLinkedList_RemoveDupesWithMap(t *testing.T) {
	ll := createLinkedList("hi", "hello", "how are you", "today", "hello", "hi", "brown")
	if !areDuplicates(ll) {
		t.Errorf("there are duplicates but we're not registering that there are")
	}
	ll.RemoveDupesWithMap()
	if areDuplicates(ll) {
		t.Errorf("RemoveDuplicates for %v returned true expected false", ll)
	}
}

func TestLinkedList_RemoveDupesInPlace(t *testing.T) {
	ll := createLinkedList("hi", "hello", "how are you", "today", "hello", "hi", "brown")
	if !areDuplicates(ll) {
		t.Errorf("there are duplicates but we're not registering that there are")
	}
	ll.RemoveDupesInPlace()
	if areDuplicates(ll) {
		t.Errorf("RemoveDuplicates for %v returned true expected false", ll)
	}
}
