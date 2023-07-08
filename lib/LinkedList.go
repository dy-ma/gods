package lib

import (
	"errors"
)

type Node struct {
	val  interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Pop() (interface{}, error) {
	if list == nil {
		return nil, errors.New("linked list is null")
	}

	if list.head == nil {
		return nil, errors.New("pop called on empty linked list")
	}

	if list.head.next == nil {
		val := list.head.val
		list.head = nil
		return val, nil
	}

	curhead := list.head.val
	list.head = list.head.next
	return curhead, nil
}

func (list *LinkedList) Push(val interface{}) error {
	if list == nil {
		return errors.New("linked list is null")
	}

	head := &Node{val: val, next: nil}

	head.next = list.head
	list.head = head

	return nil
}
