package main

import (
	"fmt"
	"strings"
)

func SayHello() {
	fmt.Println("Hello...")
}

func main() {
	broken := "suresh"
	fixed := strings.NewReplacer("s", "p").Replace(broken)
	fmt.Println(fixed)
}
