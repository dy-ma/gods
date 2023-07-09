package queue

type node struct {
	val        interface{}
	next, prev *node
}

type LinkedListQueue struct {
	back, front *node
}

// Enqueue adds a new element to the back of the queue.
// It returns a boolean indicating the success of the operation.
// If the Queue is nil, it returns false.
func (q *LinkedListQueue) Enqueue(val interface{}) bool {
	if q == nil {
		return false
	}

	// Initialize a new node with arg value
	newNode := &node{val: val, next: nil, prev: nil}

	// Zero elements
	if q.back == nil && q.front == nil {
		q.back = newNode
		q.front = newNode
	} else {
		newNode.next = q.back
		q.back.prev = newNode
		q.back = newNode
	}

	return true
}

// Dequeue removes the element at the front of the queue.
// It returns the value of the element and a boolean indicating the success of the operation.
// If the Queue is nil or empty, it returns (nil, false)
func (q *LinkedListQueue) Dequeue() (interface{}, bool) {
	if q == nil || q.front == nil {
		return nil, false
	}

	val := q.front.val
	q.front = q.front.prev

	// Update back pointer front was only element
	if q.front == nil {
		q.back = nil
	}

	return val, true
}
