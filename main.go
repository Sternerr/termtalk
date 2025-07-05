package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Chat Server")

	addr := "localhost:42069"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Client Connected, %s", conn)
	for {}
}
