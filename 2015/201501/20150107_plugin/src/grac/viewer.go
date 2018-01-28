package main

import (
	"grac/comic"
	"log"
	"os"
	"path/filepath"
	"plugin"
)

var (
	comics = make(map[string]comic.I)
)

func loadComic(n string) error {
	p, err := plugin.Open(n)
	if err != nil {
		return err
	}

	i, err := p.Lookup("Init")
	if err != nil {
		return err
	}

	c := i.(func() comic.I)()
	comics[c.Name()] = c
	log.Printf("Loaded comic %q from %q", c.Name(), n)

	return nil
}

func loadComics() {
	plugins, err := filepath.Glob(filepath.Join(os.Args[1], "*.so"))
	if err != nil {
		panic(err)
	}

	for _, n := range plugins {
		err := loadComic(n)
		if err != nil {
			log.Printf("Failed to load comic %q: %v", n, err)
		}
	}
}

func main() {
	loadComics()
}

/*
https://echo.labstack.com/recipes/twitter
https://deedlefake.com/2016/12/go-plugins/
*/
