package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int

	var wg sync.WaitGroup
	wg.Add(10)

	var mu sync.Mutex

	fmt.Println("Before goroutines: counter =", counter)

	for i := 0; i < 10; i++ {
		go func() {
			x := counter
			mu.Lock()
			x++
			counter = x
			fmt.Println("counter =", counter)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("After goroutines: counter =", counter)
	fmt.Println("Program is done")
}
