package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path
		fmt.Printf("Path: %v\n", path)
		fmt.Println(pathsToUrls)

		redirectURL, ok := pathsToUrls[path]
		if !ok {
			fmt.Println("Cannot find file file path, redirecting to default error page")

			fallback.ServeHTTP(w, r)
			return
		}

		http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)

	}

}

type pathUrl struct {
	Path string
	URL  string
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	var parsedYaml = []pathUrl{}
	err := yaml.Unmarshal(yml, &parsedYaml)
	if err != nil {
		return nil, err
	}

	pathsToUrls := map[string]string{}
	for _, ymlEntry := range parsedYaml {
		pathsToUrls[ymlEntry.Path] = ymlEntry.URL
	}

	return MapHandler(pathsToUrls, fallback), nil

}

func JSONHandler(jsonData []byte, fallback http.Handler) (http.HandlerFunc, error) {

	var parsedJson = []pathUrl{}
	err := json.Unmarshal(jsonData, &parsedJson)
	if err != nil {
		return nil, err
	}

	pathsToUrls := map[string]string{}
	for _, jsonEntry := range parsedJson {
		pathsToUrls[jsonEntry.Path] = jsonEntry.URL
	}

	return MapHandler(pathsToUrls, fallback), nil
}
