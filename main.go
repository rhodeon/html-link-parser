package main

import (
	"fmt"
	"log"
	"os"
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

	links, err := BuildLinks(htmlFile)
	if err != nil {
		log.Fatalf("Unable to build HTML: %v", err.Error())
	}

	fmt.Printf("%v\n", links)
}
