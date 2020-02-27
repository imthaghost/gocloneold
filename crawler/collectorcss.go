package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

// CSSCollector ...
func CSSCollector(site string) {

	// create a new collector
	c := colly.NewCollector(
		colly.Async(true),
	)

	//cssqueue := make([]string, 0)
	// on every link tag that has a rel attribute that is equal to stylesheet - CSS
	c.OnHTML("link[rel='stylesheet']", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		//cssqueue = append(cssqueue, link)
		// Print link
		fmt.Println("Css found", "-->", link)
		//c.Visit(e.Request.AbsoluteURL(link))
		Extractor(e.Request.AbsoluteURL(link))
	})
	c.Visit(site)
	c.Wait()

}

// CSSCollector ...
// func CSSCollector(c *colly.Collector) {
// 	c.OnHTML("link[rel='stylesheet']", func(e *colly.HTMLElement) {
// 		link := e.Attr("href")
// 		//cssqueue = append(cssqueue, link)
// 		// Print link
// 		fmt.Println("Css found", "-->", link)
// 		//c.Visit(e.Request.AbsoluteURL(link))
// 		Extractor(e.Request.AbsoluteURL(link))
// 	})
// }
