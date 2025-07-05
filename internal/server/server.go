package server

import (
	"fmt"
	"net"
	"os"
)
	
type Server struct {
	Listener net.Listener
}

func(s *Server) Listen() {
	room := Room {}
	buffer := make([]byte, 1024)

	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		
		fmt.Println("Client Connected, %s", conn)
		room.add(conn)
	
		go func() {
			for {
				numBytesRead, err := conn.Read(buffer)
				if err != nil {
					fmt.Println(err.Error())
				}

				room.send(conn, buffer[:numBytesRead])

				if numBytesRead == 0 {
					break
				}
			}

			room.remove(conn)			
			conn.Close()
			fmt.Println("Client Disconnected, %s", conn)
		}()

		fmt.Println(room)
	}
}

func NewServer(port string) *Server {
	addr := fmt.Sprintf(":%s", port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	server := Server {
		Listener: l,
	}

	return &server
}
