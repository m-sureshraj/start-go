package types

import "fmt"

type Date struct {
	year  int
	month int
	day   int
}

// note: By convention, a getter method’s name should be the same
// as the name of the field or variable it accesses.

// Getter methods don’t need to modify the receiver at all,
// so we could use a direct Date value as a receiver.
// But if any method on a type takes a pointer receiver, convention says
// that they all should, for consistency’s sake.
func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return fmt.Errorf("Invalid year! Must be greater than 1, but received %d", year)
	}

	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return fmt.Errorf("Invalid month! Must be between 1 to 12, but received %d", month)
	}

	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return fmt.Errorf("Invalid day! Must be between 1 to 31, but received %d", day)
	}

	d.day = day
	return nil
}
