package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
	// "github.com/mind1949/googletrans"
	// "golang.org/x/text/language"
)

type Content struct {
	Content string
}

// Ref: https://gist.github.com/droxey/5984bf42810ad53f03b9c465e1484449
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

//  Render contents of first-post.txt using Go Templates and print it to stdout
// func renderTemplate(filename string, data string) {
// 	t := template.Must(template.New("template.tmpl").ParseFiles(filename))
// 	content := Content{Content: data}
// 	err := t.Execute(os.Stdout, content)

// 	if err != nil {
// 		panic(err)
// 	}
// }

//  Write the HTML template to the filesystem to a file. Name it first-post.html.
// func writeTemplateToFile(filename string, data string) {
// 	t := template.Must(template.New("template.tmpl").ParseFiles(filename))
// 	content := Content{Content: data}
// 	f, err := os.Create("first-post.html")
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = t.Execute(f, content)
// 	if err != nil {
// 		panic(err)
// 	}

// }

func createPageFromTextFile(lang string, filePath string) Page {
	// Make sure we can read in the file first!
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// Get the name of the file without `.txt` at the end.
	// We'll use this later when naming our new HTML file.
	fileNameWithoutExtension := strings.Split(filePath, ".txt")[0]

	// Instantiate a new Page.
	// Populate each field and return the data.
	return Page{
		TextFilePath: filePath,
		TextFileName: fileNameWithoutExtension,
		HTMLPagePath: fileNameWithoutExtension + ".html",
		Content:      string(fileContents),
	}
}

func renderTemplateFromPage(lang string, templateFilePath string, page Page) {
	// Create a new template in memory named "template.tmpl".
	// When the template is executed, it will parse template.tmpl,
	// looking for {{ }} where we can inject content.
	t := template.Must(template.New(templateFilePath).ParseFiles(templateFilePath))

	// Create a new, blank HTML file.
	newFile, err := os.Create(page.HTMLPagePath)
	if err != nil {
		panic(err)
	}

	// Executing the template injects the Page instance's data,
	// allowing us to render the content of our text file.
	// Furthermore, upon execution, the rendered template will be
	// saved inside the new file we created earlier.
	t.Execute(newFile, page)
	fmt.Println("✅ Generated File: ", page.HTMLPagePath)
}

func writeTranslate(filename string, lang string) {

	FileText := readFile(filename)

	contents, error := translateText(lang, FileText)
	if error != nil {
		panic(error)
	}
	bytesToWrite := []byte(contents)

	err := ioutil.WriteFile(filename, bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}

func createTranslatedPageFromTextFile(filePath string, lang string) Page {

	// Make sure we can read in the file first!
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// Get the name of the file without `.txt` at the end.
	// We'll use this later when naming our new HTML file.
	fileNameWithoutExtension := strings.Split(filePath, ".txt")[0]

	// Instantiate a new Page.
	// Populate each field and return the data.
	return Page{
		TextFilePath: filePath,
		TextFileName: fileNameWithoutExtension,
		HTMLPagePath: fileNameWithoutExtension + ".html",
		Content:      string(fileContents),
	}
}

func main() {

	// This flag represents the name of any `.txt` file in the same directory as your program.
	// Run `./makesite --file=latest-post.txt` to test.
	var textFilePath string
	flag.StringVar(&textFilePath, "file", "", "Text file to turn into HTML page")
	// flag.Parse()

	//  flag to find all .txt files in the given directory
	var dir string
	flag.StringVar(&dir, "dir", "", "Directory with text files & converted HTML files")
	// flag.Parse()
	fmt.Println("Directory: ", dir)

	// flag to choose tranlation
	var lang string
	flag.StringVar(&lang, "lang", "es", "This is the language you want to translate, inputting google's language abbreviations.")
	flag.Parse()
	fmt.Println("Language:", lang)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		writeTranslate(file.Name(), lang)
		createTranslatedPageFromTextFile(file.Name(), lang)

	}

	// Make sure the `file` flag isn't blank.
	if textFilePath == "" {
		panic("Missing the --file flag! Please supply one.")
	}

	// Read the provided text file and store it's information in a struct.
	newPage := createPageFromTextFile(lang, textFilePath)

	// Use the struct to generate a new HTML page based on the provided template.
	renderTemplateFromPage(lang, "template.tmpl", newPage)

}
