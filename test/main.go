package main

import (
	"fmt"
	"start-go/test/utils"
)

func main() {
	names := []string{"foo", "bar", "baz"}
	str := utils.Join(names)

	fmt.Println(str)
}
