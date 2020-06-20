package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	defer listener.Close()
	port := listener.Addr().(*net.TCPAddr).Port
	fmt.Printf("port found: %d\n", port)
}
