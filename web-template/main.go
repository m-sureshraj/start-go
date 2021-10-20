package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetMovies() []string {
	filename := "movies.txt"
	file, err := os.Open(filename)

	if os.IsNotExist(err) {
		return nil
	}

	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())

	return lines
}

type MoviesData struct {
	Movies []string
}

type NewMoviePageData struct {
	Errors []string
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		movies := GetMovies()

		switch path {
		case "/":
			html, err := template.ParseFiles("main.html")
			check(err)
			err = html.Execute(rw, MoviesData{Movies: movies})
			check(err)
			break

		case "/new-movie":
			method := request.Method

			if method == "GET" {
				html, err := template.ParseFiles("new-movie.html")
				check(err)
				err = html.Execute(rw, nil)
				check(err)
			}

			if method == "POST" {
				movieName := request.FormValue("movie")
				movieName = strings.TrimSpace(movieName)
				data := NewMoviePageData{}

				if len(movieName) == 0 {
					data.Errors = append(data.Errors, "Movie name cannot be empty")
					html, err := template.ParseFiles("new-movie.html")
					check(err)
					err = html.Execute(rw, data)
					check(err)
				} else {
					options := os.O_APPEND | os.O_CREATE | os.O_WRONLY
					file, err := os.OpenFile("movies.txt", options, os.FileMode(0600))
					check(err)

					_, err = fmt.Fprintln(file, movieName)
					check(err)

					err = file.Close()
					check(err)
					http.Redirect(rw, request, "/", http.StatusFound)
				}
			}
			break

		default:
			html, err := template.ParseFiles("not-found.html")
			check(err)
			err = html.Execute(rw, nil)
			check(err)
		}
	})

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
