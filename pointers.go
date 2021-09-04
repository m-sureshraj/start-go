package main

import "fmt"

func increase(age *int) {
	*age += 10
}

// In Go, & is called "address of" operator

// it’s okay to return a pointer to a variable that’s local to a function.
// Even though that variable is no longer in scope, as long as you still have the
// pointer, Go will ensure you can still access the value. In JS this concept
// called as Closures
func giveMeAPointer() *int {
	age := 100
	return &age
}

func main() {
	height := 4.5
	heightRef := &height

	fmt.Println("Referent of the height: ", heightRef)
	fmt.Println("Actual value of the height var: ", *heightRef)

	someIntPointer := giveMeAPointer()
	fmt.Println(*someIntPointer)
}
