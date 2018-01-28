package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func main() {
	vecty.RenderBody(&MyComponent{})
}

type MyComponent struct {
	vecty.Core
}

func (mc *MyComponent) Render() *vecty.HTML {
	// return elem.Body(
	// 	&MyChildComponent{},
	// 	vecty.Text("some footer text"),
	// )
	// return elem.Div(
	// 	vecty.Markup(prop.Class("my-main-container")),
	// vecty.Text("Welcome to my site"),
	// )

	return elem.Body(
		elem.Title(
			vecty.Text("test"),
		),
		elem.Header(
			vecty.Text("head"),
		),
		elem.Div(
			vecty.Text("oooo"),
		),
		elem.Footer(
			vecty.Text("foot"),
		),
	)
}
