package html

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ArrangeCSS arrages css files in index file
func ArrangeCSS(projectDir string) {
	// css project directory
	indexfile := projectDir + "/index.html"
	input, err := ioutil.ReadFile(indexfile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	cssfiles, err := ioutil.ReadDir(projectDir + "/css")
	// uh oh :(
	if err != nil {
		panic(err)
	}
	fileindex := 0
	for fileindex < len(cssfiles) {
		for index, line := range lines {
			b := []byte(line)
			r := bytes.NewReader(b)
			doc, err := goquery.NewDocumentFromReader(r)
			if err != nil {
				panic(err)
			}

			// Find the review items
			doc.Find("link[rel='stylesheet']").Each(func(i int, s *goquery.Selection) {
				// For each item found, get the band and title
				data, err := s.Attr("href")
				fmt.Println(err)
				if data != "" {
					s.SetAttr("href", "css/"+cssfiles[fileindex].Name())

					data, err := s.Attr("href")
					lines[index] = fmt.Sprintf(`<link rel="stylesheet" type="text/css" href="%s"/>`, data)
					fmt.Println(data, err)
					fileindex++
				}

			})

		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(indexfile, []byte(output), 0777)
	if err != nil {
		log.Fatalln(err)
	}
}