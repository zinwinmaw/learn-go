package main

import "fmt"

// "makeCounter" is the outer function that defines a closure.
func makeCounter() func() int {

	// "count" is a variable defined in the outer function
	count := 0

	// The anonymous return function is the "Closure"
	// It captures and retains the environment (eg, count variable) from the outer function
	return func() int {
		count++
		return count
	}
}

func main() {
	counter := makeCounter()

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

}
