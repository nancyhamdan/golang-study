package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)

	go func() {
		c1 <- 15
	}()

	fmt.Println(<-c1)

	// This is a buffered channel.
	c2 := make(chan int, 1)

	c2 <- 5

	fmt.Println(<-c2)
}
