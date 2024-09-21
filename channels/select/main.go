package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)

	fmt.Println("about to exit")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the eve channel:", v)
		case v := <-o:
			fmt.Println("from the odd channel:", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from comman ok", i, ok)
				return
			} else {
				fmt.Println("from comma ok", i)
			}
			// fmt.Println("from the quit channel:", v)
			// return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	// q <- 0
	close(q)
}
