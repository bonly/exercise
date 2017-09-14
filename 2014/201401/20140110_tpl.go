package main 

import (
"fmt"
"text/template"
"os"
)

const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.{{else}}
It is a shame you couldn't make it to the wedding.{{end}}
{{with .Gift}}Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

func main(){
	tpl := template.Must(template.New("wc").Parse(letter));

	type ad struct{
		Name string;
		Attended bool;
		Gift string;
	};
	data := ad{"bonly", false, "dog"};

	err := tpl.Execute(os.Stdout, data);
	if err != nil{
		fmt.Println("tpl err: ", err);
	}
}
