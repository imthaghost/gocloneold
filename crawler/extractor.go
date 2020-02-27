package crawler

import (
	"fmt"
	"path"

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
		url := r.Ctx.Get("url")
		base := path.Base(url)
		fmt.Println(base)

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
