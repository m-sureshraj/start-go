package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// It's a variadic function. A function that accepts varying number of
// arguments. Here the `numbers` argument is `slice` data type.
// In JS it's called REST param.
func average(numbers ...float64) float64 {
	var total float64

	for _, number := range numbers {
		total += number
	}

	length := len(numbers)
	if length == 0 {
		return 0
	}

	return total / float64(len(numbers))
}

func main() {
	// slice
	var numbers []float64

	// `os.Args' first argument is the executable name. We omitted it via
	// slice operator
	for _, number := range os.Args[1:] {
		number, err := strconv.ParseFloat(number, 64)

		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, number)
	}

	// spreading
	fmt.Println(average(numbers...))
}
