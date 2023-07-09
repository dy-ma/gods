package lib

// node represents a single node in the linked list stack.
type node struct {
	val  interface{}
	next *node
}

// LinkedListStack represents a stack implemented using a linked list.
type LinkedListStack struct {
	head *node
}

// Pop removes and returns the top element from the stack.
// It returns the value and a boolean indicating the success of the operation.
// If the stack is empty or if the LinkedListStack instance is nil, it returns (nil, false).
func (list *LinkedListStack) Pop() (interface{}, bool) {
	if list == nil {
		return nil, false
	}

	if list.head == nil {
		return nil, false
	}

	if list.head.next == nil {
		val := list.head.val
		list.head = nil
		return val, true
	}

	curhead := list.head.val
	list.head = list.head.next
	return curhead, true
}

// Push adds a new element to the top of the stack.
// It returns a boolean indicating the success of the operation.
// If the LinkedListStack instance is nil, it returns false.
func (list *LinkedListStack) Push(val interface{}) bool {
	if list == nil {
		return false
	}

	head := &node{val: val, next: nil}

	head.next = list.head
	list.head = head

	return true
}

// Peek returns the value of the top element in the stack without removing it.
// It returns the value and a boolean indicating the success of the operation.
// If the stack is empty or if the LinkedListStack instance is nil, it returns (nil, false).
func (list *LinkedListStack) Peek() (interface{}, bool) {
	if list == nil {
		return nil, false
	}

	return list.head.val, true
}
