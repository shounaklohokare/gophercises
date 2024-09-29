package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"text/template"
)

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Story map[string]Chapter

func main() {

	filename := flag.String("filename", "gopher.json", "Path to the JSON file containing the story")
	webserver := flag.Bool("use_webserver", true, "Pass true to enable the game on the web server, or false to use the CLI.")
	flag.Parse()

	var story Story

	parseFile(filename, &story)

	if *webserver {
		// use web server
		useWebserver(story)
	} else {
		// use CLI
		useCLI(story)
	}

}

func useCLI(story Story) {

	nextArc := "intro"

	var op int
	for {

		arc := story[nextArc]

		fmt.Println(arc.Title + "\n")
		fmt.Println(arc.Paragraphs[0] + "\n")

		if len(arc.Options) == 0 {
			break
		}

		fmt.Printf("Please choose one of the following options to proceed:- ")
		for i, option := range arc.Options {
			fmt.Printf("\n%v. %v", (i + 1), option.Text)
		}

		fmt.Scanln(&op)

		nextArc = arc.Options[op-1].Chapter

	}
}

func useWebserver(story Story) {
	http.HandleFunc("/", story.handler)
	http.HandleFunc("/{arc}", story.handleArc)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func (story Story) handler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/index.html"))

	msg := story["intro"]

	tmpl.Execute(w, msg)

}

func (story Story) handleArc(w http.ResponseWriter, r *http.Request) {

	arc := path.Base(r.URL.Path)
	println("Arc -> ", arc)
	msg, ok := story[arc]
	if !ok {
		tmpl := template.Must(template.ParseFiles("static/error.html"))
		tmpl.Execute(w, nil)
		return
	}

	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, msg)

}

func parseFile(filename *string, story *Story) error {

	f, err := os.Open(*filename)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(f).Decode(&story); err != nil {
		return err
	}

	return nil

}
