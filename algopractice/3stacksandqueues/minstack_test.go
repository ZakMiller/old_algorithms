package stacksandqueues

import "testing"

func verify(s *pointerStack, e string, t *testing.T) {
	r, _ := s.getMin()

	if r != e {
		t.Errorf("getMin(%v) error actual %v expected %v", s, r, e)
	}
}

func TestMinStack(t *testing.T) {
	s := &pointerStack{}
	s.push("f")
	verify(s, "f", t)
	s.push("g")
	verify(s, "f", t)
	s.push("z")
	verify(s, "f", t)
	s.push("b")
	verify(s, "b", t)
	s.pop()
	verify(s, "f", t)
	s.push("a")
	verify(s, "a", t)
	s.pop()
	verify(s, "f", t)
	s.pop()
	verify(s, "f", t)
}
