package main

// Importing basic packages
import (
	"fmt"

	"net"
	"os"
)

func main() {

	// Creating a listener which uses tcp protocol and listens on port 6379
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	} else {
		fmt.Println("Listening on port 6379")
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Accepted connection: ")
	}
	defer conn.Close()
	//buf:= make([] byte, 1024)
}
