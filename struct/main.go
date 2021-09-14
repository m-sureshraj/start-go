package main

import (
	"fmt"
	"strings"

	"start-go/types"
)

// defined types (custom types). Uses `struct` as an underlaying datatype
// Basically an Object in JS
type Employee struct {
	name    string
	id      int
	country string
	height  float64
}

func printStudentInfo(student types.Student) {
	// student param is a copy of the argument. (passed by value)
	// Hence, updating the copy won't update the original argument
	// student.name = "bingo"
	fmt.Println("Student Details:")
	fmt.Println("Name: ", student.Name)
	fmt.Println("Age: ", student.Age)
	fmt.Println("Country: ", student.Country)
	fmt.Println("Height: ", student.Height)
	fmt.Println("Street: ", student.Address.Street)
	fmt.Println("City: ", student.Address.City)
	fmt.Println("State: ", student.Address.State)
	fmt.Println("Postal Code: ", student.Address.PostalCode)
}

func makeNameToUpperCase(student *types.Student) {
	student.Name = strings.ToUpper(student.Name)
}

func getFakeEmployee(name string) Employee {
	employee := Employee{
		name:    name,
		id:      100,
		country: "nl",
		height:  5.8,
	}

	return employee
}

func main() {
	student := types.Student{
		Name:    "suresh",
		Age:     29,
		Country: "india",
		Height:  5.5,
		// assigning values to nested struct
		Address: types.Address{
			Street:     "foo bar",
			City:       "baz",
			State:      "ny",
			PostalCode: 2345,
		},
	}

	makeNameToUpperCase(&student)
	printStudentInfo(student)

	// using struct in slice
	employees := []Employee{}
	employees = append(employees, getFakeEmployee("foo"))

	fmt.Println(employees)
}
