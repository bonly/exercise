package main
import (
  "net/http"
  "html/template"
)
/*
New("TN")创建时添加模板名，可执行Execute(), TN最好与tpl.tmpl文件的{{define "TN"}}同名
*/
func handler1(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("t1.html", "t2.html") 
  t.Execute(w, "Asit"); //此时就是取默认t1.html,如t1失败，取下一个的名字
}

func handler2(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("t1.html", "t2.html") 
  t.ExecuteTemplate(w, "t2.html", "Golang"); //指定了取h2.html
}

func handler_name(w http.ResponseWriter, req *http.Request){
  t, _ := template.ParseGlob("*.html");
  t.ExecuteTemplate(w, "t3.html", "in t3");
}

func main() {
  server := http.Server{
    Addr: "127.0.0.1:8080",
  }
  
  http.HandleFunc("/t1", handler1)
  http.HandleFunc("/t2", handler2)
  http.HandleFunc("/t3", handler_name)
  server.ListenAndServe()
}