package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // need to close channel, otherwise deadlock coz "receive" is waiting
	}()

	// receive
	// range loop is hanging there until the channel is closed
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("About to exit")
}
