package main

import "github.com/Shopify/go-lua"

func main() {
	l := lua.NewState()
	lua.OpenLibraries(l)
	if err := lua.DoFile(l, "20130924_hello.lua"); err != nil {
		panic(err)
	}
}
