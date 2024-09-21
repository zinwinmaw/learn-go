package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	merged := make(chan int)

	go calculate(ch1, 1)
	go calculate(ch2, 2)

	go fanIn(ch1, ch2, merged)

	for i := 0; i < 5; i++ {
		fmt.Println(<-merged)
	}
	close(merged)
	fmt.Println("about to exit")
}

func calculate(ch chan<- int, id int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		ch <- id*10 + i
	}
	close(ch)
}

func fanIn(c1, c2 <-chan int, mg chan<- int) {
	for {
		select {
		case v, ok := <-c1:
			if ok {
				mg <- v
			}
		case v, ok := <-c2:
			if ok {
				mg <- v
			}
		default:
			// check both channels are closed or not to avoid data race condition
			// both channels are closed. it will exit the loop
			if !isChannelOpen(c1) && !isChannelOpen(c2) {
				return
			}
		}
	}
}

func isChannelOpen(ch <-chan int) bool {
	select {
	case _, ok := <-ch:
		return ok // is no value, it will return false
	default:
		return true // assume channel is still open
	}
}
