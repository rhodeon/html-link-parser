package main

import (
	"flag"
)

func parseFlags() (html *string) {
	html = flag.String("HTML", "resources/ex1.html", "Name of HTML file to parse.")
	flag.Parse()
	return
}
