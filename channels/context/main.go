package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	fmt.Println("Context:\t", ctx)
	fmt.Println("Context Err:\t", ctx.Err())
	fmt.Printf("Context Type:\t%T\n", ctx)
	fmt.Println("------")

	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// Simulate some work
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Finished work before timeout")
	case <-ctx.Done():
		fmt.Println("Timeout or cancellation happened:", ctx.Err())
	}
}
