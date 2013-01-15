package main

import (
    "html/template"
    "os"
)

type Friend struct {
    Fname string
}

type Person struct {
    UserName string
    Emails   []string
    Friends  []*Friend
}

func main() {
    f1 := Friend{Fname: "minux.ma"}
    f2 := Friend{Fname: "xushiwei"}
    t := template.New("fieldname example")
    t, _ = t.Parse(`hello {{.UserName}}!
            {{range .Emails}}
                an email {{. | html}}
            {{end}}
            {{with .Friends}}
            {{range . }}
                my friend name is {{.Fname}}
            {{end}}
            {{end}}
            `)
    p := Person{UserName: "Astaxie",
        Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
        Friends: []*Friend{&f1, &f2}}
    t.Execute(os.Stdout, p)
}

/*
Go语言模板最强大的一点就是支持pipe数据，在Go语言里面任何{{}}里面的都是pipelines数据，例如我们上面输出的email里面如果还有一些可能引起XSS注入的，那么我们如何来进行转化呢？

{{. | html}}
在email输出的地方我们可以采用如上方式可以把输出全部转化html的实体，上面的这种方式和我们平常写Unix的方式是不是一模一样，操作起来相当的简便，调用其他的函数也是类似的方式。
*/

