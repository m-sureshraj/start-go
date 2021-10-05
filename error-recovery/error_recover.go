package main

import "fmt"

func one() {
	defer fmt.Println("one defer")
	fmt.Println("one")
	two()
}

func two() {
	defer fmt.Println("two defer")
	fmt.Println("two")
	three()
}

func calmDown() {
	fmt.Println(recover())
}

func three() {
	defer fmt.Println("three defer")
	fmt.Println("three")

	// since the four function recovers from the panic, this fn block continues its
	// execution normally
	four()
	fmt.Println("four fn call has panicked. but i'm still running. thanks to recover")
}

func four() {
	// Calling recover will not cause execution to resume at the point of the panic, at least not exactly. The
	// function that panicked will return immediately, and none of the code in that function’s block following
	// the panic will be executed. After the function that panicked returns, however, normal execution
	// resumes.
	defer calmDown()

	// panic will stop the function execution and return immediately
	panic("four fn panicked!")

	// this won't be run
	defer fmt.Println("four defer")
}

func main() {
	defer fmt.Println("main defer")
	one()
}
