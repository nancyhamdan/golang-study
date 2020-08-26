package main

import (
	"fmt"
)

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello! I'm a person.")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{"Nancy"}

	saySomething(&p1)

	p1.speak()
}
