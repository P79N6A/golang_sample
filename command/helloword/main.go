package main

import (
	"fmt"

	"github.com/lxb31/sample/hello"
)

func main() {
	msg := hello.CreateMessage("hello")
	//msg.
	r, err := hello.SayHello(msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r.ToString())
	}
}
