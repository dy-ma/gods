package queue

type ArrayQueue struct {
	arr []interface{}
	len int
}

// Enqueue adds an element to the end of the queue array.
// It returns a boolean indicating the success of the operation.
// If the Queue is nil, it returns false.
func (q *ArrayQueue) Enqueue(val interface{}) bool {
	if q == nil {
		return false
	}

	q.arr = append(q.arr, val)
	q.len++
	return true
}

// Dequeue removes an element from the start of the queue array.
// It returns the value of the elmenet and a boolean indicating the success of the operation.
// If the queue is nil or empty, it returns false.
func (q *ArrayQueue) Dequeue() (interface{}, bool) {
	if q == nil || q.len == 0 {
		return nil, false
	}

	val := q.arr[0]
	// Remove first element by shifting all left
	for i := 0; i < q.len-1; i++ {
		q.arr[i] = q.arr[i+1]
	}
	q.len--
	// Remove last duplicate element
	q.arr = q.arr[:q.len]
	return val, true
}
