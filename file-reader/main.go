package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFloats(fileName string) ([]float64, error) {
	file, err := os.Open(fileName)
	var numbers []float64

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	// scanner panicked while reading the file
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, err
}

// how to run? install the program. copy the height.txt to the $GOPATH/bin/ dir
func main() {
	numbers, err := getFloats("heights.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(numbers)
}
