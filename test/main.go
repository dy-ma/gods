package main

import (
	"fmt"

	"github.com/dy-ma/gods/lib"
)

func testInts(stack lib.Stack) {
	values := [5]int{100, 323, 2132, -21, 123}

	for _, val := range values {
		stack.Push(val)
	}

	for i := 4; i >= 0; i-- {
		popped, _ := stack.Pop()
		if popped != values[i] {
			fmt.Printf("testInts() Fail: %v != %v\n", values[i], popped)
		}
	}

	fmt.Println("testInts() Complete")
}

func testStrings(stack lib.Stack) {
	stack.Push("First")
	stack.Push("Second")

	popped, _ := stack.Pop()
	if popped != ("Second") {
		panic("testInts() Fail: Pop Second")
	}

	stack.Push("Third")
	popped, _ = stack.Pop()
	if popped != ("Third") {
		panic("testInts() Fail: Pop Third")
	}

	peeked, ok := stack.Peek()
	if ok && peeked != "First" {
		panic("testInts() Fail: Peek")
	}

	popped, _ = stack.Pop()
	if popped != ("First") {
		panic("testInts() Fail")
	}

	fmt.Println("testStrings() Complete")
}

func main() {
	l := new(lib.LinkedListStack)
	testInts(l)
	testStrings(l)

	a := new(lib.ArrayStack)
	testInts(a)
	testStrings(a)
}
