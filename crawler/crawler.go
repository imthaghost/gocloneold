package crawler

// Crawl asks the neccessary crawlers for collecting links for building the web page
func Crawl(site string) {
	// check to see if we should use login functionality
	// login := false
	// if !login {
	// 	// login collector
	// 	LoginCollector(site)
	// }
	// call the normal collector
	Collector(site)
}
