package main 

import (
"log"
"net/http"
"flag"
"text/template"
"fmt"
)

var srv string;
var key = "iDgGdnvNshcDrfG0qGWwvmWNu7zRUXiL";

func init(){
    flag.StringVar(&srv, "s", "0.0.0.0:9997", "srv addr"); 
}

func main(){
	http.HandleFunc("/", map_view);
	log.Printf("srv: %s\n", srv);
	log.Fatal(http.ListenAndServe(srv, nil));
}

func get_key(param string) string{
	return key;
}

func map_view(wr http.ResponseWriter, req *http.Request){
	var err error;
	tpl := template.New("head.tpl").Funcs(template.FuncMap{"get_key":get_key});
	// fmt.Printf("%#v\n", tpl);
	tpl, err = tpl.ParseFiles("main.tpl", "head.tpl");
	// fmt.Printf("%#v\n", tpl);
	
	if err != nil{
		log.Fatal(err);
	}
	if err = tpl.ExecuteTemplate(wr, "main.tpl", nil); err != nil{
		log.Fatal(err);
	}
}

