package main

import (
	"html/template"
	"os"
)

func main() {
	t, _ := template.New("demo").Parse(`{{define "T"}}Hello, {{.Username}}! Main Page: [{{.MainPage}}]{{end}}`)
	args1 := map[string]string {"Username": "Hypermind", "MainPage": "http://hypermind.com.cn/go"}
	_ = t.ExecuteTemplate(os.Stdout, "T", args1)
}

/*
map
*/