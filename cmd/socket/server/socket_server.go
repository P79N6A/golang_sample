package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("net.Listen: " + err.Error())
	}
	defer listener.Close()
	fmt.Printf("Listen on :8080")

	for {
		connection, err := listener.Accept()
		if err != nil {
			panic("Accept: " + err.Error())
		}
		go handle(connection)
	}
}

func handle(connection net.Conn) {
	buf := make([]byte, 64)
	for {
		len, err := connection.Read(buf)
		if len > 0 {
			fmt.Printf("Read %v bytes\n", len)
		}

		if err == io.EOF {
			fmt.Printf("Client close.")
			return
		} else if err != nil {
			fmt.Printf(err.Error())
			return
		}

	}

}
