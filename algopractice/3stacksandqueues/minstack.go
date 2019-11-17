package stacksandqueues

import "strings"

type pointerStack struct {
	head *minNode
	min  *minNode
}

type minNode struct {
	value       string
	previous    *minNode
	previousMin *minNode
}

func (ps *pointerStack) push(s string) {
	if ps.head == nil {
		ps.head = &minNode{s, nil, nil}
		ps.min = ps.head
		return
	}
	previous := &minNode{s, ps.head, nil}
	if s < ps.min.value {
		previousMin := ps.min
		previous.previousMin = previousMin
		ps.min = previous
	}

	ps.head = previous
}

func (ps *pointerStack) pop() (string, bool) {
	if ps.isEmpty() {
		return "", false
	}
	v := ps.head.value
	if ps.min == ps.head {
		ps.min = ps.min.previousMin
	}
	ps.head = ps.head.previous

	return v, true
}

func (ps *pointerStack) peek() (string, bool) {
	if ps.isEmpty() {
		return "", false
	}
	return ps.head.value, true
}

func (ps *pointerStack) isEmpty() bool {
	return ps.head == nil
}

func (ps *pointerStack) getMin() (string, bool) {
	if ps.min == nil {
		return "", false
	}
	return ps.min.value, true
}

func (ps *pointerStack) String() string {
	b := strings.Builder{}
	minNode := ps.head
	for minNode != nil {
		b.WriteString(minNode.value)
		minNode = minNode.previous
	}
	return b.String()
}
