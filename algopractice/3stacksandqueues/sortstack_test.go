package stacksandqueues

import "testing"

func createStack(values ...int) *stack {
	s := &stack{}
	for i := len(values) - 1; i >= 0; i-- {
		s.push(values[i])
	}
	return s
}

func TestSortStack(t *testing.T) {
	tests := []struct {
		s        *stack
		expected *stack
	}{
		{createStack(5, 4, 3, 2, 1), createStack(1, 2, 3, 4, 5)},
		{createStack(5, 4, 1, 2, 3), createStack(1, 2, 3, 4, 5)},
		{createStack(5, 4, 1, 3, 2), createStack(1, 2, 3, 4, 5)},
		{createStack(5), createStack(5)},
		{createStack(), createStack()},
	}

	for _, test := range tests {
		before := test.s.String()
		test.s.sort()
		after := test.s.String()
		if test.expected.String() != after {
			t.Errorf("sort(%v) error actual %v expected %v", before, after, test.expected)
		}
	}
}
