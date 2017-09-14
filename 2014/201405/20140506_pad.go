package main

import gc "github.com/rthornton128/goncurses"

func main() {
	stdscr, _ := gc.Init()
	defer gc.End()

	var panels [3]*gc.Panel
	y, x := 2, 4

	for i := 0; i < 3; i++ {
		window, _ := gc.NewWindow(10, 40, y+i, x+(i*5))
		window.Box(0, 0)
		panels[i] = gc.NewPanel(window)
	}

	gc.UpdatePanels()
	gc.Update()

	stdscr.GetChar()
}