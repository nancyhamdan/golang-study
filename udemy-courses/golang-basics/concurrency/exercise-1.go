package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("This is where the program starts. No goroutines started yet.")

	var wg sync.WaitGroup

	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			fmt.Println("hello! I'm a goroutine!")
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Goroutines are done. Program is done.")
}
