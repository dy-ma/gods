package stack

type ArrayStack struct {
	arr []interface{}
	len int
}

// Pop removes and returns the top element from the stack.
// It returns the value and a boolean indicating the success of the operation.
// If the stack is empty or if the ArrayStack instance is nil, it returns (nil, false).
func (stack *ArrayStack) Pop() (interface{}, bool) {
	if stack == nil || stack.len == 0 {
		return nil, false
	}

	head := stack.arr[0]
	// Remove first element
	for i := 0; i < stack.len-1; i++ {
		stack.arr[i] = stack.arr[i+1]
	}
	stack.len--
	// Slice to remove duplicated last element
	stack.arr = stack.arr[:stack.len]
	return head, true
}

// Push adds a new element to the top of the stack.
// It returns a boolean indicating the success of the operation.
// If the ArrayStack instance is nil, it returns false.
func (stack *ArrayStack) Push(val interface{}) bool {
	if stack == nil {
		return false
	}

	// Increase size of slice by one
	stack.arr = append(stack.arr, nil)

	// Right shift all elements
	for i := stack.len; i > 0; i-- {
		stack.arr[i] = stack.arr[i-1]
	}

	stack.arr[0] = val
	stack.len++
	return true
}

// Peek returns the value of the top element in the stack without removing it.
// It returns the value and a boolean indicating the success of the operation.
// If the stack is empty or if the ArrayStack instance is nil, it returns (nil, false).
func (stack *ArrayStack) Peek() (interface{}, bool) {
	if stack == nil || stack.len == 0 {
		return nil, false
	}

	return stack.arr[0], true
}
