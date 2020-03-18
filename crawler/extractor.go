package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Extractor visits a link dtermines if its a page or sublink downloads
// the contents to a correct directory in project folder
func Extractor(link string, projectPath string) {
	fmt.Println("Extracting --> ", link)
	// get the html body
	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	// Closure
	defer resp.Body.Close()
	// file base
	base := path.Base(link)
	// file extension
	extension := filepath.Ext(base)
	//
	if strings.Contains(extension, ".css") {
		var name = base[0 : len(base)-len(extension)]
		document := name + ".css"
		// get the project name and path we use the path to
		// create and write files to our project directory
		// write as it downloads and not load the whole file into memory.
		f, err := os.OpenFile(projectPath+"/"+"css/"+document, os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		htmlData, err := ioutil.ReadAll(resp.Body) //<--- here!

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		f.Write(htmlData)
		//

		// out, err := f.Create(document)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// defer out.Close()
		// Write the body to file
		// _, err = io.Copy(out, resp.Body)
		// fmt.Println(err)

	}

	//
	if strings.Contains(extension, ".js") {
		var name = base[0 : len(base)-len(extension)]
		document := name + ".js"

		f, err := os.OpenFile(projectPath+"/"+"js/"+document, os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		htmlData, err := ioutil.ReadAll(resp.Body) //<--- here!

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		f.Write(htmlData)
	}
	//
	if strings.Contains(extension, ".jpg") {
		var name = base[0 : len(base)-len(extension)]
		fmt.Println("Edited image: " + name + ".jpg")
	}
}
