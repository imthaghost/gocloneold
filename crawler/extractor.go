package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Extractor visits a link dtermines if its a page or sublink downloads
// the contents to a correct directory in project folder
func Extractor(link string, page bool) {
	// create a new collector
	c := colly.NewCollector(colly.Async(true))
	// Before making a request
	c.OnRequest(func(r *colly.Request) {
		// put the url as context in between request and response
		r.Ctx.Put("url", r.URL.String())
	})
	// After making a request
	c.OnResponse(func(r *colly.Response) {
		// get url from the response callback as context
		url := r.Ctx.Get("url")
		// Extract contents
		fmt.Println("Extracting ", "-->", url)
		// url := r.Ctx.Get("url")
		// base := path.Base(url)
		// fmt.Println(base)

		// write as it downloads and not load the whole file into memory.
		// Create the file
		// out, err := os.Create(filepath)
		// if err != nil {
		// 	return err
		// }
		// defer out.Close()

		// // Write the body to file
		// _, err = io.Copy(out, resp.Body)
		// return err

	})
	c.Visit(link)
	c.Wait()
}
