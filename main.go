package main

import (
	"fmt"

	"github.com/GoesToEleven/dog"
	"github.com/GoesToEleven/puppy"
)

func main() {
	fmt.Println(puppy.Barks())
	fmt.Println(dog.WhenGrownUp("Wof Wof!"))

	// fmt.Println("Greeting from Go")
	// fmt.Println(0o600)
	// fmt.Println(4_2)
	// fmt.Println(0xBad_Face)

	// fmt.Println(0.15e+0_2)
	// fmt.Println(0.15e-0_2)

	// var a = 1.23i
	// fmt.Println(a)

	// var b = 22.25i
	// fmt.Println(b)

	// var c = 1 + 2i
	// fmt.Println(c)

	//rune literals
	// var d = '\a'
	// fmt.Println(d)

	// var e = '\022'
	// fmt.Println(e)

	// var f = "\U000065e5\U0000672c\U00008a9e"
	// fmt.Println(f)

	fmt.Printf("%d\t %b\n", 10, 10)
	fmt.Printf("%d\t %b\n", 1<<1, 1<<1)
	fmt.Printf("%d\t %b\n", 1<<2, 1<<2)
}
