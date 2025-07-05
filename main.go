package main

import (
	"fmt"

	"github.com/sternerr/simple-chat-server/internal/server"
)


func main() {
	fmt.Println("Chat Server")
	server := server.NewServer(6969)
	server.Listen()
}
