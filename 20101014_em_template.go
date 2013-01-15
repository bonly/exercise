package main

import (
    "fmt"
    "os"
    "text/template"
)

func main() {
    s1, _ := template.ParseFiles("20101011_header.tmpl", "20101012_content.tmpl", "20101013_footer.tmpl")
    s1.ExecuteTemplate(os.Stdout, "header", nil)
    fmt.Println()
    s1.ExecuteTemplate(os.Stdout, "content", nil)
    fmt.Println()
    s1.ExecuteTemplate(os.Stdout, "footer", nil)
    fmt.Println()
    s1.Execute(os.Stdout, nil)
}

/*
template.ParseFiles把所有的嵌套模板全部解析到模板里面，其实每一个定义的{{define}}都是一个独立的模板，他们相互独立，是并行存在的关系，内部其实存储的是类似map的一种关系(key是模板的名称，value是模板的内容)，然后我们通过ExecuteTemplate来执行相应的子模板内容，我们可以看到header、footer都是相对独立的，都能输出内容，contenrt中因为嵌套了header和footer的内容，就会同时输出三个的内容。但是当我们执行s1.Execute，没有任何的输出，因为在默认的情况下没有默认的子模板，所以不会输出任何的东西。
同一个集合类的模板是互相知晓的，如果同一模板被多个集合使用，则它需要在多个集合中分别解析
*/

