package stacksandqueues

import (
	"fmt"
	"strconv"
	"strings"
)

type simpleStack struct {
	data  [100]int
	index int
}

func (ss *simpleStack) push(value int) bool {
	if ss.index > 99 {
		return false
	}
	ss.data[ss.index] = value
	ss.index++
	return true
}

func (ss *simpleStack) pop() (int, bool) {
	if ss.index <= 0 {
		return 0, false
	}
	ss.index--
	return ss.data[ss.index], true
}

type stacks struct {
	data     [100]string
	pointers [100]int
	head1    int
	head2    int
	head3    int
	empty    simpleStack
}

func (s *stacks) String() string {
	b := strings.Builder{}
	for _, d := range s.data {
		b.WriteString(fmt.Sprintf("%v,", d))
	}
	b.WriteString("  ")
	for _, p := range s.pointers {
		b.WriteString(fmt.Sprintf("%v,", strconv.Itoa(p)))
	}
	return b.String()
}

func createStacks() *stacks {
	s := &stacks{head1: 0, head2: 1, head3: 2}
	for i := 3; i < 100; i++ {
		s.empty.push(i)
	}
	return s
}

func (s *stacks) pushOne(v string) bool {
	next, success := s.empty.pop()
	if !success {
		return false
	}

	s.data[s.head1] = v
	s.pointers[next] = s.head1
	s.head1 = next
	return true
}

func (s *stacks) pushTwo(v string) bool {
	next, success := s.empty.pop()
	if !success {
		return false
	}

	s.data[s.head2] = v
	s.pointers[next] = s.head2
	s.head2 = next
	return true
}

func (s *stacks) pushThree(v string) bool {
	next, success := s.empty.pop()
	if !success {
		return false
	}

	s.data[s.head3] = v
	s.pointers[next] = s.head3
	s.head3 = next
	return true
}

func (s *stacks) popOne() (string, bool) {
	last := s.pointers[s.head1]
	if last == 0 {
		return "", false
	}
	s.pointers[s.head1] = 0
	v := s.data[s.head1]
	s.data[s.head1] = ""
	s.empty.push(s.head1)
	s.head1 = last
	return v, true
}

func (s *stacks) popTwo() (string, bool) {
	last := s.pointers[s.head2]
	if last == 0 {
		return "", false
	}
	s.pointers[s.head2] = 0
	v := s.data[s.head2]
	s.data[s.head2] = ""
	s.empty.push(s.head2)
	s.head2 = last
	return v, true
}

func (s *stacks) popThree() (string, bool) {
	last := s.pointers[s.head3]
	if last == 0 {
		return "", false
	}
	s.pointers[s.head3] = 0
	v := s.data[s.head3]
	s.data[s.head3] = ""
	s.empty.push(s.head3)
	s.head3 = last
	return v, true
}
