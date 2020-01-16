package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("My Function")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("My Function 2")
		wg.Done()
	}()

	fmt.Println("Fin.")
	wg.Wait()
}
