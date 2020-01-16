package main

import "fmt"

const version = "0.0.0"

func main() {
	// name, age := "toto", 30
	// str := "a" + "b"
	// fmt.Println(name, age, version)
	// fmt.Println(greeting("Roger"))

	// x, y := 5, 10
	// x, y = swap(x, y)
	// fmt.Println(x, y)

	a := 1
	b := &a
	*b = 20

	fmt.Println(*b)
}

func swap(a, b int) (int, int) {
	return b, a
}

func greeting(name string) string {
	return "Hello " + name
}
