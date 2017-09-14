package main 

import (
"log"
"net/http"
"flag"
"html/template"
)


var srv string;

func init(){
  flag.StringVar(&srv, "s", "0.0.0.0:9997", "srv addr"); 
}

func reg_handle(){
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("public/"))));
	http.Handle("/data/", http.FileServer(http.Dir("./")));
	http.HandleFunc("/", view);
}

func main(){
	reg_handle();
	log.Printf("srv: %s", srv);
	log.Fatal(http.ListenAndServe(srv, nil));
}

func view(wr http.ResponseWriter, req *http.Request){
	tpl, _ := template.ParseFiles("tpl/view.html");
	tpl.Execute(wr, "vv")
}