package main

import "fmt"

type teacher struct {
	Name string
	Age  int
}

type pilot struct {
	Name    string
	Age     int
	Service int
}

// func (t teacher) speak() string {
// 	return fmt.Sprintf("Name: %s, Age: %d\n", t.Name, t.Age)
// }

func (t teacher) speak() {
	fmt.Printf("Name: %s, Age: %d\n", t.Name, t.Age)
}

// func (p pilot) speak() string {
// 	return fmt.Sprintf("I am Pilot %s, service %d years and age %d.", p.Name, p.Service, p.Age)
// }

func (p pilot) speak() {
	fmt.Printf("I am Pilot %s, service %d years and age %d.", p.Name, p.Service, p.Age)
}

// to use polymorphism. Any type that has speak() method is a "human" type
// type human interface {
// 	speak() string
// }

type human interface {
	speak()
}

// polymorphic, can pass anytype of human. Here, teacher type and pilot type

// func SaySomthing(h human) {
// 	fmt.Println(h.speak())
// }

func SaySomthing(h human) {
	h.speak()
}

func main() {
	t := teacher{"James", 20}
	p := pilot{"Susan", 25, 3}

	// t.speak()
	// p.speak()

	// Illustrating Polymorphism. Value could be more than one type.
	SaySomthing(t)
	SaySomthing(p)
}
