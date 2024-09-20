package main

import (
	"fmt"
)

// func add(a, b int) int {
// 	return a + b
// }

// func add(a, b float64) float64 {
// 	return a + b
// }

// Type Constraints using generic function
func addT[T int | float64](a, b T) T {
	return a + b
}

// Type Set using Interface
type myNumbers interface {
	~int | ~float64 // "~" type integer and underlying type int ( like "myAlias" down below )
}

// is not working
// type myNumbers interface {
// 	constraints.Integer | constraints.Float
// }

func addI[T myNumbers](a, b T) T {
	return a + b
}

// Type Alias and Underlying Type Constraints
type myAlias int

func main() {

	var n myAlias = 42

	fmt.Println(addT(3, 5.5))
	fmt.Println(addI(3, 5.5))
	// fmt.Println(addI(n, 5.5)) - cannot use
	fmt.Println(addI(n, 6))
}

// Concrete Type
// A concrete type is a specific, defined type in Go, such as int, float64, string, or a user-defined struct. It has a defined implementation and can be instantiated directly.

// Interface Type
// An interface type in Go specifies a set of methods that a type must implement. It does not define how the methods are implemented; it only specifies what methods are expected. Types that implement the interface are said to "satisfy" the interface.
// A type that specifies a set of methods. It does not provide implementations but specifies what methods types must have to satisfy the interface.

/**
*
// Define an interface (contract)
type Greeter interface {
    Greet() string
}

// Define a concrete type that satisfies the interface
type Person struct {
    Name string
}

// Implement the method required by the interface
func (p Person) Greet() string {
    return "Hello, my name is " + p.Name
}

func main() {
    var g Greeter

    // Create an instance of Person
    p := Person{Name: "Alice"}

    // The Person type satisfies the Greeter interface
    g = p

    fmt.Println(g.Greet()) // Output: Hello, my name is Alice
}
	**/
