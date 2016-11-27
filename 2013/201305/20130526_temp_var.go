package main

import (
	"html/template"
	"os"
	"fmt"
)

type Person struct {
	Name   string
	Emails []string
}

//使用变量$来保存想用的值
const templ = `
{{$name := .Name}}
{{range .Emails}}
    Name is {{$name}}, email is {{.}}
{{end}}
`

func main() {
	person := Person{
		Name:   "jan",
		Emails: []string{"jan@newmarch.name", "jan.newmarch@gmail.com"},
	}

	t := template.New("Person template")
	t, err := t.Parse(templ)
	checkError(err)

	err = t.Execute(os.Stdout, person)
	checkError(err)
	
	/*
    buff := bytes.NewBufferString("")
	t.Execute(buff, person)
	*/
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}