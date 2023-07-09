package minqueue

type ArrayMinQueue struct {
	arr []Comparable
	len int
}

func (mq *ArrayMinQueue) IsEmpty() bool {
	return mq != nil && mq.len == 0
}

func (mq *ArrayMinQueue) Insert(val Comparable) bool {
	if mq == nil {
		return false
	}
	mq.arr = append(mq.arr, val)
	mq.len++
	return true
}

// Pull removes the element with the minimum value in the Queue.
// It returns the element and a boolean indicating the success of the operation.
// If the queue is nil or empty, it returns (nil, false)
func (mq *ArrayMinQueue) Pull() (Comparable, bool) {
	if mq == nil || mq.len == 0 {
		return nil, false
	}

	// Get Minimum element
	min := mq.arr[0]
	argmin := 0
	for i := 1; i < mq.len; i++ {
		if mq.arr[i].LessThan(min) {
			min = mq.arr[i]
			argmin = i
		}
	}

	// Remove Minimum element
	// Shift all elements after min to left
	for i := argmin; i < mq.len-1; i++ {
		mq.arr[i] = mq.arr[i+1]
	}
	// Remove final duplicate element
	mq.arr = mq.arr[:mq.len-1]
	mq.len--
	return min, true
}
