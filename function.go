package main

import (
	"fmt"
	"log"
)

// The parentheses around the return values are optional when there’s only one return
// value, but are required if there’s more than one return value.
func updateNumber(age int) (int, error) {
	minAge := 10
	if age < minAge {
		// no formating
		// return 0, errors.New("age should be greater than or equal to 10")

		// formats and returns a error value
		return 0, fmt.Errorf("Min age %d should be greater than or equal to minimum age %d", age, minAge)
	}

	return age + 10, nil
}

// Go is a pass-by-value language. Arguments are passed by value
// When the function is run, each parameter will be set to a copy of the value in the
// corresponding argument.

func main() {
	age := 10
	updatedAge, err := updateNumber(age)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Old age: %d, New age: %d \n", age, updatedAge)
}
