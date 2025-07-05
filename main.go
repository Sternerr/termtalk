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

	buffer := make([]byte, 1024)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Client Connected, %s", conn)
		for {
			numBytesRead, err := conn.Read(buffer)
			if err != nil {
				fmt.Println(err.Error())
			}

			if numBytesRead == 0 {
				break
			}

			_, err = conn.Write(buffer[:numBytesRead])
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		
		conn.Close()
		fmt.Println("Client Disconnected, %s", conn)
	}
}
