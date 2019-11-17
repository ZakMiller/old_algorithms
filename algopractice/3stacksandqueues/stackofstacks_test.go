package stacksandqueues

import (
	"testing"
)

func TestStackOfStacks(t *testing.T) {
	ss := New(5)
	ss.push(1)
	ss.push(2)
	ss.push(3)
	ss.push(4)
	ss.push(5)
	ss.push(1)
	ss.push(2)
	ss.push(3)
	ss.push(4)
	ss.push(5)
	ss.push(1)
	ss.push(2)
	ss.push(3)
	ss.push(4)
	ss.push(5)
	threeFull := ss.String()
	expected := `5,4,3,2,1, 5,4,3,2,1, 5,4,3,2,1, `
	if threeFull != expected {
		t.Errorf("Stack of stacks error actual %v expected %v", threeFull, expected)
	}
	ss.pop()
	ss.pop()
	ss.pop()
	ss.pop()
	ss.pop()
	ss.pop()
	fewerStacks := ss.String()
	expected = `5,4,3,2,1, 4,3,2,1, `
	if fewerStacks != expected {
		t.Errorf("Stack of stacks error actual %v expected %v", fewerStacks, expected)
	}

	ss.popAt(0)
	popAt := ss.String()
	expected = `4,3,2,1, 4,3,2,1, `
	if popAt != expected {
		t.Errorf("Stack of stacks error actual %v expected %v", popAt, expected)
	}
}
