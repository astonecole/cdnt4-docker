package main

import (
	"fmt"
	"time"
)

func greeting(msg chan string) {
	time.Sleep(time.Second * 2)
	msg <- "Bravo le message est passé"
}

func main() {
	message := make(chan string)
	go greeting(message)

	fmt.Println("Hello")
	fmt.Println(<-message)
}
