package main

import (
	"html/template"
	"os"
)

type Matrix struct {
	Array [9][9]int
}

func main() {
        var matrix Matrix;
	tmpl, _ := template.New("example").Parse(`
        {{ $a := .Array }}
        {{ range $a }}{{ $elem := . }}|{{ range $elem }}{{ printf "%d" . }}{{ end}}|
        {{end}}`)
	tmpl.Execute(os.Stdout, matrix)
}
/*
二维数组的迭代显示
*/
