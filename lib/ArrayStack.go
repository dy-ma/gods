package lib

type ArrayStack struct {
	arr []interface{}
	len int
}

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

func (stack *ArrayStack) Peek() (interface{}, bool) {
	if stack == nil || stack.len == 0 {
		return nil, false
	}

	return stack.arr[0], true
}
