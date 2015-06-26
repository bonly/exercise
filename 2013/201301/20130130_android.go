package main

import (
	"log"

	"golang.org/x/mobile/app"
)

func main() {
	app.Run(app.Callbacks{
		Draw: draw,
	})
}

func draw() {
	log.Print("In draw loop, can call OpenGL.")
}
