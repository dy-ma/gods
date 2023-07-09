package main

import (
	"fmt"

	"github.com/dy-ma/gods/lib/queue"
	"github.com/dy-ma/gods/lib/stack"
)

func testStackInts(stack stack.Stack) {
	values := [5]int{100, 323, 2132, -21, 123}

	for _, val := range values {
		stack.Push(val)
	}

	for i := 4; i >= 0; i-- {
		popped, _ := stack.Pop()
		if popped != values[i] {
			fmt.Printf("testStackInts Fail: %v != %v\n", values[i], popped)
		}
	}

	fmt.Println("testStackInts Complete")
}

func testStackStrings(stack stack.Stack) {
	stack.Push("First")
	stack.Push("Second")

	popped, _ := stack.Pop()
	if popped != ("Second") {
		panic("testStackStrings Fail: Pop Second")
	}

	stack.Push("Third")
	popped, _ = stack.Pop()
	if popped != ("Third") {
		panic("testStackStrings Fail: Pop Third")
	}

	peeked, ok := stack.Peek()
	if ok && peeked != "First" {
		panic("testStackStrings Fail: Peek First")
	}

	popped, _ = stack.Pop()
	if popped != ("First") {
		panic("testStackStrings Fail: Pop First")
	}

	fmt.Println("testStackStrings Complete")
}

func testQueueInts(queue queue.Queue) {
	queue.Enqueue(23)
	queue.Enqueue(340)
	queue.Enqueue(-120)

	dq, _ := queue.Dequeue()
	if dq != 23 {
		panic("testQueueInts Fail: 23")
	}

	dq, _ = queue.Dequeue()
	if dq != 340 {
		panic("testQueueInts Fail: 340")
	}

	queue.Enqueue(100)

	dq, _ = queue.Dequeue()
	if dq != -120 {
		panic("testQueueInts Fail: -120")
	}

	fmt.Println("testQueueInts Complete")
}

func main() {
	fmt.Println("== Testing Stacks ==")
	ls := new(stack.LinkedListStack)
	testStackInts(ls)
	testStackStrings(ls)

	as := new(stack.ArrayStack)
	testStackInts(as)
	testStackStrings(as)

	fmt.Println("== Testing Queues ==")
	lq := new(queue.LinkedListQueue)
	testQueueInts(lq)

	aq := new(queue.ArrayQueue)
	testQueueInts(aq)
}
