package main

import (
	"fmt"
	"log"
	"start-go/types"
)

func main() {
	event := types.Event{}

	err := event.SetTitle("My birthday")
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetYear(1990)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event)
}
