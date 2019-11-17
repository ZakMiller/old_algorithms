package stacksandqueues

import "testing"

func TestQueue(t *testing.T) {
	q := &queue{}
	q.push(1)
	q.push(2)
	q.push(3)
	q.push(4)
	val, ok := q.pop()
	if val != 1 || ok != true {
		t.Errorf("queue pop failure actual %v,%v expected %v,%v", val, ok, 1, true)
	}
	val, ok = q.pop()
	if val != 2 || ok != true {
		t.Errorf("queue pop failure actual %v,%v expected %v,%v", val, ok, 2, true)
	}
	val, ok = q.pop()
	if val != 3 || ok != true {
		t.Errorf("queue pop failure actual %v,%v expected %v,%v", val, ok, 3, true)
	}
	val, ok = q.pop()
	if val != 4 || ok != true {
		t.Errorf("queue pop failure actual %v,%v expected %v,%v", val, ok, 4, true)
	}
	val, ok = q.pop()
	if val != 0 || ok != false {
		t.Errorf("queue pop failure actual %v,%v expected %v,%v", val, ok, 0, false)
	}
}
