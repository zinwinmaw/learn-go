package main

import "fmt"

func main() {
	c := make(chan int)

	// c <- 42 // it is blocked

	// 1st method, since go routnine run, channel is no longer blocked
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	// 2nd method, using buffer
	bc := make(chan int, 2) // one buffer for one channel value
	bc <- 43
	bc <- 45

	fmt.Println(<-bc)
	fmt.Println(<-bc)
}
