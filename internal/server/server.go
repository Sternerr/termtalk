package server

import (
	"fmt"
	"net"
	"os"
)
	
type Server struct {
	Listener net.Listener
	Room Room
}

func(s *Server) Listen() {
	buffer := make([]byte, 1024)

	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		
		s.Room.Add(conn)

		msg := fmt.Sprintf("%s Connected", conn.RemoteAddr())
		s.Room.BroadcastSystemMessage([]byte(msg))
	
		go func() {
			for {
				numBytesRead, err := conn.Read(buffer)
				if err != nil {
					fmt.Println(err.Error())
				}

				s.Room.Send(conn, buffer[:numBytesRead])

				if numBytesRead == 0 {
					break
				}
			}

			s.Room.Remove(conn)			
			msg := fmt.Sprintf("%s Disconnected", conn.RemoteAddr())
			s.Room.BroadcastSystemMessage([]byte(msg))

			conn.Close()
		}()
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
		Room: Room{},
	}

	return &server
}
