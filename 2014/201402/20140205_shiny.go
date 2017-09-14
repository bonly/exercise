package main

import (
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"log"
)

func main() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(nil)
		if err != nil {
			// handleError(err)
			log.Println(err);
			return
		}
		for e := range w.Events() {
			log.Println(e);
			// handleEvent(e)
		}
	})
}
