package main

import "fmt"

func main() {
	x := 42

	fmt.Println("Value ", x)
	fmt.Println("Memory Address", &x)
	fmt.Printf("%v\t%T\n", &x, &x)

	s := "Hello"
	fmt.Printf("%v\t%T\n", &s, &s)

	y := &x //address of x
	fmt.Printf("Address of x: %v\t%T\n", y, y)
	fmt.Printf("Get Value: %v\t%T\n", *y, *y)

	*y = 43 // dereferencing the pointer
	fmt.Printf("New value of Address: %v\n", x)
	fmt.Println(x)
	fmt.Println(*y)
}
