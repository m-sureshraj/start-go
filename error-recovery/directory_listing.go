package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}

	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	} else {
		// if the panic value isn't an error, resume the panicking with the same value
		panic(err)
	}
}

func WalkDir(path string) error {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		// panic is like throw new Error() in js
		panic(err)
	}

	for _, file := range files {
		fileName := file.Name()
		fullPath := filepath.Join(path, fileName)

		if file.IsDir() {
			WalkDir(fullPath)
		} else {
			fmt.Println(fullPath)
		}
	}

	return nil
}

func main() {
	defer reportPanic()
	WalkDir("sample_dir")
}
