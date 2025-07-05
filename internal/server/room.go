package server

import (
	"net"
	"fmt"
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

