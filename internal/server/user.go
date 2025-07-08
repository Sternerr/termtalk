package server

import (
	"net"
)

type User struct {
	Conn net.Conn
}
