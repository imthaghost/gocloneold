package crawler

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"

	"github.com/imthaghost/goclone/request"
)

// Crawl ...
func Crawl(site string) {
	//cssqueue := make([]string, 0)
	// create a new collector
	c := colly.NewCollector(colly.Async(true))
	// on every link tag that has a rel attribute
	c.OnHTML("link[rel='stylesheet']", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Println("Css found", "-->", link)

	})
	// on every script tag
	c.OnHTML("script[src]", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		// Print link
		fmt.Println("Js found", "-->", link)

	})

	// on every img tag
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		// Print link
		fmt.Println("Img found", "-->", link)
	})
	// // on every img tag
	// c.OnHTML("video", func(e *colly.HTMLElement) {
	// 	link := e.Text
	// 	// Print link
	// 	fmt.Println("Video found", "-->", link)
	// 	// Visit link found on page
	// 	// Only those links are visited which are in AllowedDomains
	// 	//c.Visit(e.Request.AbsoluteURL(link))
	// })

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", "-->", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		host := request.GetDomain(site)
		projectPath := path + "/" + host + "/"
		f, err := os.Create(projectPath + "index.html")
		if err != nil {
			fmt.Println(err)

		}
		f.WriteString(string(r.Body))
		f.Close()

		//file.WriteTo(projectPath+"index.html", data)

		// c.OnHTML("link[href]", func(e *colly.HTMLElement) {
		// 	// grab the hyper link reference
		// 	link := e.Attr("href")
		// 	// push css link to the queue
		// 	cssqueue = append(cssqueue, link)
		// })
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
