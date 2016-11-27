package main

import (
"log"

"golang.org/x/mobile/app"
"golang.org/x/mobile/event"
"golang.org/x/mobile/gl"
"golang.org/x/mobile/app/debug"
)

func main() {
	app.Run(app.Callbacks{
		Start: onStart,
		Draw:  onDraw,
		Touch: onTouch,
		Stop:  onStop,
	});
}

func onStart() {
	log.Println("application start")
}

func onStop() {
	log.Println("application stop");
}

func onTouch(t event.Touch) {
        log.Println("application touch");

}

func onDraw() {
	// draw background.
	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	debug.DrawFPS()
}
