package stacksandqueues

type queue struct {
	front stack
	back  stack
}

func (q *queue) push(value int) {
	if q.front.node == nil {
		node := &node{nil, value}
		q.front.node = node
		q.back.node = node
	} else {
		next := &node{nil, value}
		q.front.node.previous = next
		q.front.node = next
	}
}

func (q *queue) pop() (int, bool) {
	if q.back.node == nil {
		return 0, false
	}
	value := q.back.node.value
	q.back.node = q.back.node.previous
	return value, true
}
