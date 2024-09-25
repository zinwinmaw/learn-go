package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// create file
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// write to file
	s := strings.NewReader("James Bond")
	io.Copy(f, s)

	// open file
	f, err = os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// read contents
	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}
