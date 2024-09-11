package main

import (
	"log"
	"os"
)

func main() {

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	s := []byte("Hello Gophers! Nice to meet you all.")
	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}
}
