package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// to test defer function while os.Exit
	defer foo()

	// log file create first
	f, err2 := os.Create("log.txt")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("names.txt")
	if err != nil {
		// output error log in console
		fmt.Println("Errors:", err)

		// out error log in console with Timestamp if there is no log file
		// if log.txt file exist, it will output logs in the log file
		// log.Println("Errors:", err)

		// Fatal function call os.Exit(1) after writing the error meesage log
		// it won't even run defer function and total exit the program
		// log.Fatalln(err)

		// defer function runs, then exist status 2
		// can use recover
		log.Panicln(err) // log the error message using Go log package including timestamp
		panic(err)       // just print the messages
	}
	defer f2.Close()

}

func foo() {
	fmt.Println("It is deferred function. When os.Exit")
}
