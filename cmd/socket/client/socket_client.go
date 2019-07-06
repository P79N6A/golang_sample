package main

import (
	"fmt"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic("net.Dial: " + err.Error())
	}
	connection.Write([]byte("hello, world, this is socket client."))
	fmt.Fprintf(connection, "GET / HTTP/1.0\r\n\r\n")
	connection.Close()
}
