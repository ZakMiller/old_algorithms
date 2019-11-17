package stacksandqueues

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	previous *node
	value    int
}

type stack struct {
	node *node
}

type stackWrapper struct {
	stack
	size int
}

func (s *stack) push(value int) {
	if s.node == nil {
		s.node = &node{nil, value}
	} else {
		s.node = &node{s.node, value}
	}
}

func (s *stack) pop() int {
	value := s.node.value
	s.node = s.node.previous
	return value
}

func (s *stack) isEmpty() bool {
	return s.node == nil
}

func (s *stack) peek() int {
	return s.node.value
}

func (s *stack) String() string {
	b := strings.Builder{}
	node := s.node
	for node != nil {
		b.WriteString(fmt.Sprintf("%v ", strconv.Itoa(node.value)))
		node = node.previous
	}
	return b.String()
}

type stackOfStacks struct {
	stacks    []*stackWrapper
	stackSize int
}

func (ss *stackOfStacks) String() string {
	b := strings.Builder{}
	for _, s := range ss.stacks {
		node := s.node
		for node != nil {
			val := fmt.Sprintf("%v,", strconv.Itoa(node.value))
			b.WriteString(val)
			node = node.previous
		}
		b.WriteString(" ")
	}
	return b.String()
}

func New(stackSize int) *stackOfStacks {
	return &stackOfStacks{stacks: []*stackWrapper{{}}, stackSize: stackSize}
}

func (ss *stackOfStacks) push(value int) {
	if ss.stacks[len(ss.stacks)-1].size == ss.stackSize {
		ss.stacks = append(ss.stacks, &stackWrapper{})
	}
	ss.stacks[len(ss.stacks)-1].push(value)
	ss.stacks[len(ss.stacks)-1].size++
}

func (ss *stackOfStacks) pop() (int, bool) {
	if ss.stacks[len(ss.stacks)-1].size == 0 {
		if len(ss.stacks)-1 == 0 {
			return 0, false
		} else {
			ss.stacks = ss.stacks[0 : len(ss.stacks)-1]
		}
	}

	return ss.popAt(len(ss.stacks) - 1)
}

func (ss *stackOfStacks) popAt(index int) (int, bool) {
	if index >= len(ss.stacks) {
		return 0, false
	}

	value := ss.stacks[index].pop()
	ss.stacks[index].size--
	return value, true
}
