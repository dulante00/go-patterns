package data

import (
	"io"
	"os"
)

type myFile struct {
	filename string
	*os.File
}

func newDiskStorage() Store {
	return &myFile{}
}

//type myFile struct {
//	mFilename string
//	fd        *os.File
//}
//
//func (f *myFile) Read(p []byte) (n int, err error) {
//	n, err = f.fd.Read(p)
//
//	if n > 0 {
//		if n, err := f.fd.Write(p[:n]); err != nil {
//			return n, err
//		}
//	}
//
//	return
//}
//
//func (f *myFile) Write(p []byte) (n int, err error) {
//	n, err = f.fd.Write(p)
//	//runtime.KeepAlive(f)
//	return n, err
//}
//
//func (f *myFile) Close() error {
//	err := f.fd.Close()
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (f *myFile) Open(s string) (io.ReadWriteCloser, error) {
	fd, err := os.OpenFile(s, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &myFile{s, fd}, nil
}
