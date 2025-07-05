package main

import (
	"fmt"
	"net"
	"os"
)

type Room struct {
	users []net.Conn
}

func(r *Room) add(c net.Conn) {
	r.users = append(r.users, c)
}


func(r *Room) remove(c net.Conn) {
	for i := 0; i < len(r.users); i += 1 {
		if r.users[i] == c {
			r.users = append(r.users[:i], r.users[i+1:]...)
		}
	}
}

func(r *Room) send(from net.Conn, messageBuffer []byte) {
	for i := 0; i < len(r.users); i += 1 {
		if r.users[i] == from {
			continue
		}

		sender := from.RemoteAddr().String()
		msg := append([]byte(sender + ": "), messageBuffer...)
		_, err := r.users[i].Write(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func main() {
	fmt.Println("Chat Server")

	addr := "localhost:42069"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer l.Close()

	room := Room {}
	buffer := make([]byte, 1024)

	for {
		conn, err := l.Accept()
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
