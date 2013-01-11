package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl := template.New("")
	tmpl.Parse("{{range $k,$v := .}}map[{{$k}}] = {{$v}}\n{{end}}")
	tmpl.Execute(os.Stdout, &map[string]int{
		"key1": 101,
		"key2": 102,	
		"key3": 103,	
	})
}
