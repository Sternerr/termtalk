package client

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

type Client struct {
	Conn net.Conn
}

func(c *Client) StartClient() {
	go func() {
		scanner := bufio.NewScanner(c.Conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	input := bufio.NewScanner(os.Stdin)
	for {
		if !input.Scan() {
			break
		}

		text := input.Text()
		c.Conn.Write([]byte(text))
	}
}

func NewClient(host, port string) *Client {
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	
	client := Client{
		Conn: conn,
	}

	return &client
}


