package models

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

const (
	hyperlink = "a"
	href      = "href"
)

// Link represents HTML hyperlink elements.
// Url corresponds to href.
// Text corresponds to the content.
type Link struct {
	Url  string
	Text string
}

type Links []Link

// BuildLinks returns a list of links from the given Reader.
func BuildLinks(r io.Reader) (Links, error) {
	var links Links
	node, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	getLinks(node, &links)
	return links, nil
}

// GetUrls maps links to a list of their Urls.
func (links Links) GetUrls() (urls []string) {
	for _, link := range links {
		urls = append(urls, link.Url)
	}
	return
}

// GetTexts maps links to a list of their Texts.
func (links Links) GetTexts() (texts []string) {
	for _, link := range links {
		texts = append(texts, link.Text)
	}
	return
}

// getLinks recursively iterates through a Node to
// populate the inputted list with found links.
func getLinks(node *html.Node, links *Links) {
	// app link nodes to links list
	if node.Type == html.ElementNode && node.Data == hyperlink {
		link := linkFromNode(node)
		*links = append(*links, link)
		return
	}

	// recurse over non-link nodes
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		getLinks(child, links)
	}
}

// linkFromNode returns a Link instance from an HTML Node.
func linkFromNode(node *html.Node) (link Link) {
	// set Url
	for _, attr := range node.Attr {
		if attr.Key == href {
			link.Url = attr.Val
			break
		}
	}

	// set Text
	link.Text = iterateTextNodes(node, &link)
	return link
}

// iterateTextNodes recursively iterates over Text Nodes and child
// to retrieve a concatenation of the data contents.
func iterateTextNodes(node *html.Node, link *Link) string {
	var text string
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		// append the data when the child is a text
		if child.Type == html.TextNode {
			text += " " + strings.TrimSpace(child.Data)

			// recurse if the child is an element
		} else if child.Type == html.ElementNode {
			text += iterateTextNodes(child, link)
		}
	}
	return text
}
