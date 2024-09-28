package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/shounaklohokare/gophercises/urlshort/urlshort"
)

func main() {

	yamlFileName := flag.String("yaml", "urls.yaml", "Path to YAML file containing urls")
	jsonFileName := flag.String("json", "urls.json", "Path to JSON file containing urls")
	// parse above file names
	flag.Parse()

	// web router
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlFile, err := os.Open(*yamlFileName)
	if err != nil {
		fmt.Printf("Failed to load file %v due to %v", *yamlFileName, err)
		return
	}

	defer yamlFile.Close()

	yamlData, err := io.ReadAll(yamlFile)
	if err != nil {
		fmt.Printf("Failed to load file %v due to %v", *yamlFileName, err)
		return
	}

	yamlHandler, err := urlshort.YAMLHandler(yamlData, mapHandler)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonFile, err := os.Open(*jsonFileName)
	if err != nil {
		fmt.Printf("Failed to load file %v due to %v", *jsonFileName, err)
		return
	}

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("Failed to load file %v due to %v", *jsonFileName, err)
		return
	}

	jsonHandler, err := urlshort.JSONHandler(jsonData, mapHandler)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.ListenAndServe(":9090", jsonHandler)
	http.ListenAndServe(":8080", mapHandler)
	http.ListenAndServe(":7832", yamlHandler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Error Page!")
}
