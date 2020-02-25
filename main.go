package main

import (
	"fmt"
	"os"
)

func checkargs(params []string) bool {
	// no url or too many arguments
	if len(params) > 2 || len(params) == 1 {
		return false
	}
	return true
}
func main() {
	// var url string
	if checkargs(os.Args) == true {
		fmt.Println(os.Args)
	} else {
		// description := `
		// Easily copy websites to your computer with goclone
		// goclone is a utility that allows you to download a website
		// from the Internet to a local directory, building recursively
		// all directories, getting html, images, and other files from
		// the server to your computer.

		// See Full Documentation: https://github.com/imthaghost/goclone`

		incorrect := `
		goclone <url>
		example: goclone https://tesla.com

		See Full Command List: 
		
		goclone --help

		`
		fmt.Println(incorrect)
	}

	// Crawler
	//crawler.Crawl()
}
