package linkedlist

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		ll       *LinkedList
		expected bool
	}{
		{createLinkedList("a", "b", "c", "c", "b", "a"), true},
		{createLinkedList("a", "b", "c", "b", "a"), true},
		{createLinkedList("f", "b", "c", "c", "b", "a"), false},
		{createLinkedList("a"), true},
		{createLinkedList("a", "b"), false},
		{createLinkedList("a", "a"), true},
	}

	for _, test := range tests {
		actual := IsPalindrome(test.ll)
		if actual != test.expected {
			t.Errorf("IsPalindrome(%v) error actual %v expected %v", test.ll, actual, test.expected)
		}
	}
}
