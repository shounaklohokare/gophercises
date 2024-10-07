package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/shounaklohokare/gophercises/sitemap/linkparser"
)

func main() {

	url := flag.String("url", "", "The URL to create a sitemap for")
	depth := flag.Int("depth", 2, "The depth of the links tree")
	xmlFileName := flag.String("xml", "sitemap.xml", "The sitemap file location to be saved")
	flag.Parse()

	if *url == "" {
		fmt.Println("-url is required")
		return
	}

	sitemapsUrls, err := buildSiteMap(*url, *depth)
	if err != nil {
		fmt.Printf("Failed to create sitemap : %v", err)
		return
	}

	fmt.Println(sitemapsUrls)

	if err := createFile(sitemapsUrls, *xmlFileName); err != nil {
		fmt.Printf("Failed to generate a sitemap in %s : %v", *xmlFileName, err)
	}

	fmt.Printf("Generated sitemap successfully with %d link(s) for %s in %s", len(sitemapsUrls), *url, *xmlFileName)

}

func buildSiteMap(baseUrl string, depth int) ([]string, error) {

	urlsMap := map[string]bool{}

	urls := []string{baseUrl}
	for d := 0; d < depth; d++ {

		var newUrls []string
		for _, url := range urls {

			pageUrls, err := getUrls(url)
			if err != nil {
				return nil, err
			}

			var uniqueSubUrls []string
			for _, subUrl := range pageUrls {

				if !urlsMap[subUrl] {
					uniqueSubUrls = append(uniqueSubUrls, subUrl)
					urlsMap[subUrl] = true
				}

			}

			newUrls = append(newUrls, uniqueSubUrls...)
		}

		urls = newUrls

	}

	return urls, nil

}

func getUrls(pageUrl string) ([]string, error) {

	pageUrl = strings.TrimSuffix(pageUrl, "/")

	res, err := http.Get(pageUrl)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	links, err := linkparser.GetUrls(res.Body)
	if err != nil {
		return nil, err
	}

	var urls []string

	for _, link := range links {
		urls = append(urls, link.Href)
	}

	var domainUrls []string

	for _, url := range urls {

		if (strings.HasPrefix(url, "http") && !strings.HasSuffix(url, pageUrl)) || (strings.Contains(url, "@")) {
			continue
		}

		if strings.HasSuffix(url, pageUrl) {
			domainUrls = append(domainUrls, url)
			continue
		}

		url = pageUrl + formatUrl(url)

		domainUrls = append(domainUrls, url)

	}

	return domainUrls, nil

}

func formatUrl(url string) string {

	if i := strings.Index(url, "#"); i != -1 {
		url = url[:i]
	}

	if url == "" || url[0] != '/' {
		url = "/" + url
	}

	return url

}

type SitemapXML struct {
	XMLName xml.Name        `xml:"urlset"`
	Xmlns   string          `xml:"xmlns,attr"`
	URLs    []SitemapXMLURL `xml:"url"`
}

type SitemapXMLURL struct {
	Loc string `xml:"loc"`
}

func createFile(urls []string, pathToXML string) error {

	var sitemap SitemapXML

	sitemap.Xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

	for _, url := range urls {
		sitemap.URLs = append(sitemap.URLs, SitemapXMLURL{
			Loc: url,
		})
	}

	f, err := os.Create(pathToXML)
	if err != nil {
		fmt.Printf("Error while creating file %s:- %v", pathToXML, err)
		return err
	}

	defer f.Close()

	return xml.NewEncoder(f).Encode(sitemap)

}
