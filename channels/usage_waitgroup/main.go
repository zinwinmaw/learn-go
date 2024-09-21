package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	c := make(chan int)

	wg.Add(2)

	// send-only
	go func() {
		defer wg.Done()
		send(c)
	}()

	// recieve-only
	go func() {
		defer wg.Done()
		receive(c)
	}()
	wg.Wait()
}

func send(c chan<- int) {
	c <- 42
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
