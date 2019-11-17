package treesandgraphs

import "testing"

func TestQueueIsEmpty(t *testing.T) {
	tests := []struct {
		add      int
		remove   int
		expected bool
	}{
		{0, 0, true},
		{1, 1, true},
		{2, 2, true},
		{1, 0, false},
		{2, 0, false},
		{3, 1, false},
	}

	for _, test := range tests {
		q := &queue{}
		for i := 0; i < test.add; i++ {
			q.enqueue(&node{nil, nil, i})
		}
		for i := 0; i < test.remove; i++ {
			q.dequeue()
		}

		result := q.isEmpty()
		if result != test.expected {
			t.Errorf("queue.isEmpty() with %v enqueues and %v dequeues failure got %v expected %v", test.add, test.remove, result, test.expected)
		}
	}
}

func TestQueueDequeue(t *testing.T) {
	tests := []struct {
		toEnqueue []int
	}{
		{[]int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		q := &queue{}
		for _, v := range test.toEnqueue {
			q.enqueue(&node{nil, nil, v})
		}
		for _, v := range test.toEnqueue {
			dequeued := q.dequeue()
			if dequeued.value != v {
				t.Errorf("dequeue failure for %v got %v expected %v", test.toEnqueue, dequeued.value, v)
			}
		}
	}
}
