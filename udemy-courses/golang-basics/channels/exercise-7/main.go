package main

import (
	"fmt"
	"sync"
)

func main() {
	maxGoroutines := 10
	ch := make(chan int)
	doneCh := make(chan struct{})

	go putCh(ch, maxGoroutines, doneCh)

	pullCh(ch, doneCh)

	fmt.Println("About to exit")
}

func putCh(ch chan<- int, maxGoroutines int, doneCh chan<- struct{}) {
	var wg sync.WaitGroup

	for i := 0; i < maxGoroutines; i++ {
		wg.Add(1)
		go func(c chan<- int, i int) {
			//fmt.Println("goroutine:", i)
			for i := 1; i <= 5; i++ {
				c <- i
			}
			wg.Done()
		}(ch, i)
	}

	wg.Wait()
	doneCh <- struct{}{}
}

func pullCh(ch <-chan int, doneCh <-chan struct{}) {
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-doneCh:
			return
		}
	}
}
