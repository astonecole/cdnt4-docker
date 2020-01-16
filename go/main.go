package main

import (
	"fmt"
)

type Runner interface {
	Run()
}

type Voiture struct {
	marque string
}

func (v Voiture) Run() {
	fmt.Println(v.marque + " I'm running")
}

func startApp(r Runner) {
	r.Run()
}

func main() {
	v := Voiture{marque: "Peugeot"}
	startApp(v)
}
