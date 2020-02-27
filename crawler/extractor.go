package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Extractor visits a link and downloads its contents to a file
func Extractor(link string) {
	// create a new collector
	c := colly.NewCollector(colly.Async(true))
	//Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {

		r.Ctx.Put("url", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {

		fmt.Println("Extracting ", "-->", r.Ctx.Get("url"))
		// data := r.Body
		// fmt.Println(string(data))

	})
	c.Visit(link)
	c.Wait()
}
