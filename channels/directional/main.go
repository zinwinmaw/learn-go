package main

import "fmt"

func main() {
	c := make(chan int)    // Bidirectional
	cs := make(chan<- int) // Send-only to channel
	cr := make(<-chan int) // receive-only from channel

	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cs\t%T\n", cs)
	fmt.Printf("cr\t%T\n", cr)

	// General to Specific assgins and can be converted
	// cs = c
	// cr = c
	// type conversion bidirectional to (chan<- int) send-only channel
	fmt.Printf("cs\t%T\n", (chan<- int)(c))
	fmt.Printf("cr\t%T\n", (<-chan int)(c))

	// Specific to General can't be converted
	// c = cr
	// c = cs
	fmt.Printf("c\t%T\n", (chan int)(cs))
	fmt.Printf("c\t%T\n", (chan int)(cr))
}
