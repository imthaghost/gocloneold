package crawler

// Crawl asks the neccessary crawlers for collecting links for building the web page
func Crawl(site string) {
	login := false
	if !login {
		return
	}
	// call the collector
	Collector(site)

}
