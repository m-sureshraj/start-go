package main

import "fmt"

// How to initialize constant list?

func main() {
	names := [3]int{1, 2, 3}

	// typical loop
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	heights := [3]float32{6.3, 5.7, 5.54}

	// for range approach
	fmt.Println("Heights")

	for _, value := range heights {
		fmt.Printf("index = %.2f \n", value)
	}
}
