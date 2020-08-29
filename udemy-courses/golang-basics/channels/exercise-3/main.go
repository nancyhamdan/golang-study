package main

import (
	"fmt"
)

func main() {
	c := genChan()

	receive(c)

	fmt.Println("About to exit.")
}

func genChan() <-chan int {
	cr := make(chan int)
	go func(ch chan<- int) {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}(cr)
	return cr
}

func receive(cr <-chan int) {
	for v := range cr {
		fmt.Println(v)
	}
}
