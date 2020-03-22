package html

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ArrangeJS arranges javascript in index file
func ArrangeJS(projectDir string) {
	// css project directory
	indexfile := projectDir + "/index.html"
	input, err := ioutil.ReadFile(indexfile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	jsfiles, err := ioutil.ReadDir(projectDir + "/js")
	// uh oh :(
	if err != nil {
		panic(err)
	}
	// fileindex := 0

	for _, jsfile := range jsfiles {
		for index, line := range lines {
			b := []byte(line)
			r := bytes.NewReader(b)
			doc, err := goquery.NewDocumentFromReader(r)
			if err != nil {
				panic(err)
			}

			// Find the review items
			doc.Find("script[src]").Each(func(i int, s *goquery.Selection) {
				// For each item found, get the band and title
				data, err := s.Attr("src")
				fmt.Println(err)
				if data != "" {
					s.SetAttr("src", "js/"+jsfile.Name())
					data, err := s.Attr("src")
					lines[index] = fmt.Sprintf(`<script src="%s"></script>`, data)
					fmt.Println(data, err)

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
