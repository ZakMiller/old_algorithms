package linkedlist

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		one      *LinkedList
		two      *LinkedList
		expected *LinkedList
	}{
		{createLinkedList("1", "2", "3"), createLinkedList("4", "5", "6"), createLinkedList("5", "7", "9")},
		{createLinkedList("1", "2", "9"), createLinkedList("4", "5", "6"), createLinkedList("5", "8", "5")},
		{createLinkedList("1", "2"), createLinkedList("4", "0", "0"), createLinkedList("4", "1", "2")},
	}

	for _, test := range tests {
		actual := Sum(test.one, test.two)
		if actual.String() != test.expected.String() {
			t.Errorf("Sum(%v, %v) error actual %v expected %v", test.one, test.two, actual, test.expected)
		}
	}
}

func TestSumReversed(t *testing.T) {
	tests := []struct {
		one      *LinkedList
		two      *LinkedList
		expected *LinkedList
	}{
		{createLinkedList("1", "2", "3"), createLinkedList("4", "5", "6"), createLinkedList("5", "7", "9")},
		{createLinkedList("9", "2", "9"), createLinkedList("8", "5", "6"), createLinkedList("7", "8", "5", "1")},
		{createLinkedList("1", "2"), createLinkedList("0", "0", "4"), createLinkedList("1", "2", "4")},
	}

	for _, test := range tests {
		actual := SumReversed(test.one, test.two)
		if actual.String() != test.expected.String() {
			t.Errorf("SumReversed(%v, %v) error actual %v expected %v", test.one, test.two, actual, test.expected)
		}
	}
}
