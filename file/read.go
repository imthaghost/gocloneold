package file

import (
	"io/ioutil"

	erros "github.com/imthaghost/goclone/errors"
)

// read contents of file
func ReadContents(path string) []byte {
	content, err := ioutil.ReadFile(path)
	// nil check
	erros.Check(err)
	return content
}
