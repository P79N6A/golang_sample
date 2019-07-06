package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Args contains rpc arguments
type Args struct {
	A, B int
}

// Arith is rpc service
type Arith int

// Add implements rpc method
func (t *Arith) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func main() {
	// register and handle
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	// established http server
	listener, e := net.Listen("tcp", ":8080")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Fatal(http.Serve(listener, nil))
}
