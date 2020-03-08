package crawler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Extractor visits a link dtermines if its a page or sublink downloads
// the contents to a correct directory in project folder
func Extractor(link string) {
	fmt.Println("Extracting --> ", link)
	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	base := path.Base(link)
	extension := filepath.Ext(base)
	if strings.Contains(extension, ".css") {
		var name = base[0 : len(base)-len(extension)]
		fmt.Println(name + ".css")
	}
	if strings.Contains(extension, ".js") {
		var name = base[0 : len(base)-len(extension)]
		fmt.Println(name + ".js")
	}
	//write as it downloads and not load the whole file into memory.
	out, err := os.Create(base + extension)
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	fmt.Println(err)

}
