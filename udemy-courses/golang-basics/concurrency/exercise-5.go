package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64

	var wg sync.WaitGroup
	wg.Add(10)

	fmt.Println("Before goroutines: counter =", counter)

	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter =", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("After goroutines: counter =", counter)
	fmt.Println("Program is done")
}
