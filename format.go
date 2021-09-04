package main

import "fmt"

func main() {
	// following percent sign %<> are called formatting verbs

	// string
	fmt.Printf("String: %s \n", "hello world")

	// decimal integer
	fmt.Printf("Decimal Integer: %d \n", 10)

	// float point number
	fmt.Printf("Float: %f \n", 10.567)

	// boolean type (true | false)
	fmt.Printf("Boolean: %t \n", true)

	// Could be any value. Chooses an approriate format based on the given value's type
	fmt.Printf("Values: %v \n", 4.5)

	// Could be any value. Formatted as it would appear in the Go code
	fmt.Printf("Values: %#v \n", "hello		")

	// Prints the given value's type
	fmt.Printf("Types: %T \n", 4.5)

	// control the width when formatting (padding)

	// total length should be 5 characters (including the decimal point)
	fmt.Printf("hello: %8.2f world \n", 3.908674934) // '    3.91'
	fmt.Printf("hello: %.2f world \n", 3.908674934)  // '3.91
}
