package main

import (
	"fmt"
	"reflect"
	"strings"
)

// https://stackoverflow.com/questions/39092925/why-are-interfaces-needed-in-golang
// The interface contains only the signatures of the methods, but not their implementation.
type MyInterface interface {
	toUpperCase() string
	toLowerCase() string
}

// To satisfy the above interface, following type must have
// all methods of the interface. Though, the type can have methods
// which are not listed in the interface.
type MyName string

func (m MyName) toUpperCase() string {
	return strings.ToUpper(string(m))
}

func (m MyName) toLowerCase() string {
	return strings.ToLower(string(m))
}

// this is an extra method
func (m MyName) toTitle() string {
	return strings.Title(string(m))
}

// pointer example
type Switch bool

func (s *Switch) Toggle() {
	*s = !*s
}

type Bulb interface {
	Toggle()
}

// another example
type Animal interface {
	Say() string
}

type Cat struct{}

func (c Cat) Say() string {
	return "Meow"
}

type Dog struct{}

func (d Dog) Say() string {
	return "woof"
}

type Parrot struct{}

func (d Parrot) Say() string {
	return "ki ki"
}

func (d Parrot) Sing() string {
	return "ki ko ki ko ku u"
}

// It's called empty interface. it matches all the types
type anthing interface{}

// this param type can be anything
func customPrint(s interface{}) {
	fmt.Println(s)
}

func customVariadicPrint(s ...interface{}) {
	fmt.Println(s...)
}

func main() {
	// When we have a variable with an interface type,
	// it can hold values of any type that satisfies the interface.
	var value MyInterface = MyName("suresh")

	// we can only call methods defined as part of the interface
	fmt.Println(value.toUpperCase())

	s1 := Switch(true)
	var bulb Bulb = &s1

	bulb.Toggle()
	bulb.Toggle()
	bulb.Toggle()
	fmt.Println(s1)

	cat := Cat{}
	dog := Dog{}
	parrot := Parrot{}

	animals := []Animal{cat, dog, parrot}

	for _, animal := range animals {
		fmt.Println(reflect.TypeOf(animal).Name(), "says:", animal.Say())

		// type assertion
		p, ok := animal.(Parrot)
		if ok {
			fmt.Println("I'm a parrot and I can sing as well", p.Sing())
		}
	}

	customPrint("ss")
	customVariadicPrint(10, "sure", true, 30.54)
}
