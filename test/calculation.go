package main

import "fmt"

type mathCalculation func(int, int) int

// Add function receive two integer values
// and return the total result
func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

// doMath function do some math functions
func DoMath(a int, b int, calc mathCalculation) {
	r := calc(a, b)
	fmt.Println("Result is : ", r)
}

func main() {
	DoMath(5, 6, Add)
	DoMath(20, 30, Subtract)
}

// To show all documentations in command line -
// go doc -cmd -all
// go doc -cmd -u
