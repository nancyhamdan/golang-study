package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int

	var wg sync.WaitGroup
	wg.Add(10)

	fmt.Println("Before goroutines: counter =", counter)

	for i := 0; i < 10; i++ {
		go func() {
			x := counter
			runtime.Gosched()
			x++
			counter = x
			fmt.Println("counter =", counter)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("After goroutines: counter =", counter)
	fmt.Println("Program is done")
}
