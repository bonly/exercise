package main

import (
   "fmt"
   "html/template"
   "log"
   "net/http"
   _ "strings"
)


func table(w http.ResponseWriter, r *http.Request) {
   fmt.Println("method:", r.Method) //获取请求的方法
   if r.Method == "GET" {
      t, _ := template.ParseFiles("index.html");
      t.Execute(w, nil);
   } else {
      //请求的是登陆数据，那么执行登陆的逻辑判断
      r.ParseForm() //必须有这个才能解释   
      fmt.Println("username:", r.Form["username"])
      fmt.Println("password:", r.Form["password"])
      /*
      Request本身也提供了FormValue()函数来获取用户提交的参数。
      如r.Form["username"]也可写成r.FormValue("username")。
      调用r.FormValue时会自动调用r.ParseForm，所以不必提前调用。
      r.FormValue只会返回同名参数中的第一个，若参数不存在则返回空字符串。
      */
   }
}

func main() {
   http.HandleFunc("/", table)         //设置访问的路由
   err := http.ListenAndServe(":9090", nil) //设置监听的端口
   if err != nil {
      log.Fatal("ListenAndServe: ", err)
   }
}