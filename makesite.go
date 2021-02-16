package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Content represents the text in .txt files 
// to be generated into HTML files 
type Content struct {
	Content string
}

// Page referenced at https://gist.github.com/droxey/5984bf42810ad53f03b9c465e1484449
// Page holds all the information we need to generate a new
// HTML page from a text file on the filesystem.
type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content      string
}

func readFile(file string) string {

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

func createTranslatedPageFromTextFile(lang string, filePath string) {
	// Make sure we can read in the file first!
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		fmt.Println(filepath.Ext(file.Name()))

		// Get the name of the file without `.txt` at the end.
		// Use this later when naming new HTML file.
		if filepath.Ext(file.Name()) == ".txt" {
			fileContents := readFile(file.Name())
			translated, err := translateText(lang, fileContents)

			// Instantiate a new Page.
			// Populate each field with data
			var page Page
			page.Content = translated

			if err != nil {
				log.Fatal(err)
			}

			// Create a new template in memory named "template.tmpl".
			// When the template is executed, it will parse template.tmpl,
			// looking for {{ }} where we can inject content.
			t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

			// Create a new, blank HTML file.
			newFile, err := os.Create(strings.Split(file.Name(), ".txt")[0] + ".html")

			if err != nil {
				panic(err)
			}

			// Executing the template injects the Page instance's data,
			// allowing us to render the content of our text file.
			// Furthermore, upon execution, the rendered template will be
			// saved inside the new file we created earlier.
			err = t.Execute(newFile, page)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {
	var dir string  // flag to find all .txt files in the given directory
	var lang string // flag to choose tranlation
	flag.StringVar(&dir, "dir", "", "Directory with text files & converted HTML files")
	flag.StringVar(&lang, "lang", "zh", "This is the language you want to translate, inputting google's language abbreviations.")
	flag.Parse()

	createTranslatedPageFromTextFile(lang, dir)
}
