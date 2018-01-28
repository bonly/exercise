package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var Filelist []string

func main() {
	err := filepath.Walk("./", listfunc)
	if err != nil {
		fmt.Printf("filepath.Walk() error: %v\n", err)
	}
	fmt.Printf("%+v", Filelist)
}

func listfunc(path string, f os.FileInfo, err error) error {
	ok := strings.HasSuffix(path, ".go")
	if ok {
		Filelist = append(Filelist, path)
	}

	return nil
}