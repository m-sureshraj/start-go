package main

import (
	"fmt"
	"strings"
)

// defined types can use any types as an underlaying type, although
// structs are most commonly used
type person struct {
	firstName string
	lastName  string
	age       int
}

// adding a method to the `person` type

// the initial parameter is called "receiver parameter"
// p (name of the parameter), person (type)
// the method becomes associated with all values of that type.
// in other languages this concept known as `this`

// method's name should be capitalized to be exported
func (p person) getFullName() string {
	return p.firstName + " " + p.lastName
}

// by default, the receiver parameter receives a copy of the receiver value
func (p *person) upperCaseFirstName() {
	p.firstName = strings.ToUpper(p.firstName)
}

func main() {
	foo := person{
		firstName: "bar",
		lastName:  "raj",
		age:       30,
	}

	fmt.Println(foo.getFullName())
	foo.upperCaseFirstName()
	fmt.Println(foo.getFullName())
}
