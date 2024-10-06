package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type link struct {
	Text string
	Href string
}

func main() {

	htmlFileName := flag.String("html", "ex1.html", "Path of the HTML file to be parsed")
	flag.Parse()

	f, err := os.Open(*htmlFileName)
	if err != nil {
		fmt.Printf("Failed to open the %v file :- %v", *htmlFileName, err)
	}

	defer f.Close()

	root, err := html.Parse(f)
	if err != nil {
		fmt.Printf("Failed to parse the HTML file:- %s", err)
	}

	aChan := make(chan *html.Node)
	go findAnchors(root, aChan)

	for a := range aChan {
		fmt.Printf("%+v\n", link{
			Text: getText(a),
			Href: getHref(a),
		})
	}

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
