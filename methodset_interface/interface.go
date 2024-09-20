// Interface Implementations
package main

import "fmt"

// Define an interface
type Animal interface {
	Speak() string
}

// Puppy implements the Animal interface
type Puppy struct{}

func (d Puppy) Speak() string {
	return "Woof!"
}

// Cat implements the Animal interface
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// Define an interface
type Speaker interface {
	Speak() string
}

// Type with pointer receiver
type WildCat struct {
	name string
}

func (c *WildCat) Speak() string {
	return "Meow! My name is " + c.name
}

func makeSound(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	var d Animal = Puppy{}
	var c Animal = Cat{}

	makeSound(d) // Output: Woof!
	makeSound(c) // Output: Meow!

	cat := &WildCat{name: "Whiskers"} // cat is a pointer to Cat

	// Cat satisfies Speaker interface via pointer receiver
	var speaker Speaker = cat
	fmt.Println(speaker.Speak()) // Output: Meow! My name is Whiskers

}
