package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

// ImgCollector ...
func ImgCollector(site string) {

	// create a new collector
	c := colly.NewCollector(
		colly.Async(true),
	)

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		// Print link
		fmt.Println("Img found", "-->", link)
		Extractor(e.Request.AbsoluteURL(link))
	})
	c.Visit(site)
	c.Wait()

}

/* ********************************** Extension of same collector *******************************
//ImgCollector ...
func ImgCollector(c *colly.Collector) {
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		// Print link
		fmt.Println("Img found", "-->", link)
		Extractor(e.Request.AbsoluteURL(link))
	})
}
********************************* Extension of same collector *********************************/
