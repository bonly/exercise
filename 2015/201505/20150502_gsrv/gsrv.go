package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"plugin"
)

var so_list []string;

func main(){
	err := filepath.Walk("./", get_so);
	if err != nil {
		fmt.Printf("lib list error: %v\n", err);
		return;
	}

	fmt.Printf("%+v", so_list);

	plg, err := plugin.Open(so_list[0]);
	if err != nil {
		panic(err);
	}

	sv, err := plg.Lookup("Srv");
	if err != nil {
		panic(err);
	}

	srv := sv.(func()(error));
	log.Fatal(srv());
}

func get_so(path string, f os.FileInfo, err error) error {
	ok := strings.HasSuffix(path, ".pg");
	if ok {
		so_list = append(so_list, path);
	}

	return nil;
}

