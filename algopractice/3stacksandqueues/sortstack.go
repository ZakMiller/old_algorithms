package stacksandqueues

func (s *stack) sort() {
	s2 := &stack{}

	for !s.isEmpty() {
		val := s.pop()
		for !s2.isEmpty() && s2.peek() > val {
			s.push(s2.pop())
		}
		s2.push(val)
	}

	for !s2.isEmpty() {
		s.push(s2.pop())
	}
}
