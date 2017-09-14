package main

import "honnef.co/go/js/dom"

func main() {
	d := dom.GetWindow().Document()
	h := d.GetElementByID("foo")
	h.SetInnerHTML("Hello World")
}

/*
gopherjs build dom.go -o dom.js
*/