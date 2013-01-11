package main

import (
	"html/template"
	"os"
)

type Info struct{
	Username string
	MainPage string
}

func main() {
	t, _ := template.New("demo").Parse(`{{define "T"}}Hello, {{.Username}}! Main Page: [{{.MainPage}}]{{end}}`)
	args2 := Info{Username: "Hypermind", MainPage: "http://hypermind.com.cn/go"}
	_ = t.ExecuteTemplate(os.Stdout, "T", args2)
}

/*

*/