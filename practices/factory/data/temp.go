package data

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type temp struct {
	TmpFilename string
	*os.File
}

func newTempStorage() Store {
	return &temp{}
}

func (t *temp)  Open(prefix string) (io.ReadWriteCloser, error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), prefix)
	if err != nil {
		fmt.Printf("Cannot create temporary file, %v", err)
		return nil, err
	}
	fmt.Println("Created File: " + tmpFile.Name())
	name := tmpFile.Name()
	return &temp{name, tmpFile}, nil
}
