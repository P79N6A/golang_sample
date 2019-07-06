package stack

import (
	"fmt"
)

func ExampleStack() {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for len(s) > 0 {
		fmt.Println(s.Pop())
	}
	// OUtput:
	//3 true
	//2 true
	//1 true

}
