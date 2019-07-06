package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// Args contains rpc arguments
// 协议需要独立出来
type Args struct {
	A, B int
}

func main() {
	// 获取http客户端
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 发起调用
	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Add", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d+%d=%d", args.A, args.B, reply)
}
