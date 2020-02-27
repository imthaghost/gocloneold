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

	// queue for css links
	cssqueue := make([]string, 0)

	// on every link tag that has a rel attribute that is equal to stylesheet - CSS
	c.OnHTML("link[rel='stylesheet']", func(e *colly.HTMLElement) {

		// hyperlink reference
		link := e.Attr("href")

		// enqueue
		cssqueue = append(cssqueue, link)

		// print css file was found
		fmt.Println("Css found", "-->", link)

		// Extract contents at specified links
		Extractor(e.Request.AbsoluteURL(link))
	})
	c.Visit(site)
	c.Wait()

}

/* ********************************** Extension of same collector *******************************
//SSCollector ...
func CSSCollector(c *colly.Collector) {
	c.OnHTML("link[rel='stylesheet']", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		cssqueue = append(cssqueue, link)
		Print link
		fmt.Println("Css found", "-->", link)
		c.Visit(e.Request.AbsoluteURL(link))
		Extractor(e.Request.AbsoluteURL(link))
	})
}
********************************* Extension of same collector *********************************/
