package main

import "fmt"

type Dog struct{}

func (d Dog) Bark() { // Value receiver
	fmt.Println("Woof!")
}

func (d *Dog) Wag() { // Pointer receiver
	fmt.Println("Wagging tail!")
}

func main() {
	var d Dog
	d.Bark() // Valid, calls method with value receiver

	var pd *Dog = &d
	pd.Wag() // Valid, calls method with pointer receiver

	// pd.Bark() would work as well due to automatic conversion
}
