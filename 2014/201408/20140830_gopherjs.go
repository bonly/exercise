package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Call("alert", "Hello, JavaScript")
	println("Hello, JS console")
}

/*
gopherjs build -o main.js
*/
