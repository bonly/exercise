package main
import (
    "text/template"
    "os"
    "fmt"
)
func myfunc(text string) string {
    return fmt.Sprintf("The length of '%s' is %d", text, len(text))
}
func chkerr(err error) {
    if err != nil {
        panic(err)
    }
}
func main(){
    T0 := `{{with .var | myfunc}}{{.}}{{end}}
    {{with $a := "This is var a" | myfunc}}{{$a}}{{end}}`
    t,err := template.New("T0").Funcs(template.FuncMap{"myfunc":myfunc}).Parse(T0)
    chkerr(err)
    err = t.Execute(os.Stdout, map[string]string{
        "var":"This is a test",
    })
    chkerr(err)
}