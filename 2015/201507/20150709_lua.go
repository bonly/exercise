package main

import "github.com/Shopify/go-lua"
import "time"

func tlo() {
	l := lua.NewState()
	lua.OpenLibraries(l)
	if err := lua.DoFile(l, "hello.lua"); err != nil {
		panic(err)
	}

	var out bytes.Buffer
	f := l.stack[l.top-1].(*luaClosure)
	err = l.Dump(&out)
	if err != nil {
		panic(err)
	}
}

func main() {
	for i := 0; i < 100; i++ {
		go tlo()
	}
	time.Sleep(10 * time.Second)
}
