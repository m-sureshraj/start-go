package types

// As with other declarations, type name must be capitalized to be exported.
// the above rule is true for the stuct's fields as well. Only the capitalized
// names will be expored
type Student struct {
	Name    string
	Age     int
	Country string
	Height  float64
	// embedding struct allows us to add all the fields from the inner struct
	// we could access it via:
	// Student.Address.Street
	// or Student.Street
	Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode int
}
