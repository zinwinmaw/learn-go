package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

type count int

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title, "\n")
}

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

// wrapper function. Passing Stringer interface from "fmt" package. To distingush Stringer interface from other package.
func LogInfo(s fmt.Stringer) {
	log.Println("From log stringer: ", s.String())
}

func main() {

	var n count = 42
	b := book{title: "Gone with the wind"}

	// fmt.Println(b)
	// fmt.Println(n)

	// log.Println(b)
	LogInfo(b)
	LogInfo(n)
}
