package main

import (
	"fmt"
	"strings"
)

// type definition
// must accept "single string" argument and return a "string"
type StringProcessor func(string) string

// function signature
// accepts a single argument of type "string" and return a "string"
func toUpperCase(s string) string {
	return strings.ToUpper(s)
}

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

// for the function to match "StringProcesser" type signature, it must have the same input and output types as defined by "StringProcessor"
// this is Polymorphic Behavior. This function doesnt need to know the specific implimentation ( toUpperCase, toLowerCase ). It only requires a function that matches the "StringProcessor" type.
func processString(s string, callback StringProcessor) {
	result := callback(s)
	fmt.Println("Processed string - ", result)
}

func main() {
	str := "heLLO worlD"
	processString(str, toUpperCase)
	processString(str, toLowerCase)
}
