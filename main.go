package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
	"html-link-parser/models"
)

func main() {
	htmlFilename := parseFlags()
	htmlFile, err := os.Open(*htmlFilename)
	if err != nil {
		log.Fatalf("Unable to open HTML file: %v", err.Error())
	}
	defer func() {
		err := htmlFile.Close()
		if err != nil {
			log.Fatalf("Unable to close HTML file: %v", err.Error())
		}
	}()

	node, err := html.Parse(htmlFile)
	if err != nil {
		log.Fatalf("Unable to parse HTML: %v", err.Error())
	}

	links := models.BuildLinks(node)
	fmt.Printf("%v\n", links)
}
