package utils

import (
	"bytes"
	"io/ioutil"
	"os"
)

// FileContent represents byte type
type FileContent []byte

// GetFileContent gets the content of the `path` file
func GetFileContent(path string) FileContent {
	file, err := os.Open(path)
	if err != nil {
		return []byte("")
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte("")
	}

	defer file.Close()

	return data
}

// AsString converts FileContent ([]byte) to string
func (fc FileContent) AsString() string {
	return bytes.NewBuffer(fc).String()
}
