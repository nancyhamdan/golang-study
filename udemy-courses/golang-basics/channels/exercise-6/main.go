package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go fillCh(ch)
	printCh(ch)

	fmt.Println("About to exit program.")
}

func fillCh(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func printCh(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
