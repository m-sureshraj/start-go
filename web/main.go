package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		message := []byte("Page Not found")

		switch path {
		case "/":
			message = []byte("Home Page!")
		case "/about":
			message = []byte("About Page")
		case "/contact":
			message = []byte("Contact Page")
		}

		_, err := rw.Write(message)
		if err != nil {
			log.Fatal(err)
		}
	})

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

// Notes:
// Go supports firstclass functions.
// Like in JS we can assign function to a variable, pass into other functions.

// var foo func()
// var bar func(a int, b int) int

// foo = hello
// bar = func(a, b int) int {
// 	return a + b
// }
