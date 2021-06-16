package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Link struct {
	Href string `json:"href"`
	Text string `json:"text"`
}

func getText(n *html.Node) string {

	text := ""
	for c := n; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			text = fmt.Sprintf("%s %s", text, strings.TrimSpace(c.Data))
		} else if c.Type == html.ElementNode && c.FirstChild != nil {
			text = fmt.Sprintf("%s %s", text, strings.TrimSpace(c.FirstChild.Data))
		}

	}
	return text
}

func main() {

	s := `
	<body>
	<a href="/homepage">
	<p>Before text</p>
	Home
	<h5>After text</h5>
	<h6>After text 2</h6>
	<!-- comment -->
	</a>
	</body>
	`
	

	doc, err := html.Parse(strings.NewReader(s))
	check(err)
	links := make([]Link, 0, 0)
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Parent.Data != "a" {
			if n.Data == "a" {
				text := ""
				href := ""

				if len(n.FirstChild.Data) > 0 {
					text = text + strings.TrimSpace(n.FirstChild.Data)
				}
				text += getText(n.FirstChild.NextSibling)

				for _, a := range n.Attr {
					if a.Key == "href" {
						href = a.Val
					}
				}
				l := Link{href, text}
				links = append(links, l)
			}

		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}

	}

	f(doc)
	jsonLinks, _ := json.Marshal(links)

	fmt.Println(string(jsonLinks))
}
