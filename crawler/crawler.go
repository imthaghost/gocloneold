package crawler

// Crawl collects neccessary links for building the web page and passes the links to the extractor for download
func Crawl(site string) {
	// CSSCollector(site)
	// JSCollector(site)
	LoginCollector()
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
	// 		fmt.Println(found)
	// 	}
	// }()
	// ############################# Some Go routines #######################

}
