package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"regexp"
	"strings"
	"text/template"
)

const src = `
package example

type MyType int

const (
	Alpha MyType = iota // "short:a" "long:alpha"
	Beta                // "short:b" "long:beta"
	Gamma               // "short:g" "long:gamma"
	Delta               // "short:d" "long:delta"
)
`

const xref = `
package {{.PackageName}}
{{ range $type, $annotations := .TypeAnnotations }}{{ range $tag, $values := $annotations }}
var x{{$type}}_{{$tag}} = map[{{$type}}]string{ {{ range $const, $value := $values }}
	{{ $const }} : "{{$value}}",{{ end }}
}
	{{ end }}
{{ end }}
`

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	items := make(map[string]map[string]map[string]string)

	add := func(typ, name, cat, value string) {
		m, ok := items[typ]
		if !ok {
			m = make(map[string]map[string]string)
			items[typ] = m
		}
		m2, ok2 := m[cat]
		if !ok2 {
			m2 = make(map[string]string)
			m[cat] = m2
		}
		m2[name] = value
	}

	tag := regexp.MustCompile("\"[^\"]+\"")

	parse := func(typ, name, comment string) {
		defs := tag.FindAllString(comment, -1)
		for _, def := range defs {
			def = strings.Trim(def, "\"")
			sides := strings.SplitN(def, ":", 2)
			if len(sides) != 2 {
				continue
			}
			add(typ, name, strings.TrimSpace(sides[0]), strings.TrimSpace(sides[1]))
		}
	}

	lastType := ""
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ValueSpec:
			if x.Type != nil {
				if id, ok := x.Type.(*ast.Ident); ok {
					lastType = id.Name
				}
			}
			if (lastType == "") || (x.Comment == nil) {
				return false
			}

			for _, ident := range x.Names {
				for _, cmt := range x.Comment.List {
					parse(lastType, ident.Name, cmt.Text)
				}
			}
		}
		return true
	})

	t, err := template.New("").Parse(xref)
	if err != nil {
		panic(err)
	}

	t.Execute(os.Stdout, map[string]interface{}{
		"PackageName":     f.Name.Name,
		"TypeAnnotations": items,
	})
}

/*
package example

var xMyType_long = map[MyType]string{ 
	Alpha : "alpha",
	Beta : "beta",
	Delta : "delta",
	Gamma : "gamma",
}
	
var xMyType_short = map[MyType]string{ 
	Alpha : "a",
	Beta : "b",
	Delta : "d",
	Gamma : "g",
}
	
*/