package main

import "fmt"

type InvalidCowError struct {
	cows    int
	details string
}

func (e *InvalidCowError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.details)
}

func main() {
	cow := -20
	err := ValidateNumberofCows(cow)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("No. of cows %d is valid", cow)

}

func ValidateNumberofCows(c int) error {
	if c < 0 {
		return &InvalidCowError{
			cows:    c,
			details: "there is no negative cows.",
		}
	} else if c == 0 {
		return &InvalidCowError{
			c,
			"o cows need the food.",
		}
	} else {
		return nil
	}
}
