package main

import (
	"fmt"
	"log"
	"net"
)

type Chat struct {
	users map[net.Conn]bool
	join  chan net.Conn
	leave chan net.Conn
}

func NewChat() *Chat {
	return &Chat{
		users: make(map[net.Conn]bool),
		join:  make(chan net.Conn),
		leave: make(chan net.Conn),
	}
}

func (c *Chat) Serve() {
	server, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			c.users[conn] = true
			c.join <- conn
		}()
	}
}

func (c *Chat) Handle() {
	for {
		select {
		case conn := <-c.join:
			fmt.Println("SYSTEM: Welcome")
		case conn := <-c.leave:
			fmt.Println("SYSTEM: Bye")
			delete(c.users, conn)
		}
	}
}

func main() {

}
