package link

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestBuildLinks(t *testing.T) {
	tests := []struct {
		name string
		in   io.Reader
		want Links
	}{
		{
			name: "single link",
			in:   strings.NewReader(testHtml["single link"]),
			want: Links{
				Link{
					Url:  "/other-page",
					Text: "A link to another page",
				},
			},
		},
		{
			name: "multiple links",
			in:   strings.NewReader(testHtml["multiple links"]),
			want: Links{
				Link{
					Url:  "/other-page",
					Text: "A link to another page",
				},
				Link{
					Url:  "/another-page",
					Text: "A link to yet another page",
				},
			},
		},
		{
			name: "links with nested elements",
			in:   strings.NewReader(testHtml["links with nested elements"]),
			want: Links{
				Link{
					Url:  "https://www.github.com/rhodeon",
					Text: "Check me out on GitHub!",
				},
				Link{
					Url:  "https://github.com/rhodeon/html-link-parser",
					Text: "Self-referencing link!",
				},
			},
		},
		{
			name: "link with comment",
			in:   strings.NewReader(testHtml["link with comment"]),
			want: Links{
				Link{
					Url:  "/cipher",
					Text: "zero cipher",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := BuildLinks(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nGot: \t%#v;\nWant: \t%#v", got, tt.want)
			}
		})
	}

}

func TestLinks_GetTexts(t *testing.T) {
	links := testLinks
	want := []string{"github.com", "reddit.com", ""}
	got := links.GetUrls()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nGot: \t%#v;\nWant: \t%#v", got, want)
	}
}

func TestLinks_GetUrls(t *testing.T) {
	links := testLinks
	want := []string{"share repository", "", "view profile"}
	got := links.GetTexts()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nGot: \t%#v;\nWant: \t%#v", got, want)
	}
}

// TODO: Add Examples
