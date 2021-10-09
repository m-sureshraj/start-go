package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responseSize(url string, channel chan Response) {
	// following line will block the function execution till it
	// fetches the response
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	channel <- Response{url: url, size: len(body)}
}

type Response struct {
	url  string
	size int
}

// The main function of every Go program is started using a
// goroutine, so every Go program runs at least one goroutine.
func main() {
	channel := make(chan Response)

	urls := []string{
		"https://example.com",
		"https://sureshraj.me",
		"https://google.lk",
	}

	for _, url := range urls {
		go responseSize(url, channel)
	}

	for i := 0; i < len(urls); i++ {
		page := <-channel
		fmt.Printf("%s: %d\n", page.url, page.size)
	}
}
