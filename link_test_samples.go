package main

var testHtml = map[string]string{
	"single link": `<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`,
	"multiple links": `<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
  <a href="/another-page">A link to yet another page</a>
</body>
</html>
`,
	"links with nested elements": `<html>
<head>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
    <h1>GitHub</h1>
    <div>
        <a href="https://www.github.com/rhodeon">
            Check me out on <em>GitHub</em>! 
            <i aria-hidden="true"></i>
        </a>
        <a href="https://github.com/rhodeon/html-link-parser">
            Self-<em>referencing</em> <strong>link</strong>!
        </a>
    </div>
</body>
</html>
`,
	"link with comment": `<html>

<body>
    <a href="/cipher">zero cipher
        <!-- commented text SHOULD NOT be included! -->
    </a>
</body>

</html>
`,
}

var testLinks = Links{
	Link{
		Url:  "github.com",
		Text: "share repository",
	},
	Link{
		Url: "reddit.com",
	},
	Link{
		Text: "view profile",
	},
}
