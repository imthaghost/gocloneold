package main

import (
	"fmt"
	"os"

	"github.com/imthaghost/goclone/crawler"
	"github.com/imthaghost/goclone/file"
	"github.com/imthaghost/goclone/request"
)

func main() {
	// if the domain is valid and the url is valid return the text
	url := os.Args[1]
	// grab the url from the
	if !request.ValidateURL(url) && !request.ValidateDomain(url) {
		fmt.Println("goclone <url>")
	} else if request.ValidateDomain(url) {
		// use the domain as the project name
		name := url
		// CreateProject
		file.CreateProject(name)
		// create the url
		validURL := request.CreateURL(name)
		// Crawler
		crawler.Crawl(validURL)
		//crawler.Extractor(validURL)
	} else if request.ValidateURL(url) {
		// get the hostname
		name := request.GetDomain(url)
		// create project
		file.CreateProject(name)
		// Crawler
		crawler.Crawl(url)
		//crawler.Extractor(url)
	} else {
		fmt.Print(url)
	}

}
