package file

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// WriteTo will write the contents to desired file
func WriteTo(filename string, stream []byte) {
	ioutil.WriteFile(filename, []byte(stream), 0666)
	dst, err := os.Create(filepath.Join(dir, filepath.Base(file.Filename))) // dir is directory where you want to save file.
	if err != nil {
		checkErr(err)
		return
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		checkErr(err)
		return
	}

}

// CreateProject initializes the project directory
func CreateProject(projectName string) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	projectPath := path + "/" + projectName
	os.MkdirAll(projectPath, 0666)
}
