package main

import (
	"fmt"
)

func main() {
	doneCh := make(chan struct{})
	ch1 := gen(doneCh)

	receive(ch1, doneCh)

	fmt.Println("About to exit")
}

func gen(doneCh chan<- struct{}) chan int {
	c := make(chan int)

	go func(ch chan<- int) {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		fmt.Println("goroutine about to finish")
		close(ch)
		doneCh <- struct{}{}
	}(c)
	return c
}

func receive(ch1 <-chan int, doneCh <-chan struct{}) {
	for {
		select {
		case v := <-ch1:
			fmt.Println(v)
		case <-doneCh:
			return
		}
	}
}
