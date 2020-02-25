package main

import "github.com/imthaghost/goclone/crawler"

// func checkargs(params []string) bool {
// 	// no url or too many arguments
// 	if len(params) > 2 || len(params) == 1 {
// 		return false
// 	}
// 	return true
// }
func main() {
	// flags

	// help, v := flags.ParseFlags()
	// fmt.Print(v)
	// if help != false {
	// 	log.Println("hi")
	// }
	// if v != true {
	// 	fmt.Print("Verbose disabled")
	// }
	// if checkargs(os.Args) == true {
	// 	url := os.Args[1]
	// 	//fmt.Println(url)
	// 	fmt.Println(request.Validate(url))
	// } else {
	// description := `
	// 	Easily copy websites to your computer with goclone
	// 	goclone is a utility that allows you to download a website
	// 	from the Internet to a local directory, building recursively
	// 	all directories, getting html, images, and other files from
	// 	the server to your computer.

	// 	// See Full Documentation: https://github.com/imthaghost/goclone`

	// 	incorrect := `
	// 	goclone <url>
	// 	example: goclone https://tesla.com

	// 	See Full Command List:

	// 	goclone --help

	// 	`
	// 	fmt.Println(incorrect)
	// }

	// Crawler
	crawler.Crawl()
}
