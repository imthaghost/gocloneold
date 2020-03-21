package html

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// LinkRestructure grabs all html files in project directory
// reorganizes each file with local links (css js images)
func LinkRestructure(projectDir string) {
	// css project directory
	cssfiles, err := ioutil.ReadDir(projectDir + "/css")
	// uh oh :(
	if err != nil {
		panic(err)
	}
	// loop through each css file and
	for _, f := range cssfiles {
		fmt.Println(f.Name())
		//

	}
	// read each line of the HTML file
	file, err := os.Open(projectDir + "/index.html")
	// uh oh :(
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// scanner
	// re := regexp.MustCompile(`href="([^"]+\.css)`)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	defer file.Close()
	// re.FindAllStringSubmatch(txtlines[0], -1)
	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}

}

// input, err := ioutil.ReadFile(projectDir + "/index.html")
// if err != nil {
// 	panic(err)
// }

// output := bytes.Replace(input, []byte("replaceme"), []byte("ok"), -1)

// if err = ioutil.WriteFile(projectDir+"/index.html", output, 0777); err != nil {
// 	panic(err)
// }
