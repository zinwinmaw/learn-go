package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	name string
}

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.name))
}

func main() {

	p := person{
		"James Doe",
	}

	// file has Write() function - so, it is also writer interface type
	f, err := os.Create("output2.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	// byte buffer has Write() function - so, it is also writer interface type
	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String())

}
