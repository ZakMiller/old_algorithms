package linkedlist

import "testing"

func TestLinkedList_GetKthToLast(t *testing.T) {
	tests := []struct {
		ll       *LinkedList
		k        int
		expected string
	}{
		{createLinkedList("first", "second", "third", "fourth", "fifth", "sixth"), 0, "sixth"},
		{createLinkedList("first", "second", "third", "fourth", "fifth", "sixth"), 1, "fifth"},
		{createLinkedList("first", "second", "third", "fourth", "fifth", "sixth"), 2, "fourth"},
		{createLinkedList("first", "second", "third", "fourth", "fifth", "sixth"), 3, "third"},
		{createLinkedList("first", "second", "third", "fourth", "fifth", "sixth"), 4, "second"},
		{createLinkedList("first", "second", "third", "fourth", "fifth", "sixth"), 5, "first"},
	}

	for _, test := range tests {
		actual := GetKthToLast(test.ll, test.k)
		if actual.value != test.expected {
			t.Errorf("GetKthToLast(%v, %v) failed actual %v expected %v", test.ll, test.k, actual.value, test.expected)
		}
	}
}
