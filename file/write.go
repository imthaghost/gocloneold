package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// WriteTo will write the contents to desired file
func WriteTo(filename string, stream []byte) {
	ioutil.WriteFile(filename, []byte(stream), 0666)
}

// CreateProject initializes the project directory
func CreateProject(projectName string) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	fmt.Println(projectName)
}
