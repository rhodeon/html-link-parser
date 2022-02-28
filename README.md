# html-link-parser

A Go library for parsing HTML hyperlink elements into native Go objects.

Useful for creating sitemap generators, web crawlers and web scrapers.

## Install

```sh
go get -u github.com/rhodeon/html-link-parser
```

## Usage

A Link struct represents an HTML hyperlink element:

```go
type Link struct {
Url  string // href
Text string // content
}
```

To parse an HTML reader to a list of Link objects:

```go
html := `<html>
<body>
  <h1>Hello!</h1>
  <a href="github.com">share repository</a>
  <a href="/another-page">A link to yet another page</a>
</body>
</html>
`
links, _ := BuildLinks(strings.NewReader(html))
fmt.Printf("%+v", links)
// Output: [{Url:github.com Text:share repository} {Url:/another-page Text:A link to yet another page}]
```

To map links to a list of URLs:

```go
urls := links.GetUrls()
fmt.Printf("%#v", urls)
// Output: []string{"github.com", "/another-page"}
```

To map links to a list of Texts:

```go
texts := links.GetTexts()
fmt.Printf("%#v", texts)
// Output: []string{"share repository", "A link to yet another page"}
```