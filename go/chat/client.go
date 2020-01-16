package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var addr = flag.String("addr", ":3000", "Server host")

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Dial is ok")

	go io.Copy(conn, os.Stdin)
	io.Copy(os.Stdout, conn)
}
