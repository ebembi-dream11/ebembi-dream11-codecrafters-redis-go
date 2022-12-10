package main

import (
	"fmt"
	"net"
)

func main() {
	// Create a listener on localhost:8080
	ln, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	// Accept incoming connections
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Send a ping message
	_, err = conn.Write([]byte("ping"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Wait for a response
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the response
	fmt.Println("This is my response", string(buf[:n]))
	//fmt.Println()
}
