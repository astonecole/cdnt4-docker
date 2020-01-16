package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

var addr = flag.String("addr", "0.0.0.0:3000", "Server host")

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
	server, err := net.Listen("tcp", *addr)
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
			go c.Broadcast(conn)
		case conn := <-c.leave:
			fmt.Println("SYSTEM: Bye")
			delete(c.users, conn)
		}
	}
}

func (c *Chat) Broadcast(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Print(err)
			break
		}

		fmt.Print(conn.RemoteAddr().String() + ": " + message)

		for connection := range c.users {
			if connection != conn {
				io.WriteString(connection, message)
			}
		}
	}

	c.leave <- conn
}

func main() {
	flag.Parse()

	fmt.Println("Server is running on " + *addr)

	chat := NewChat()
	go chat.Handle()
	chat.Serve()
}
