package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

// JSCollector ...
func JSCollector(site string) {

	// create a new collector
	c := colly.NewCollector(
		colly.Async(true),
	)

	c.OnHTML("script[src]", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		// Print link
		fmt.Println("Js found", "-->", link)
		Extractor(e.Request.AbsoluteURL(link), true)
	})

	// // Before making a request print "Visiting ..."
	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", "-->", r.URL.String())
	// 	r.Ctx.Put("url", r.URL.String())
	// })

	// c.OnResponse(func(r *colly.Response) {

	// 	fmt.Println("Visited: ", r.Ctx.Get("url"))

	// })
	c.Visit(site)
	c.Wait()

}

// JSCollector ...
// func JSCollector(c *colly.Collector) {
// 	// on every script tag that has a src attribute - JS
// 	c.OnHTML("script[src]", func(e *colly.HTMLElement) {
// 		link := e.Attr("src")
// 		// Print link
// 		fmt.Println("Js found", "-->", link)
// 		Extractor(e.Request.AbsoluteURL(link))

// 	})

// }
