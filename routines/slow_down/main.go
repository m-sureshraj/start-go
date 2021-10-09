package main

import (
	"fmt"
	"time"
)

func sleep(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(time.Second)
	}
	fmt.Println(name, "wakes up")
}

func send(channel chan string) {
	sleep("sending routine", 2)

	fmt.Println("***sending value***")
	channel <- "foo"

	fmt.Println("***sending value***")
	channel <- "bar"
}

func main() {
	myChannel := make(chan string)
	go send(myChannel)

	sleep("main routine", 5)

	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
