package main

import (
	"fmt"
	"io/ioutil"
)

type Data struct {
	Content string
}

func readFile() {
	// read file 
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
	}
	fmt.Print(string(fileContents))
}

func writeFile() {
	// write file 
	bytesToWrite := []byte("hello\ngo\n")
	err := ioutil.WriteFile("first-post.txt", bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}	
}

func writeTemplateToFilesystem() {
	
}


func main() {
	fmt.Println("Hello, world!")

	readFile()
	// writeFile()

}