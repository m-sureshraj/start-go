package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func generateRandomNumber() int {
	rand.Seed(time.Now().Unix())

	return rand.Intn(100)
}

func readUserInput() int {
	fmt.Print("Guess the number: ")
	reader := bufio.NewReader(os.Stdin)
	rawInput, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	rawInput = strings.TrimSpace(rawInput)
	guess, err := strconv.Atoi(rawInput)

	if err != nil {
		log.Fatal(err)
	}

	return guess
}

func main() {
	randomNumber := generateRandomNumber()

	maxAttempt := 10
	attempt := 0
	var input int
	hasGussed := false

	for attempt < maxAttempt {
		fmt.Println("you have", maxAttempt-attempt, "guesses left")
		attempt++
		input = readUserInput()

		if input < randomNumber {
			fmt.Println("Oops. Your guess was LOW\n")
		} else if input > randomNumber {
			fmt.Println("Oops. Your gues was HIGH\n")
		} else {
			hasGussed = true
			fmt.Println("Good Job! you guessed it!")
			break
		}
	}

	if !hasGussed {
		fmt.Println("Sorry. You didn't guess my number. It was: ", randomNumber)
	}
}
