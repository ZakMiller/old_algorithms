package treesandgraphs

type node struct {
	left  *node
	right *node
	value int
}

type queue struct {
	front *queueNode
	back  *queueNode
}

type queueNode struct {
	value *node
	next  *queueNode
}

func (q *queue) dequeue() *node {
	node := q.front
	q.front = q.front.next
	if q.front == nil {
		q.back = nil
	}
	return node.value
}

func (q *queue) enqueue(value *node) {
	if q.back == nil {
		q.back = &queueNode{value, nil}
		q.front = q.back
	} else {
		next := &queueNode{value, nil}
		q.back.next = next
		q.back = next
	}
}

func (q *queue) isEmpty() bool {
	return q.back == nil
}
