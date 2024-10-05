package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

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

}
