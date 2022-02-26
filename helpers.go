package main

import (
	"flag"
	"strings"
	"unicode"
)

func parseFlags() (html *string) {
	html = flag.String("HTML", "resources/ex2.html", "Name of HTML file to parse.")
	flag.Parse()
	return
}

// trimDuplicateSpaces reduces multiple whitespaces at the beginning and end
// of a string to one each.
// The unmodified string is returned if it has less than 2 characters.
func trimDuplicateSpaces(s string) string {
	if len(s) <= 1 {
		return s
	}

	// check beginning and end of the string for spaces
	var begin bool
	if unicode.IsSpace(rune(s[0])) {
		begin = true
	}
	var end bool
	if unicode.IsSpace(rune(s[len(s)-1])) {
		end = true
	}

	// trim and add a space each to the beginning and end of the string
	// if it previously had them
	s = strings.TrimSpace(s)
	if begin {
		s = " " + s
	}
	if end {
		s = s + " "
	}

	return s
}
