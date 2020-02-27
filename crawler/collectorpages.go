package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

// PageCollector ...
func PageCollector(site string) {

	// create a new collector
	c := colly.NewCollector(
		colly.Async(true),
	)

	// on every a tag that has a href attribute - All links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		// hyperlink reference
		link := e.Attr("href")

		// print page was found
		fmt.Println("Page found", "-->", link)

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
