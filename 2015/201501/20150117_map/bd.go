package main 

import (
"log"
"net/http"
"flag"
"text/template"
// "fmt"
)

var srv string;
var ssl bool;
var key = "iDgGdnvNshcDrfG0qGWwvmWNu7zRUXiL";

func init(){
    flag.StringVar(&srv, "s", "0.0.0.0:9997", "srv addr"); 
    flag.BoolVar(&ssl, "ssl", false, "https or not");
}

func main(){
	flag.Parse();
	http.HandleFunc("/", map_view);
	log.Printf("srv: %s\n", srv);
	if ssl == true{
	   log.Fatal(http.ListenAndServeTLS(srv, "cert.pem", "key.pem", nil));
    }else{
	   log.Fatal(http.ListenAndServe(srv, nil));
    }
}

func get_key(param string) string{
	// log.Println("call get_key");
	return key;
}

func map_view(wr http.ResponseWriter, req *http.Request){
	var err error;
	tpl := template.New("fun_img").Funcs(template.FuncMap{"get_key":get_key});
	
	// tpl, err = tpl.ParseFiles("geo.html");
	// if err != nil{
	// 	log.Fatal("geo.html ", err);
	// }
	// tpl, err = tpl.ParseFiles("head.html"); 
	// if err != nil{
	// 	log.Fatal(err);
	// }
	// log.Printf("%#v\n", tpl.DefinedTemplates());
	tpl, err = tpl.ParseFiles("main.html", "geo.html", "head.html");
	if err != nil{
		log.Fatal("main.html ", err);
	}
	// log.Printf("%#v\n", tpl.DefinedTemplates());

	// fmt.Printf("%#v\n", tpl);
	if err = tpl.ExecuteTemplate(wr, "main.html", nil); err != nil{
		log.Fatal(err);
	}
}

