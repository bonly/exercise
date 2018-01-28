package main

import (
	"html/template"
	"os"
)

type versionKey struct {
	Environment string
	Service     string
}

type templateData struct {
	Environments []string
	Services     []string
	Versions     map[versionKey]string
}

func (t *templateData) Version(environment, service string) string {
	return t.Versions[versionKey{
		Environment: environment,
		Service:     service,
	}]
}

func main() {
	t := template.Must(template.New("").Parse(templ))
	td := &templateData{
		Environments: []string{"EnvA", "EnvB"},
		Services:     []string{"ServA", "ServB"},
		Versions: map[versionKey]string{
			{"EnvA", "ServA"}: "1.0.0",
			{"EnvA", "ServB"}: "1.0.1",
			{"EnvB", "ServA"}: "1.0.2",
		},
	}
	if err := t.Execute(os.Stdout, td); err != nil {
		panic(err)
	}
}

const templ = `{{$environments := .Environments}}
{{$services := .Services}}
{{$versions := .Versions}}

{{range $service := $services -}}
  {{range $environment := $environments}}
    {{$environment}} - {{$service}} version: {{$.Version $environment $service}}
  {{end}}
{{end}}`
