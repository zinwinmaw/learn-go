package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	/// For loop ///
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// init and post are optional. Only expression is needed.
	b := 30
	a := 2
	for a < b {
		a *= 2
		fmt.Println(a)
	}

}
