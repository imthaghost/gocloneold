package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Crawl ...
func Crawl(site string) {

	// create a new collector
	c := colly.NewCollector(colly.Async(true))
	// on every link tag that has a rel attribute that is equal to stylesheet - CSS
	c.OnHTML("link[rel='stylesheet']", func(e *colly.HTMLElement) {
		cssqueue := make([]string, 0)
		link := e.Attr("href")
		cssqueue = append(cssqueue, link)
		// Print link
		fmt.Println("Css found", "-->", link)
		c.Visit(e.Request.AbsoluteURL(link))
	})
	// on every script tag that has a src attribute - JS
	c.OnHTML("script[src]", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		// Print link
		fmt.Println("Js found", "-->", link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// on every img tag that has a src attribute - Images
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		// Print link
		fmt.Println("Img found", "-->", link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// todo figure how the fuck to pull videos from pages that create their own data attributes
	// // on every img tag that has a src attribute - Images
	// c.OnHTML("video", func(e *colly.HTMLElement) {
	// 	link := e.Attr("src")
	// 	// Print link
	// 	fmt.Println("Video found", "-->", link)
	// 	c.Visit(e.Request.AbsoluteURL(link))
	// })

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", "-->", r.URL.String())
		r.Ctx.Put("url", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {

		fmt.Println("Visited: ", r.Ctx.Get("url"))

	})
	c.Visit(site)
	c.Wait()

	// const (
	// 	csrfTokenSelector = "#main-container > section.content > main > div > div.auth-form.sign-in-form > form > input[type=hidden]:nth-child(2)"
	// 	verifySelector    = "#screenshot-10392298 > div > div.dribbble-shot.with-actions > div.dribbble-img.gif > a.dribbble-over"
	// )
	// // create a new collector
	// //c := colly.NewCollector()
	// err := c.Post("https://www.makeschool.com/login", map[string]string{"user[email]": "gary.frederick@smash.lpfi.org", "user[password]": "ouoyou12"})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Print(c.Cookies("https://www.makeschool.com/dashboard"))
	// // HTML
	// // c.OnHTML(csrfTokenSelector, func(e *colly.HTMLElement) {
	// // 	// authenticity token
	// // 	//token := e.Attr("value")
	// // 	// authenticate
	// // 	// err := c.Post("https://www.makeschool.com/login", map[string]string{"user[email]": "gary.frederick@smash.lpfi.org", "user[password]": "ouoyou12"})
	// // 	// if err != nil {
	// // 	// 	panic(err)
	// // 	// }
	// // 	// fmt.Print(c.Cookies("https://www.makeschool.com/dashboard"))
	// // })
	// // attach callbacks after login
	// c.OnResponse(func(r *colly.Response) {
	// 	// var body []string
	// 	// fmt.Print(string(r.Body))
	// 	// c.OnHTML(verifySelector, func(e *colly.HTMLElement) {
	// 	// 	body = append(body, e.Attr("href"))
	// 	// 	fmt.Print(body)
	// 	// })
	// 	log.Print(string(r.Body))
	// })
	// // start scraping
	// c.Visit("https://www.makeschool.com/login")
}
