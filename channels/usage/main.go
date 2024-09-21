package main

import "fmt"

func main() {

	c := make(chan int)

	// send-only
	go send(c)

	// recieve-only
	receive(c)
}

func send(c chan<- int) {
	c <- 42
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
