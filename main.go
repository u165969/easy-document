package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
)

// DataInDoc is the data that will be applied to HTML template
type DataInDoc struct {
	Title    string
	CSS      string
	JS       string
	Document string
}

func document(input, output, tmpl, title, css, js string) error {
	// read a markdown file
	markdown, err := ioutil.ReadFile(input)
	if err != nil && err != errors.New("EOF") {
		return err
	}

	// convert the markdown into HTML and sanitize it
	unsafeHTML := blackfriday.Run(markdown)
	safeHTML := bluemonday.UGCPolicy().SanitizeBytes(unsafeHTML)

	// generate HTML output
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()
	data := DataInDoc{Title: title, CSS: css, JS: js, Document: string(safeHTML)}
	if err = t.Execute(file, data); err != nil {
		return err
	}

	return nil
}

func main() {
	// set variables
	var input, output, tmpl, title, css, js string
	flag.StringVar(&input, "input", "input/document.md", "input filename")
	flag.StringVar(&output, "output", "output/document.html", "output filename")
	flag.StringVar(&tmpl, "tmpl", "tmpl/document.tmpl", "template filename")
	flag.StringVar(&title, "title", "title", "title in HTML")
	flag.StringVar(&css, "css", "css/document.css", "css file in HTML")
	flag.StringVar(&js, "js", "js/document.js", "javascript file in HTML")
	flag.Parse()

	// create a document
	if err := document(input, output, tmpl, title, css, js); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create a document: %v\n", err)
		os.Exit(1)
	}
}
