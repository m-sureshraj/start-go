package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getStrings(fileName string) ([]string, error) {
	var lines []string

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = strings.ToLower(strings.TrimSpace(scanner.Text()))

		if line != "" {
			lines = append(lines, line)
		}
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return lines, nil
}

// map literals
// ranks := map[string]int{"gold": 4}

func main() {
	names, err := getStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	voteCount := map[string]int{}

	for _, name := range names {
		// accessing the map[key] returns an optional second bool type value.
		// if the key exist in the map, the value will be true else false.

		// if _, ok := voteCount[name]; ok {
		// 	voteCount[name] += 1
		// } else {
		// 	voteCount[name] = 1
		// }

		// As with arrays and slices, zero values can make it safe
		// to manipulate a map value even if you haven’t explicitly
		// assigned to it yet.
		voteCount[name]++
	}

	for key, value := range voteCount {
		fmt.Printf("%s: %d\n", key, value)
	}
}
