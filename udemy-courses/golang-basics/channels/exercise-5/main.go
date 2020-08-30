package main

import "fmt"

func main() {
	ch := make(chan int)

	go func(c chan<- int) {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}(ch)

	// Trying out the comma ok idiom.
	for {
		if v, ok := <-ch; ok {
			fmt.Println(v, ok)
		} else {
			break
		}
	}

}
