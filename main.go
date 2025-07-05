package main

import (
	"fmt"
	"os"

	"github.com/sternerr/termtalk/internal/server"
	"github.com/sternerr/termtalk/internal/client"
)


func main() {
	fmt.Println("Chat Server")
	
	args := os.Args
	if len(args) < 2 {
		fmt.Println("must atleast 2 arguments")
		os.Exit(0)
	}
	
	mode := args[1]
	switch mode {
	case "host":
		port := parseFlag(args[2:], "--port")
		if port == "" {
			fmt.Println("Must provide --port for host mode")
			os.Exit(1)
		}

		server := server.NewServer(port)
		server.Listen()
	case "client":
		port := parseFlag(args[2:], "--port")
		client := client.NewClient("localhost", port)
		client.StartClient()
	default:
		fmt.Printf("%s is not a command\n", mode)
		os.Exit(0)
	}
}

func parseFlag(args []string, name string) string {
	for i := 0; i < len(args); i += 1 {
		if args[i] == name && i+1 < len(args) {
			return args[i+1]
		}
	}

	return ""
}
