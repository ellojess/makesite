package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	directory := "."
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name())
		}

	}
}
