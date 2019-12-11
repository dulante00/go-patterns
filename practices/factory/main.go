package main

import (
	"fmt"
	"github.com/sharelinux/go-patterns/practices/factory/data"
	"io/ioutil"
	"os"
	"reflect"
)

func main() {
	//fileWrite() // 测试写入磁盘文件
	//fileRead()  // 测试从磁盘文件读取文件

	fileWriteTmpFile() // 测试写入临时文件
}

func fileRead() {
	s := data.NewStore(data.DiskStorage)
	f, err := s.Open("./practices/factory/myfile.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	content, err := ioutil.ReadAll(f)
	fmt.Printf("%s", string(content))
}

func fileWrite() {
	s := data.NewStore(data.DiskStorage)
	f, err := s.Open("./practices/factory/myfile.txt")
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	n, _ := f.Write([]byte("This is a test data!\n"))
	fmt.Printf("Write: %v\n", n)
}


func fileWriteTmpFile() {
	s := data.NewStore(data.TempStorage)
	f, err := s.Open("myfile-*.txt")
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	n, _ := f.Write([]byte("This is a test data!\n"))
	fmt.Printf("Write: %v\n", n)

	// Remember to clean up the file afterwards
	typ := reflect.TypeOf(f)
	val := reflect.ValueOf(f)
	fmt.Printf("type: %v\n",typ)
	fmt.Printf("value: %v\n",val)

	name := val.FieldByName("TmpFilename")
	fmt.Printf("name: %v\n",name)

}