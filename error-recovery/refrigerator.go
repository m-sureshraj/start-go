package main

import (
	"fmt"
)

type Refrigerator []string

func (r Refrigerator) Open() {
	fmt.Println("Opening the door")
}

func (r Refrigerator) Close() {
	fmt.Println("closing the door")
}

func (r Refrigerator) FindFood(food string) error {
	r.Open()
	defer r.Close()

	if find(r, food) {
		fmt.Println("Found:", food)
	} else {
		return fmt.Errorf("%s not found", food)
	}

	return nil
}

func find(items []string, value string) bool {
	for _, item := range items {
		if item == value {
			return true
		}
	}

	return false
}
