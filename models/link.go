package models

import (
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Url  string
	Text string
}

// BuildLinks is a wrapper function for getLinks to return a list of links in a HTML Node
func BuildLinks(node *html.Node) (links []Link) {
	links = []Link{}
	getLinks(node, &links)
	return
}

// getLinks recursively iterates through a Node to
// populate the inputted list with found links
func getLinks(node *html.Node, links *[]Link) {
	// app link nodes to links list
	if node.Type == html.ElementNode && node.Data == "a" {
		link := linkFromNode(node)
		*links = append(*links, link)
		return
	}

	// recurse over non-link nodes
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		getLinks(child, links)
	}
}

// linkFromNode returns a Link instance from a HTML Node.
func linkFromNode(node *html.Node) (link Link) {
	// set href
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			link.Url = attr.Val
			break
		}
	}

	// set text
	link.Text = iterateLinkTexts(node, &link)
	return link
}

// iterateLinkTexts recursively iterates over Text Nodes and child
// to retrieve a concatenation of the data contents.
func iterateLinkTexts(node *html.Node, link *Link) string {
	var text string
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		// append the data when the child is a text
		if child.Type == html.TextNode {
			text += " " + strings.TrimSpace(child.Data)

			// recurse if the child is an element
		} else if child.Type == html.ElementNode {
			text += iterateLinkTexts(child, link)
		}
	}
	return text
}
