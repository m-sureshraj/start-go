package types

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	// embedding the Date field from date.go
	// un-exported fields don't get included to the Event
	Date
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 25 {
		return errors.New("title must be less than 25 characters")
	}

	e.title = title
	return nil
}
