package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	file, err := ioutil.ReadFile("files/names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(file)
	fmt.Println(str)

	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

}
