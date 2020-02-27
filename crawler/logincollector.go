package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

// LoginCollector ...
func LoginCollector() {
	const (
		csrfTokenSelector = "#main-container > section.content > main > div > div.auth-form.sign-in-form > form > input[type=hidden]:nth-child(2)"
		verifySelector    = "#screenshot-10392298 > div > div.dribbble-shot.with-actions > div.dribbble-img.gif > a.dribbble-over"
	)
	c := colly.NewCollector()
	err := c.Post("https://www.makeschool.com/login", map[string]string{"user[email]": "gary.frederick@smash.lpfi.org", "user[password]": "ouoyou12"})
	if err != nil {
		panic(err)
	}
	fmt.Print(c.Cookies("https://www.makeschool.com/dashboard"))
	// HTML
	c.OnHTML(csrfTokenSelector, func(e *colly.HTMLElement) {
		//authenticity token
		//token := e.Attr("value")
		//authenticate
		err := c.Post("https://www.makeschool.com/login", map[string]string{"user[email]": "gary.frederick@smash.lpfi.org", "user[password]": "ouoyou12"})
		if err != nil {
			panic(err)
		}
		fmt.Print(c.Cookies("https://www.makeschool.com/dashboard"))
	})
	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		// var body []string
		// fmt.Print(string(r.Body))
		// c.OnHTML(verifySelector, func(e *colly.HTMLElement) {
		// 	body = append(body, e.Attr("href"))
		// 	fmt.Print(body)
		// })
		fmt.Println(r.Body)
	})
	// start scraping
	c.Visit("https://www.makeschool.com/login")
}
