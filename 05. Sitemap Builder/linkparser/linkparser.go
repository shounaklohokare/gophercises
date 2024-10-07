package linkparser

import (
	"fmt"
	"io"

	"strings"

	"golang.org/x/net/html"
)

type link struct {
	Text string
	Href string
}

func GetUrls(r io.Reader) ([]link, error) {

	root, err := html.Parse(r)
	if err != nil {
		fmt.Printf("Failed to parse the HTML file:- %s", err)
	}

	aChan := make(chan *html.Node)
	go findAnchors(root, aChan)

	var links []link

	for a := range aChan {
		links = append(links, link{
			Text: getText(a),
			Href: getHref(a),
		})
	}

	return links, nil

}

func findAnchors(node *html.Node, aChan chan *html.Node) {

	if node.Type == html.ElementNode && node.Data == "a" {
		aChan <- node
		return
	}

	for next := node.FirstChild; next != nil; next = next.NextSibling {
		findAnchors(next, aChan)
	}

	if node.Parent == nil {
		close(aChan)
	}

}

func getHref(node *html.Node) string {

	for _, attr := range node.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

func getText(node *html.Node) string {

	var text string

	for next := node.FirstChild; next != nil; next = next.NextSibling {
		if next.Type == html.TextNode {
			text += next.Data
		} else {
			text += getText(next)
		}
	}

	return strings.TrimSpace(text)

}
