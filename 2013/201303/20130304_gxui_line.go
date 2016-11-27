package main

import (

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
	"github.com/google/gxui/math"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	canvas := driver.CreateCanvas(math.Size{W: 1000, H: 1000})

	p := gxui.Polygon{
		gxui.PolygonVertex{
			Position:      math.Point{X: 0, Y: 0},
			RoundedRadius: 0,
		},
		gxui.PolygonVertex{
			Position:      math.Point{X: 100, Y: 100},
			RoundedRadius: 0,
		},
	}
	canvas.DrawLines(p, gxui.CreatePen(3, gxui.White))

	canvas.Complete()

	image := theme.CreateImage()
	image.SetCanvas(canvas)

	window := theme.CreateWindow(800, 600, "test the txt box")
	window.AddChild(image)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}
