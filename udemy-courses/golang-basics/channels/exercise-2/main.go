package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)

	go func() {
		cs <- 15
	}()

	// this would not work if cs is a send only channel.
	fmt.Println(<-cs)
}
