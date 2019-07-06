package queue

import (
	"fmt"
)

func ExampleFifoQueue() {
	q := NewFifoQueue()
	q.Add(1)
	q.Add(2)

	fmt.Println(q)
	fmt.Println(q.Peek())
	fmt.Println(q.Poll())
	fmt.Println(q.Poll())
	//Output:
	//[1 2]
	//1
	//1
	//2
}
