package main 

import(
"net/http"
"html/template"
//"text/template" //效果一致
"fmt"
)

func handler(wr http.ResponseWriter, req *http.Request){
	tpl, err := template.ParseFiles("t1.html", "t2.html");
	if err != nil{
		fmt.Println(err);
		return;
	}
	tpl.ExecuteTemplate(wr, "t1.html", "main");
}


func main(){
	http.HandleFunc("/home", handler);

	err := http.ListenAndServe("0.0.0.0:9997", nil);
	if err != nil{
		fmt.Println(err);
	}
}