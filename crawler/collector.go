package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Collector ...
func Collector(url string) {
	// create a new collector
	c := colly.NewCollector(
		colly.Async(true),
	)
	// on every link tag that has a rel attribute that is equal to stylesheet - CSS
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		// hyperlink reference
		link := e.Attr("href")

		// print css file was found
		fmt.Println("Link found", "-->", e.Request.AbsoluteURL(link))

		// go c.Visit(e.Request.AbsoluteURL(link))

	})
	// on every link tag that has a rel attribute that is equal to stylesheet - CSS
	c.OnHTML("link[rel='stylesheet']", func(e *colly.HTMLElement) {

		// hyperlink reference
		link := e.Attr("href")

		// print css file was found
		fmt.Println("Css found", "-->", link)

		go c.Visit(e.Request.AbsoluteURL(link))

	})
	c.OnHTML("script[src]", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		// Print link
		fmt.Println("Js found", "-->", link)
		// Extractor(e.Request.AbsoluteURL(link))
		go c.Visit(e.Request.AbsoluteURL(link))
	})
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		// Print link
		fmt.Println("Img found", "-->", link)
		// Extractor(e.Request.AbsoluteURL(link))
		go c.Visit(e.Request.AbsoluteURL(link))
	})
	// Before making a request
	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})
	// On
	c.OnResponse(func(r *colly.Response) {
		link := r.Ctx.Get("url")
		// check if the url being visited is the root for searching if so write it as a page
		if url == link {
			fmt.Print(r.Request.URL)
			//Extractor(link, true)
			return
		}
		// else call the extractor as false
		Extractor(r.Ctx.Get("url"), false)
	})
	c.Visit(url)
	c.Wait()
}
