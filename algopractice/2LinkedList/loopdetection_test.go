package linkedlist

import "testing"

func getLoopDetectionCaseOne() *LinkedList {
	loop1 := &Node{"end", nil}
	loop2 := &Node{"a4", loop1}
	loop3 := &Node{"a3", loop2}
	loop1.next = loop3
	a2 := &Node{"a2", loop1}
	a1 := &Node{"a1", a2}
	a := &LinkedList{a1}
	return a
}

func getLoopDetectionCaseTwo() *LinkedList {
	loop1 := &Node{"end", nil}
	loop2 := &Node{"a4", loop1}
	loop3 := &Node{"a3", loop2}
	loop4 := &Node{"a3", loop3}
	loop5 := &Node{"a3", loop4}
	loop6 := &Node{"a3", loop5}
	loop1.next = loop6
	a2 := &Node{"a2", loop1}
	a1 := &Node{"a1", a2}
	a := &LinkedList{a1}
	return a
}

func getLoopDetectionCaseThree() *LinkedList {
	loop1 := &Node{"end", nil}
	loop2 := &Node{"a4", loop1}
	loop3 := &Node{"a3", loop2}
	loop4 := &Node{"a3", loop3}
	loop5 := &Node{"a3", loop4}
	loop6 := &Node{"a3", loop5}
	loop1.next = loop6
	a6 := &Node{"a6", loop1}
	a5 := &Node{"a5", a6}
	a4 := &Node{"a4", a5}
	a3 := &Node{"a3", a4}
	a2 := &Node{"a2", a3}
	a1 := &Node{"a1", a2}
	a := &LinkedList{a1}
	return a
}

func TestFindStartOfLoop(t *testing.T) {
	tests := []struct {
		ll       *LinkedList
		expected string
	}{
		{getLoopDetectionCaseOne(), "end"},
		{getLoopDetectionCaseTwo(), "end"},
		{getLoopDetectionCaseThree(), "end"},
	}

	for _, test := range tests {
		actual := FindStartOfLoop(test.ll).value
		if actual != test.expected {
			t.Errorf("FindStartOfLoop(%v) error actual %v expected %v", test.ll, actual, test.expected)
		}
	}
}
