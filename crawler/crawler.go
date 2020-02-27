package crawler

// Crawl asks the neccessary crawlers for collecting links for building the web page
func Crawl(site string) {

	CSSCollector(site) // CSS Collection
	JSCollector(site)  // JS Collection
	ImgCollector(site) // IMG Collection

	// collecting behind login
	//LoginCollector()

	// ############################# Some Go routines #######################
	// linkChan := make(chan string)
	// csslist := CSSCollector(site)
	// go func(csslist []string) {
	// 	for _, item := range csslist {
	// 		linkChan <- item //send

	// 	}
	// }(csslist)
	// go func() {
	// 	for _, item := range csslist {
	// 		log.Println(item)
	// 		found := <-linkChan //receive
	// 		fmt.Println("Found: ", found)
	// 	}
	// }()
	// ############################# Some Go routines #######################

}
