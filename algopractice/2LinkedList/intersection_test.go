package linkedlist

import "testing"

func getIntersectionCaseOne() (*LinkedList, *LinkedList, string) {
	shared := &Node{"end", nil}
	a4 := &Node{"a4", shared}
	a3 := &Node{"a3", a4}
	a2 := &Node{"a2", a3}
	a1 := &Node{"a1", a2}
	a := &LinkedList{a1}

	b2 := &Node{"b2", shared}
	b1 := &Node{"b1", b2}
	b := &LinkedList{b1}
	return a, b, "end"
}

func getIntersectionCaseTwo() (*LinkedList, *LinkedList, string) {
	shared := &Node{"end", nil}
	a4 := &Node{"a4", shared}
	a3 := &Node{"a3", a4}
	a2 := &Node{"a2", a3}
	a1 := &Node{"a1", a2}
	a := &LinkedList{a1}
	b := &LinkedList{shared}

	return a, b, "end"
}

func isValid(node *Node, value string) bool {
	if node == nil {
		if value == "" {
			return true
		}
		return false
	}
	return node.value == value
}

func TestGetIntersection(t *testing.T) {
	v1One, v1Two, v1Value := getIntersectionCaseOne()
	v2One, v2Two, v2Value := getIntersectionCaseTwo()
	tests := []struct {
		one                 *LinkedList
		two                 *LinkedList
		valueAtIntersection string
	}{
		{createLinkedList("one", "two", "three"), createLinkedList("one", "two", "three"), ""},
		{v1One, v1Two, v1Value},
		{v2One, v2Two, v2Value},
	}

	for _, test := range tests {
		actual := GetIntersection(test.one, test.two)
		if !isValid(actual, test.valueAtIntersection) {
			t.Errorf("GetIntersection(%v,%v) error actual %v expected %v", test.one, test.two, actual, test.valueAtIntersection)
		}
	}
}
