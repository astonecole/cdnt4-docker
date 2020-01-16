package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u User) greeting() {
	fmt.Println(u.name, u.age)
}

type Admin struct {
	User
}

func (a Admin) superGreeting() {
	fmt.Println("Je suis un super admin " + a.name)
}

func main() {
	u := User{
		name: "toto",
		age:  30,
	}
	a := Admin{User: u}

	a.greeting()
}
