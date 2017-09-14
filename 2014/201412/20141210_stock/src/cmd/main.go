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


type Val map[string]interface{};


type st struct{
	Title string;
	Box []Ctl;
};

type Ctl struct{
	BoxID string;
	Status string;
};

func view(wr http.ResponseWriter, req *http.Request){
	tpl := template.Must(template.ParseGlob("tpl/*.html"));

	// tp := tpl.Lookup("index.html");
	// if err := tp.Execute(wr, map[string]string{"Title":"One"}); err != nil{
	// 	log.Fatal(err);
	// }

	// if err := tpl.ExecuteTemplate(wr, "index.html", Val{"Title": "Xbed"}); err != nil{
	// 	log.Fatal(err);
	// }

	var vs st;
	vs.Title = "Four";
	vs.Box = append(vs.Box, Ctl{"1", "open"});
	vs.Box = append(vs.Box, Ctl{"2", "close"});
	if err := tpl.ExecuteTemplate(wr, "index.html", vs); err != nil{
		log.Fatal(err);
	}	

}