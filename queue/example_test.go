package queue

import (
	"fmt"
)

func Example() {
	// Create new queue and push objects to the ned of the queue including a duplicate
	q := New()
	q.Push(2)
	q.Push(5)
	q.Push(1)
	q.Push(3)
	q.Push(5)

	// Print all pops, the duplicate value '5' should be at the end of the que
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	// Output:
	// 2
	// 1
	// 3
	// 5
}
