package html

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ArrangeImgs arranges all images in project
func ArrangeImgs(projectDir string) {
	// css project directory
	indexfile := projectDir + "/index.html"
	input, err := ioutil.ReadFile(indexfile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	images, err := ioutil.ReadDir(projectDir + "/imgs")
	// uh oh :(
	if err != nil {
		panic(err)
	}
	fileindex := 0
	for fileindex < len(images) {
		for index, line := range lines {
			b := []byte(line)
			r := bytes.NewReader(b)
			doc, err := goquery.NewDocumentFromReader(r)
			if err != nil {
				panic(err)
			}
			// Find the review items
			doc.Find("img[src]").Each(func(i int, s *goquery.Selection) {
				// For each item found, get the band and title
				data, err := s.Attr("src")
				fmt.Println(err)
				if data != "" {
					s.SetAttr("src", "imgs/"+images[fileindex].Name())
					data, err := s.Attr("src")
					lines[index] = fmt.Sprintf(`<img src="%s">`, data)
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
