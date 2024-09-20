package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func workers(id int) {
	defer wg.Done()
	fmt.Printf("Worker %d is starting\n", id)
	time.Sleep(2 * time.Second) // Simulate work
	fmt.Printf("Worker %d is done\n", id)
}

func main() {

	const numWorkers = 3
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go workers(i)
	}
	wg.Wait()
	fmt.Println("All workers are completed")
}
