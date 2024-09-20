package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64

	const gs = 10
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			// runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))

			wg.Done()
		}()
		fmt.Println("Goroutines loop:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Counters:", counter)
}
