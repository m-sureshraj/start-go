package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(filename string) (*os.File, error) {
	fmt.Println("Opening:", filename)
	return os.Open(filename)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing:", file.Name())
	file.Close()
}

func getFloats(fileName string) ([]float64, error) {
	file, err := OpenFile(fileName)
	var numbers []float64

	if err != nil {
		return nil, err
	}
	// defere function ensures a function call takes place even if the calling
	// function exists early
	defer CloseFile(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	// scanner panicked while reading the file
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, err
}

func main() {
	numbers, err := getFloats(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(numbers)
}
