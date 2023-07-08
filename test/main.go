package main

import (
	"fmt"

	"github.com/dy-ma/gods/lib"
)

func testInts() {
	values := [5]int{100, 323, 2132, -21, 123}
	l := new(lib.LinkedList)

	for _, val := range values {
		l.Push(val)
	}

	for i := 4; i >= 0; i-- {
		popped, _ := l.Pop()
		if popped != values[i] {
			fmt.Printf("testInts() Fail: %v != %v\n", values[i], popped)
		}
	}

	fmt.Println("testInts() Complete")
}

func testStrings() {
	l := new(lib.LinkedList)

	l.Push("First")
	l.Push("Second")

	popped, _ := l.Pop()
	if popped != ("Second") {
		panic("testInts() Fail")
	}

	l.Push("Third")
	popped, _ = l.Pop()
	if popped != ("Third") {
		panic("testInts() Fail")
	}

	popped, _ = l.Pop()
	if popped != ("First") {
		panic("testInts() Fail")
	}

	fmt.Println("testStrings() Complete")
}

func main() {
	testInts()
	testStrings()
}
