package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 42 // channel block

	fmt.Println(<-c) // fatal error: all goroutines are asleep - deadlock!
}
