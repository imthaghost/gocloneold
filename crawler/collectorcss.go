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
	// cssqueue := make([]string, 0)

	// on every link tag that has a rel attribute that is equal to stylesheet - CSS
	c.OnHTML("link[rel='stylesheet']", func(e *colly.HTMLElement) {

		// hyperlink reference
		link := e.Attr("href")

		// enqueue
		// cssqueue = append(cssqueue, link)

		// print css file was found
		fmt.Println("Css found", "-->", link)

		// Extract contents at specified links
		Extractor(e.Request.AbsoluteURL(link), true)
	})

	// Before making a request
	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})
	// On
	c.OnResponse(func(r *colly.Response) {

		// fmt.Println("Visited: ", r.Ctx.Get("url"))

	})
	c.Visit(site)
	c.Wait()

}

//CSSCollector ...
// func CSSCollector(c *colly.Collector, linkChan chan<- string) {
// 	c.OnHTML("link[rel='stylesheet']", func(e *colly.HTMLElement) {
// 		link := e.Attr("href")
// 		//Print link
// 		// sends link
// 		select {
// 		case linkChan <- link:
// 			fmt.Println(link)
// 		default:
// 			fmt.Println("no link")
// 		}
// 		fmt.Println("Css found", "-->", link)
// 	})
// }
