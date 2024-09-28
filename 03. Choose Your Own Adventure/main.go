package main

import (
	"encoding/json"
	"flag"
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

var story Story

func main() {

	fileName := flag.String("filename", "gopher.json", "path to the JSON file containing the story")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&story) // if decoding is successful it returns nil else it returns an error
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/{arc}", handleArc)

	log.Fatal(http.ListenAndServe(":8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	msg := story["intro"]

	tmpl.Execute(w, msg)

}

func handleArc(w http.ResponseWriter, r *http.Request) {

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
