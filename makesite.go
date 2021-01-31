package main

import (
	"fmt"
	"io/ioutil"
	"html/template"
	"os"
)

type Content struct {
	Content string
}

func readFile(file string) string{

	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(fileContents)
}

func writeFile(file string, data string) {
	bytesToWrite := []byte(data)
	err := ioutil.WriteFile(file, bytesToWrite, 0644)

	if err != nil {
		panic(err)
	}	
}

//  Render contents of first-post.txt using Go Templates and print it to stdout
func renderTemplate(filename string, data string) {
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))
	content := Content{Content: data}
	err := t.Execute(os.Stdout, content)

	if err != nil {
	  panic(err)
	}
}

//  Write the HTML template to the filesystem to a file. Name it first-post.html.
func writeTemplateToFile(filename string, data string) {
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))
	content := Content{Content: data}
	f, err := os.Create("first-post.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, content)
	if err != nil {
		panic(err)
	  }

}


func main() {
	fmt.Println("Hello, world!")

	// readFile()
	renderTemplate("template.tmpl", readFile("first-post.txt"))
	writeTemplateToFile("template.tmpl", readFile("first-post.txt"))

}