package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	"path"
	"runtime"
)

func main() {
	fmt.Printf("js: %v\n", js.Global.Get("process").Get("version"))
	fmt.Printf("electron: %v\n", js.Global.Get("process").Get("versions").Get("electron"))

	app := js.Global.Call("require", "app")
	browserWindow := js.Global.Call("require", "browser-window")

	crashReporter := js.Global.Call("require", "crash-reporter")
	crashReporter.Call("start")

	var mainWindow *js.Object

	app.Call("on", "window-all-closed", func() {
		if js.Global.Get("process").Get("platform").String() != "darwin" {
			app.Call("quit")
		}
	})

	app.Call("on", "ready", func() {
		geom := map[string]int{"width": 800, "height": 600}
		mainWindow = browserWindow.New(geom)

		// and load the index.html of the app.
		_, filename, _, _ := runtime.Caller(1)
		url := "file://" + path.Join(path.Dir(filename), "index.html")
		mainWindow.Call("loadUrl", url)

		// Emitted when the window is closed.
		mainWindow.Call("on", "closed", func() {
			mainWindow = nil
		})
	})
}