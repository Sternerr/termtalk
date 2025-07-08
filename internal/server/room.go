package server

import (
	"net"
	"fmt"
	"bytes"
)

type Room struct {
	Users []User
}

func(r *Room) Add(c net.Conn) {
	user := User{
		Conn: c,
	}
	r.Users = append(r.Users, user)
}


func(r *Room) Remove(c net.Conn) {
	for i := 0; i < len(r.Users); i += 1 {
		if r.Users[i].Conn == c {
			r.Users = append(r.Users[:i], r.Users[i+1:]...)
		}
	}
}

func(r *Room) Send(from net.Conn, messageBuffer []byte) {
	for i := 0; i < len(r.Users); i += 1 {
		if r.Users[i].Conn == from {
			continue
		}

		sender := from.RemoteAddr().String()
		trimmedMessageBuffer := bytes.TrimRight(messageBuffer, "\r\n")
		msg := append([]byte(sender + ": "), trimmedMessageBuffer...)
		msg = append(msg, '\n')

		_, err := r.Users[i].Conn.Write(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func(r *Room) BroadcastSystemMessage(messageBuffer []byte) {
	trimmedMessageBuffer := bytes.TrimRight(messageBuffer, "\r\n")
	msg := append([]byte("[server]: "), trimmedMessageBuffer...)
	msg = append(msg, '\n')

	for i := 0; i < len(r.Users); i += 1 {
		_, err := r.Users[i].Conn.Write(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

